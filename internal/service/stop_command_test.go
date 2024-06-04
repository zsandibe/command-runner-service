package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock "github.com/zsandibe/command-runner-service/internal/service/mocks"
)

func TestService_StopCommandById(t *testing.T) {
	type fields struct {
		Command      *mock.MockCommand
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name: "test_1, error not expected",
			args: args{
				id: 1,
			},
			wantErr: true,
			prepare: func(args args, fields fields) {
				fields.ExecCmdCache.EXPECT().Get(int64(1)).Return(nil, errors.New("not found value"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Command:      mock.NewMockCommand(ctrl),
				ScriptsCache: mock.NewMockCache(ctrl),
				ExecCmdCache: mock.NewMockCache(ctrl),
			}
			tt.prepare(tt.args, f)

			service := NewService(f.Command, f.ScriptsCache, f.ExecCmdCache)
			err := service.StopCommandById(tt.args.id)

			service.StopRunner()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
		})
	}
}
