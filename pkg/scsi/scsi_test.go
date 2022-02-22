/*
*
* Copyright © 2020 Dell Inc. or its subsidiaries. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*   http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
 */

package scsi

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	mh "github.com/dell/gobrick/internal/mockhelper"
	wrp "github.com/dell/gobrick/internal/wrappers"
	"github.com/golang/mock/gomock"
	"golang.org/x/sync/singleflight"
)

const (
	testCHRoot = "/fakeroot"
)

func getValidHCTL() HCTL {
	return HCTL{Host: "34", Channel: "0", Target: "0", Lun: "9"}
}

func getHostOnlyHCTL() HCTL {
	return HCTL{Host: "34"}
}

type scsiFields struct {
	chroot     string
	fileReader *wrp.MockLimitedIOUtil
	filePath   *wrp.MockLimitedFilepath
	os         *wrp.MockLimitedOS
	osexec     *wrp.MockLimitedOSExec
	singleCall *singleflight.Group
}

func getDefaultSCSIFields(ctrl *gomock.Controller) scsiFields {
	ioutilMock := wrp.NewMockLimitedIOUtil(ctrl)
	filePathMock := wrp.NewMockLimitedFilepath(ctrl)
	osMock := wrp.NewMockLimitedOS(ctrl)
	osExecMock := wrp.NewMockLimitedOSExec(ctrl)
	return scsiFields{
		chroot:     testCHRoot,
		fileReader: ioutilMock,
		filePath:   filePathMock,
		os:         osMock,
		osexec:     osExecMock,
		singleCall: &singleflight.Group{},
	}
}

func TestNewSCSI(t *testing.T) {
	scsi := NewSCSI("")
	if scsi == nil {
		t.Errorf("NewSCSI() should not return nil")
	}
}

func Test_scsi_IsDeviceExist(t *testing.T) {

	type args struct {
		ctx    context.Context
		device string
	}

	ctx := context.Background()
	defaultArgs := args{ctx: ctx, device: mh.ValidDeviceName}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := mh.MockHelper{
		Ctrl:           ctrl,
		OSStatCallPath: mh.ValidDevicePath,
	}

	tests := []struct {
		name        string
		fields      scsiFields
		stateSetter func(fields scsiFields)
		args        args
		want        bool
	}{
		{
			name:   "device exist",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mocks.OSStatCallOK(fields.os)
			},
			args: defaultArgs,
			want: true,
		},
		{
			name:   "device not found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mocks.OSStatCallErr(fields.os)
			},
			args: defaultArgs,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			if got := s.IsDeviceExist(tt.args.ctx, tt.args.device); got != tt.want {
				t.Errorf("IsDeviceExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scsi_RescanSCSIHostByHCTL(t *testing.T) {
	type args struct {
		ctx  context.Context
		addr HCTL
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, addr: getValidHCTL()}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := mh.MockHelper{
		Ctrl: ctrl,
		OSOpenFileCallPath: fmt.Sprintf(
			"/sys/class/scsi_host/host%s/scan", defaultArgs.addr.Host),
		FileWriteStringCallData: fmt.Sprintf("%s %s %s", defaultArgs.addr.Channel,
			defaultArgs.addr.Target, defaultArgs.addr.Lun),
	}

	tests := []struct {
		name        string
		fields      scsiFields
		stateSetter func(fields scsiFields)
		args        args
		wantErr     bool
	}{
		{
			name:   "open file error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mocks.OSOpenFileCallErr(fields.os)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "write to file error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				_, fileMock := mocks.OSOpenFileCallOK(fields.os)
				mocks.FileWriteStringErr(fileMock)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "file close error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				_, fileMock := mocks.OSOpenFileCallOK(fields.os)
				mocks.FileWriteStringOK(fileMock)
				mocks.FileCloseErr(fileMock)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "rescan without error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				_, fileMock := mocks.OSOpenFileCallOK(fields.os)
				mocks.FileWriteStringOK(fileMock)
				mocks.FileCloseOK(fileMock)
			},
			args:    defaultArgs,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			if err := s.RescanSCSIHostByHCTL(tt.args.ctx, tt.args.addr); (err != nil) != tt.wantErr {
				t.Errorf("RescanSCSIHostByHCTL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scsi_RescanSCSIDeviceByHCTL(t *testing.T) {
	type args struct {
		ctx context.Context
		h   HCTL
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, h: getValidHCTL()}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mh.MockHelper{
		Ctrl: ctrl,
		OSOpenFileCallPath: fmt.Sprintf(
			"/sys/class/scsi_device/%s:%s:%s:%s/device/rescan",
			defaultArgs.h.Host,
			defaultArgs.h.Channel,
			defaultArgs.h.Target,
			defaultArgs.h.Lun),
		FileWriteStringCallData: "1",
	}

	tests := []struct {
		name        string
		fields      scsiFields
		stateSetter func(fields scsiFields)
		args        args
		wantErr     bool
	}{
		{
			name:        "HCTL is not full",
			fields:      getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {},
			args:        args{ctx: ctx, h: getHostOnlyHCTL()},
			wantErr:     true,
		},
		{
			name:   "file open error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.OSOpenFileCallErr(fields.os)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "write to file error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				_, fileMock := mock.OSOpenFileCallOK(fields.os)
				mock.FileWriteStringErr(fileMock)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "file close error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				_, fileMock := mock.OSOpenFileCallOK(fields.os)
				mock.FileWriteStringOK(fileMock)
				mock.FileCloseErr(fileMock)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "rescan without error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				_, fileMock := mock.OSOpenFileCallOK(fields.os)
				mock.FileWriteStringOK(fileMock)
				mock.FileCloseOK(fileMock)
			},
			args:    defaultArgs,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			if err := s.RescanSCSIDeviceByHCTL(tt.args.ctx, tt.args.h); (err != nil) != tt.wantErr {
				t.Errorf("RescanSCSIDeviceByHCTL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type t struct {
	name        string
	fields      scsiFields
	stateSetter func(fields scsiFields)
	args        interface{}
	wantErr     bool
}

func getDeleteSCSIDeviceTestCases(mocks mh.MockHelper, defaultArgs interface{},
	ctrl *gomock.Controller) []t {
	return []t{
		{
			name:   "error read state file",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mocks.IOUTILReadFileErr(fields.fileReader)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "device in blocked state",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mocks.IOUTILReadFileOKReturn = "blocked"
				mocks.IOUTILReadFileOK(fields.fileReader)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "error open delete file",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mocks.IOUTILReadFileOKReturn = "running"
				mocks.IOUTILReadFileOK(fields.fileReader)
				mocks.OSOpenFileCallErr(fields.os)
			},
			args:    defaultArgs,
			wantErr: false,
		},
		{
			name:   "error write to delete file",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mocks.IOUTILReadFileOK(fields.fileReader)
				_, fileMock := mocks.OSOpenFileCallOK(fields.os)
				mocks.FileWriteStringErr(fileMock)
				mocks.FileCloseErr(fileMock)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "delete without error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mocks.IOUTILReadFileOK(fields.fileReader)
				_, fileMock := mocks.OSOpenFileCallOK(fields.os)
				mocks.FileWriteStringOK(fileMock)
				mocks.FileCloseOK(fileMock)
			},
			args:    defaultArgs,
			wantErr: false,
		},
	}
}

func Test_scsi_DeleteSCSIDeviceByHCTL(t *testing.T) {
	type args struct {
		ctx context.Context
		h   HCTL
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, h: getValidHCTL()}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sysPath := fmt.Sprintf("/sys/class/scsi_device/%s:%s:%s:%s/device/",
		defaultArgs.h.Host,
		defaultArgs.h.Channel,
		defaultArgs.h.Target,
		defaultArgs.h.Lun)

	mock := mh.MockHelper{
		Ctrl:                    ctrl,
		OSOpenFileCallPath:      sysPath + "delete",
		FileWriteStringCallData: "1",
		IOUTILReadFileCallPath:  sysPath + "state",
	}

	tests := getDeleteSCSIDeviceTestCases(mock, defaultArgs, ctrl)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			if err := s.DeleteSCSIDeviceByHCTL(tt.args.(args).ctx, tt.args.(args).h); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSCSIDeviceByHCTL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scsi_DeleteSCSIDeviceByName(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, name: mh.ValidDeviceName}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mh.MockHelper{
		Ctrl: ctrl,
		OSOpenFileCallPath: fmt.Sprintf(
			"/sys/class/block/%s/device/delete", mh.ValidDeviceName),
		FileWriteStringCallData: "1",
		IOUTILReadFileCallPath: fmt.Sprintf(
			"/sys/class/block/%s/device/state", mh.ValidDeviceName),
	}

	tests := getDeleteSCSIDeviceTestCases(mock, defaultArgs, ctrl)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			if err := s.DeleteSCSIDeviceByName(tt.args.(args).ctx, tt.args.(args).name); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSCSIDeviceByName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scsi_DeleteSCSIDeviceByPath(t *testing.T) {
	type args struct {
		ctx     context.Context
		devPath string
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, devPath: fmt.Sprintf("/sys/class/block/%s/device", mh.ValidDeviceName)}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sysPath := fmt.Sprintf("/sys/class/block/%s/device/", mh.ValidDeviceName)

	mock := mh.MockHelper{
		Ctrl:                    ctrl,
		OSOpenFileCallPath:      sysPath + "delete",
		FileWriteStringCallData: "1",
		IOUTILReadFileCallPath:  sysPath + "state",
	}
	tests := getDeleteSCSIDeviceTestCases(mock, defaultArgs, ctrl)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			if err := s.DeleteSCSIDeviceByPath(
				tt.args.(args).ctx, tt.args.(args).devPath); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSCSIDeviceByPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scsi_GetDeviceWWN(t *testing.T) {
	type args struct {
		ctx     context.Context
		devices []string
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, devices: []string{mh.ValidDeviceName}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mh.MockHelper{
		Ctrl: ctrl,
		IOUTILReadFileCallPath: fmt.Sprintf(
			"/sys/block/%s/device/wwid", mh.ValidDeviceName),
		OSEXECCommandContextName: "chroot",
		OSEXECCommandContextArgs: []string{testCHRoot, scsiIDPath, "-g", "-p", "0x83", "/dev/" + mh.ValidDeviceName},
	}

	tests := []struct {
		name        string
		fields      scsiFields
		stateSetter func(fields scsiFields)
		args        args
		want        string
		wantErr     bool
	}{
		{
			name:   "not found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.IOUTILReadFileErr(fields.fileReader)
				mock.OSIsNotExistOK(fields.os)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.IOUTILReadFileOKReturn = mh.ValidSYSFCWWID
				mock.IOUTILReadFileOK(fields.fileReader)
			},
			args:    defaultArgs,
			wantErr: false,
			want:    mh.ValidWWID,
		},
		{
			name:   "not found with scsi_id",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.IOUTILReadFileErr(fields.fileReader)
				mock.OSIsNotExistOKReturn = true
				mock.OSIsNotExistOK(fields.os)
				_, cmdMock := mock.OSExecCommandContextOK(fields.osexec)
				mock.OSExecCmdErr(cmdMock)
			},
			args:    defaultArgs,
			wantErr: true,
			want:    "",
		},
		{
			name:   "found with scsi_id",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.IOUTILReadFileErr(fields.fileReader)
				mock.OSIsNotExistOKReturn = true
				mock.OSIsNotExistOK(fields.os)
				_, cmdMock := mock.OSExecCommandContextOK(fields.osexec)
				mock.OSEXECCmdOKReturn = mh.ValidWWID
				mock.OSExecCmdOK(cmdMock)
			},
			args:    defaultArgs,
			wantErr: false,
			want:    mh.ValidWWID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				chroot:     tt.fields.chroot,
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			got, err := s.GetDeviceWWN(tt.args.ctx, tt.args.devices)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDeviceWWN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDeviceWWN() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scsi_GetDevicesByWWN(t *testing.T) {
	type args struct {
		ctx context.Context
		wwn string
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, wwn: mh.ValidWWID}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wwidPath := "/device/wwid"
	devicePath := "/sys/block/" + mh.ValidDeviceName
	devicePath2 := "/sys/block/" + mh.ValidDeviceName2
	deviceWWIDPath := devicePath + wwidPath
	deviceWWIDPath2 := devicePath2 + wwidPath

	mock := mh.MockHelper{
		Ctrl:                    ctrl,
		FilePathGlobCallPattern: "/sys/block/sd*",
		FilePathGlobOKReturn:    []string{devicePath},
		IOUTILReadFileCallPath:  deviceWWIDPath,
		IOUTILReadFileOKReturn:  mh.ValidSYSFCWWID,
	}

	tests := []struct {
		name        string
		fields      scsiFields
		args        args
		stateSetter func(fields scsiFields)
		want        []string
		wantErr     bool
	}{
		{
			name:   "error resolve glob pattern",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobErr(fields.filePath)
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name:   "found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOK(fields.filePath)
				mock.IOUTILReadFileOK(fields.fileReader)
			},
			args:    defaultArgs,
			want:    []string{mh.ValidDeviceName},
			wantErr: false,
		},
		{
			name:   "multipile found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOKReturn = []string{devicePath, devicePath2}
				mock.FilePathGlobOK(fields.filePath)
				mock.IOUTILReadFileOK(fields.fileReader)
				mock.IOUTILReadFileCallPath = deviceWWIDPath2
				mock.IOUTILReadFileOK(fields.fileReader)
			},
			args:    defaultArgs,
			want:    []string{mh.ValidDeviceName, "sdf"},
			wantErr: false,
		},
		{
			name:   "not found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOKReturn = []string{}
				mock.FilePathGlobOK(fields.filePath)
			},
			args:    defaultArgs,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			got, err := s.GetDevicesByWWN(tt.args.ctx, tt.args.wwn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDevicesByWWN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDevicesByWWN() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scsi_GetDMDeviceByChildren(t *testing.T) {
	type args struct {
		ctx     context.Context
		devices []string
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, devices: []string{mh.ValidDeviceName, mh.ValidDeviceName2}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sysPath1 := fmt.Sprintf("/sys/block/%s/holders/dm-", mh.ValidDeviceName)
	sysPath2 := fmt.Sprintf("/sys/block/%s/holders/dm-", mh.ValidDeviceName2)

	sysPath1Glob := sysPath1 + "*"
	sysPath2Glob := sysPath2 + "*"

	dm1Path := sysPath1 + "1"
	dm2Path := sysPath1 + "2"

	mock := mh.MockHelper{
		Ctrl: ctrl,
	}

	tests := []struct {
		name        string
		fields      scsiFields
		args        args
		stateSetter func(fields scsiFields)
		want        string
		wantErr     bool
	}{
		{
			name:   "error while resolve glob",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobCallPattern = sysPath1Glob
				mock.FilePathGlobErr(fields.filePath)
			},
			args:    defaultArgs,
			want:    "",
			wantErr: true,
		},
		{
			name:   "not found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOKReturn = []string{}
				mock.FilePathGlobCallPattern = sysPath1Glob
				mock.FilePathGlobOK(fields.filePath)
				mock.FilePathGlobCallPattern = sysPath2Glob
				mock.FilePathGlobOK(fields.filePath)
			},
			args:    defaultArgs,
			want:    "",
			wantErr: true,
		},
		{
			name:   "check another devices if one failed",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOKReturn = []string{dm1Path}
				mock.FilePathGlobCallPattern = sysPath1Glob
				mock.FilePathGlobOK(fields.filePath)
				mock.FilePathGlobOKReturn = []string{dm2Path}
				mock.FilePathGlobCallPattern = sysPath2Glob
				mock.FilePathGlobOK(fields.filePath)
				mock.IOUTILReadFileCallPath = dm1Path + "/dm/uuid"
				mock.IOUTILReadFileErr(fields.fileReader)
				mock.IOUTILReadFileOKReturn = "mpath"
				mock.IOUTILReadFileCallPath = dm2Path + "/dm/uuid"
				mock.IOUTILReadFileOK(fields.fileReader)
			},
			args:    defaultArgs,
			want:    "dm-2",
			wantErr: false,
		},
		{
			name:   "devices have different parents",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOKReturn = []string{dm1Path}
				mock.FilePathGlobCallPattern = sysPath1Glob
				mock.FilePathGlobOK(fields.filePath)
				mock.FilePathGlobOKReturn = []string{dm2Path}
				mock.FilePathGlobCallPattern = sysPath2Glob
				mock.FilePathGlobOK(fields.filePath)
				mock.IOUTILReadFileOKReturn = "mpath"
				mock.IOUTILReadFileCallPath = dm1Path + "/dm/uuid"
				mock.IOUTILReadFileOK(fields.fileReader)
				mock.IOUTILReadFileCallPath = dm2Path + "/dm/uuid"
				mock.IOUTILReadFileOK(fields.fileReader)
			},
			args:    defaultArgs,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			got, err := s.GetDMDeviceByChildren(tt.args.ctx, tt.args.devices)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDMDeviceByChildren() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDMDeviceByChildren() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scsi_CheckDeviceIs_Valid(t *testing.T) {
	type args struct {
		ctx    context.Context
		device string
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, device: mh.ValidDevicePath}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mh.MockHelper{
		Ctrl:                     ctrl,
		OSStatCallPath:           mh.ValidDevicePath,
		OSEXECCommandContextName: "dd",
		OSEXECCommandContextArgs: []string{
			"if=" + mh.ValidDevicePath, "of=/dev/null", "bs=1k", "count=1"},
		OSEXECCmdOKReturn: "invalid",
	}

	tests := []struct {
		name        string
		fields      scsiFields
		stateSetter func(fields scsiFields)
		args        args
		want        bool
	}{
		{
			name:   "not exist",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.OSStatCallErr(fields.os)
			},
			args: defaultArgs,
			want: false,
		},
		{
			name:   "dd command failed",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.OSStatCallOK(fields.os)
				_, cmdMock := mock.OSExecCommandContextOK(fields.osexec)
				mock.OSExecCmdErr(cmdMock)
			},
			args: defaultArgs,
			want: false,
		},
		{
			name:   "dd invalid output",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.OSStatCallOK(fields.os)
				_, cmdMock := mock.OSExecCommandContextOK(fields.osexec)
				mock.OSExecCmdOK(cmdMock)
			},
			args: defaultArgs,
			want: false,
		},
		{
			name:   "valid",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.OSStatCallOK(fields.os)
				_, cmdMock := mock.OSExecCommandContextOK(fields.osexec)
				mock.OSEXECCmdOKReturn = "1+0 records in\n" +
					"1+0 records out\n" +
					"1024 bytes (1,0 kB, 1,0 KiB) copied, 0,000110574 s, 9,3 MB/s"
				mock.OSExecCmdOK(cmdMock)
			},
			args: defaultArgs,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			if got := s.CheckDeviceIsValid(tt.args.ctx, tt.args.device); got != tt.want {
				t.Errorf("CheckDeviceIsmh.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scsi_GetDeviceNameByHCTL(t *testing.T) {
	type args struct {
		ctx context.Context
		h   HCTL
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, h: getValidHCTL()}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sysPath := fmt.Sprintf("/sys/class/scsi_device/%s:%s:%s:%s/device/block/",
		defaultArgs.h.Host,
		defaultArgs.h.Channel,
		defaultArgs.h.Target,
		defaultArgs.h.Lun)

	mock := mh.MockHelper{
		Ctrl:                    ctrl,
		FilePathGlobCallPattern: sysPath + "*",
		FilePathGlobOKReturn:    []string{sysPath + mh.ValidDeviceName},
	}

	tests := []struct {
		name        string
		fields      scsiFields
		args        args
		stateSetter func(fields scsiFields)
		want        string
		wantErr     bool
	}{
		{
			name:        "invalid HCTL",
			fields:      getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {},
			args:        args{ctx: ctx, h: getHostOnlyHCTL()},
			want:        "",
			wantErr:     true,
		},
		{
			name:   "glob resolve error",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobErr(fields.filePath)
			},
			args:    defaultArgs,
			want:    "",
			wantErr: true,
		},
		{
			name:   "device found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOK(fields.filePath)
			},
			args:    defaultArgs,
			want:    mh.ValidDeviceName,
			wantErr: false,
		},
		{
			name:   "device not found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOKReturn = []string{}
				mock.FilePathGlobOK(fields.filePath)
			},
			args:    defaultArgs,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			got, err := s.GetDeviceNameByHCTL(tt.args.ctx, tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDeviceNameByHCTL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDeviceNameByHCTL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scsi_GetDMChildren(t *testing.T) {
	type args struct {
		ctx context.Context
		dm  string
	}

	ctx := context.Background()

	defaultArgs := args{ctx: ctx, dm: mh.ValidDMName}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sysPath := fmt.Sprintf("/sys/block/%s/slaves/", mh.ValidDMName)

	mock := mh.MockHelper{
		Ctrl:                    ctrl,
		FilePathGlobCallPattern: sysPath + "*",
		FilePathGlobOKReturn: []string{
			sysPath + mh.ValidDeviceName,
			sysPath + mh.ValidDeviceName2},
	}

	tests := []struct {
		name        string
		fields      scsiFields
		stateSetter func(fields scsiFields)
		args        args
		want        []string
		wantErr     bool
	}{
		{
			name:   "glob err",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobErr(fields.filePath)
			},
			args:    defaultArgs,
			want:    nil,
			wantErr: true,
		},
		{
			name:   "found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOK(fields.filePath)
			},
			args:    defaultArgs,
			want:    []string{mh.ValidDeviceName, mh.ValidDeviceName2},
			wantErr: false,
		},
		{
			name:   "not found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathGlobOKReturn = nil
				mock.FilePathGlobOK(fields.filePath)
			},
			args:    defaultArgs,
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			got, err := s.GetDMChildren(tt.args.ctx, tt.args.dm)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDMChildren() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDMChildren() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scsi_WaitUdevSymlink(t *testing.T) {
	type args struct {
		ctx        context.Context
		deviceName string
		wwn        string
	}

	ctx := context.Background()

	devArgs := args{ctx: ctx, deviceName: mh.ValidDeviceName, wwn: mh.ValidWWID}
	dmArgs := args{ctx: ctx, deviceName: mh.ValidDMName, wwn: mh.ValidWWID}

	devPath := fmt.Sprintf("/dev/disk/by-id/scsi-%s", mh.ValidWWID)
	dmPath := fmt.Sprintf("/dev/disk/by-id/dm-uuid-mpath-%s", mh.ValidWWID)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mh.MockHelper{
		Ctrl: ctrl,
	}

	tests := []struct {
		name        string
		fields      scsiFields
		stateSetter func(fields scsiFields)
		args        args
		wantErr     bool
	}{
		{
			name:   "dev found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathEvalSymlinksCallPath = devPath
				mock.FilePathEvalSymlinksOKReturn = mh.ValidDevicePath
				mock.FilePathEvalSymlinksOK(fields.filePath)
			},
			args:    devArgs,
			wantErr: false,
		},
		{
			name:   "not found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathEvalSymlinksErr(fields.filePath)
			},
			args:    devArgs,
			wantErr: true,
		},
		{
			name:   "symlink point to unexpected device",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathEvalSymlinksOKReturn = mh.ValidDevicePath2
				mock.FilePathEvalSymlinksOK(fields.filePath)
			},
			args:    devArgs,
			wantErr: true,
		},
		{
			name:   "dm found",
			fields: getDefaultSCSIFields(ctrl),
			stateSetter: func(fields scsiFields) {
				mock.FilePathEvalSymlinksCallPath = dmPath
				mock.FilePathEvalSymlinksOKReturn = mh.ValidDMPath
				mock.FilePathEvalSymlinksOK(fields.filePath)
			},
			args:    dmArgs,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scsi{
				fileReader: tt.fields.fileReader,
				filePath:   tt.fields.filePath,
				os:         tt.fields.os,
				osexec:     tt.fields.osexec,
				singleCall: tt.fields.singleCall,
			}
			tt.stateSetter(tt.fields)
			if err := s.WaitUdevSymlink(tt.args.ctx, tt.args.deviceName, tt.args.wwn); (err != nil) != tt.wantErr {
				t.Errorf("WaitUdevSymlink() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
