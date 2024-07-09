/*
Copyright © 2020-2022 Dell Inc. or its subsidiaries. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gobrick

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/dell/gonvme"

	intmultipath "github.com/dell/gobrick/internal/multipath"
	intscsi "github.com/dell/gobrick/internal/scsi"
	wrp "github.com/dell/gobrick/internal/wrappers"
	"github.com/golang/mock/gomock"
	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
)

var (
	validNVMEPortal1     = "1.1.1.1:3260"
	validNVMETarget1     = "nqn.2014-08.org.nvmexpress:uuid:csi_master"
	validNVMEPortal2     = "1.1.1.1:3260"
	validNVMETarget2     = "nqn.2014-08.org.nvmexpress:uuid:csi_worker"
	validNVMETargetInfo1 = NVMeTargetInfo{
		Portal: validNVMEPortal1,
		Target: validNVMETarget1,
	}
	validNVMETargetInfo2 = NVMeTargetInfo{
		Portal: validNVMEPortal2,
		Target: validNVMETarget2,
	}
	validNVMEVolumeInfo = NVMeVolumeInfo{
		Targets: []NVMeTargetInfo{validNVMETargetInfo1, validNVMETargetInfo2},
		WWN:     validNQN,
	}

	validLibNVMETarget1 = gonvme.NVMeTarget{
		TargetNqn: validNVMETarget1,
		Portal:    validNVMEPortal1,
	}

	validLibNVMETarget2 = gonvme.NVMeTarget{
		TargetNqn: validNVMETarget2,
		Portal:    validNVMEPortal2,
	}

	validNVMEInitiatorName = "nqn.2014-08.org.nvmexpress:uuid:csi_worker:e16da41ba075"

	validLibNVMESession1 = gonvme.NVMESession{
		Target:            validNVMETarget1,
		Portal:            validNVMEPortal1,
		Name:              "nvme1",
		NVMESessionState:  "live",
		NVMETransportName: "tcp",
	}
	validLibNVMESession2 = gonvme.NVMESession{
		Target:            validNVMETarget2,
		Portal:            validNVMEPortal2,
		Name:              "nvme2",
		NVMESessionState:  "live",
		NVMETransportName: "tcp",
	}
	validLibNVMESessions = []gonvme.NVMESession{validLibNVMESession1, validLibNVMESession2}
)

type NVMEFields struct {
	baseConnector                          *baseConnector
	multipath                              *intmultipath.MockMultipath
	scsi                                   *intscsi.MockSCSI
	nvmeLib                                *gonvme.MockNVMe
	filePath                               *wrp.MockLimitedFilepath
	manualSessionManagement                bool
	waitDeviceTimeout                      time.Duration
	waitDeviceRegisterTimeout              time.Duration
	failedSessionMinimumLoginRetryInterval time.Duration
	loginLock                              *rateLock
	limiter                                *semaphore.Weighted
	singleCall                             *singleflight.Group
}

func getDefaultNVMEFields(ctrl *gomock.Controller) NVMEFields {
	con := NewNVMeConnector(NVMeConnectorParams{})
	bc := con.baseConnector
	mpMock := intmultipath.NewMockMultipath(ctrl)
	scsiMock := intscsi.NewMockSCSI(ctrl)
	nvmeMock := gonvme.NewMockNVMe(map[string]string{})
	bc.multipath = mpMock
	bc.scsi = scsiMock
	return NVMEFields{
		baseConnector:                          bc,
		multipath:                              mpMock,
		scsi:                                   scsiMock,
		nvmeLib:                                nvmeMock,
		filePath:                               wrp.NewMockLimitedFilepath(ctrl),
		manualSessionManagement:                con.manualSessionManagement,
		waitDeviceTimeout:                      con.waitDeviceTimeout,
		waitDeviceRegisterTimeout:              con.waitDeviceRegisterTimeout,
		failedSessionMinimumLoginRetryInterval: con.waitDeviceTimeout,
		loginLock:                              con.loginLock,
		limiter:                                con.limiter,
		singleCall:                             con.singleCall,
	}
}

func TestNVME_Connector_ConnectVolume(t *testing.T) {
	type args struct {
		ctx  context.Context
		info NVMeVolumeInfo
	}

	ctx := context.Background()
	defaultArgs := args{ctx: ctx, info: validNVMEVolumeInfo}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := baseMockHelper{
		Ctx: ctx,
	}

	tests := []struct {
		name        string
		fields      NVMEFields
		args        args
		stateSetter func(fields NVMEFields)
		want        Device
		wantErr     bool
		isFC        bool
	}{
		{
			name:        "empty request",
			fields:      getDefaultNVMEFields(ctrl),
			stateSetter: func(_ NVMEFields) {},
			args:        args{ctx: ctx, info: NVMeVolumeInfo{}},
			want:        Device{},
			wantErr:     true,
			isFC:        false,
		},
		{
			name:   "not found-single device",
			fields: getDefaultNVMEFields(ctrl),
			stateSetter: func(fields NVMEFields) {
				mock.MultipathIsDaemonRunningOKReturn = false
				mock.MultipathIsDaemonRunningOK(fields.multipath)
			},
			args:    defaultArgs,
			want:    Device{},
			wantErr: true,
			isFC:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &NVMeConnector{
				baseConnector:             tt.fields.baseConnector,
				multipath:                 tt.fields.multipath,
				scsi:                      tt.fields.scsi,
				nvmeLib:                   tt.fields.nvmeLib,
				manualSessionManagement:   tt.fields.manualSessionManagement,
				waitDeviceTimeout:         tt.fields.waitDeviceTimeout,
				waitDeviceRegisterTimeout: tt.fields.waitDeviceRegisterTimeout,
				loginLock:                 tt.fields.loginLock,
				limiter:                   tt.fields.limiter,
				singleCall:                tt.fields.singleCall,
				filePath:                  tt.fields.filePath,
			}
			tt.stateSetter(tt.fields)
			got, err := c.ConnectVolume(tt.args.ctx, tt.args.info, tt.isFC)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectVolume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectVolume() got = %v, want %v", got, tt.want)
			}
		})
	}
}
