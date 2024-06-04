package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zsandibe/command-runner-service/internal/entity"
	mock "github.com/zsandibe/command-runner-service/internal/service/mocks"
)

func TestService_DeleteCommandById(t *testing.T) {
	type fields struct {
		Command      *mock.MockCommand
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		c  context.Context
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Command
		wantErr bool
		prepare func(args2 args, fields fields)
	}{
		{
			name: "test_1, error not expected",
			args: args{
				c:  context.Background(),
				id: 1,
			},
			want: nil,
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().DeleteCommandById(
					context.Background(),
					int64(1)).Return(nil)
			},
		}, {
			name: "test_2, error not expected",
			args: args{
				c:  context.Background(),
				id: 2,
			},
			want: nil,
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().DeleteCommandById(
					context.Background(),
					int64(2)).Return(nil)
			},
		}, {
			name: "test_3, error not expected",
			args: args{
				c:  context.Background(),
				id: 3,
			},
			want: nil,
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().DeleteCommandById(
					context.Background(),
					int64(3)).Return(nil)
			},
		}, {
			name: "test_4, error not expected",
			args: args{
				c:  context.Background(),
				id: 4,
			},
			want: nil,
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().DeleteCommandById(
					context.Background(),
					int64(4)).Return(nil)
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
			err := service.DeleteCommandById(tt.args.c, tt.args.id)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopRunner()
		})
	}
}
