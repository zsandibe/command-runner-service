package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zsandibe/command-runner-service/internal/entity"
	mock "github.com/zsandibe/command-runner-service/internal/service/mocks"
)

func TestService_GetCommandById(t *testing.T) {
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
			want: &entity.Command{
				Id:          1,
				Script:      "echo Hello world!",
				Description: "simple script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().GetCommandById(
					context.Background(),
					int64(1)).Return(&entity.Command{
					Id:          1,
					Script:      "echo Hello world!",
					Description: "simple script",
				}, nil)
				fields.ScriptsCache.EXPECT().Set(int64(1), "echo Hello world!").Return(nil)
			},
		}, {
			name: "test_2, error not expected",
			args: args{
				c:  context.Background(),
				id: 2,
			},
			want: &entity.Command{
				Id:          2,
				Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
				Description: "combined script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().GetCommandById(
					context.Background(),
					int64(2)).Return(&entity.Command{
					Id:          2,
					Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					Description: "combined script",
				}, nil)
				fields.ScriptsCache.EXPECT().Set(int64(2), "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami").Return(nil)
			},
		}, {
			name: "test_3, error not expected",
			args: args{
				c:  context.Background(),
				id: 3,
			},
			want: &entity.Command{
				Id:          3,
				Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
				Description: "combined script with sleeping",
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().GetCommandById(
					context.Background(),
					int64(3)).Return(&entity.Command{
					Id:          3,
					Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					Description: "combined script with sleeping",
				}, nil)
				fields.ScriptsCache.EXPECT().Set(int64(3), "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"").Return(nil)
			},
		}, {
			name: "test_4, error not expected",
			args: args{
				c:  context.Background(),
				id: 4,
			},
			want: &entity.Command{
				Id:          4,
				Script:      "ls",
				Description: "simple script for look files",
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().GetCommandById(
					context.Background(),
					int64(4)).Return(&entity.Command{
					Id:          4,
					Script:      "ls",
					Description: "simple script for look files",
				}, nil)
				fields.ScriptsCache.EXPECT().Set(int64(4), "ls").Return(nil)
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
			got, err := service.GetCommandById(tt.args.c, tt.args.id)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopRunner()
			require.Equal(t, tt.want, got)
		})
	}
}
