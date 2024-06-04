package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zsandibe/command-runner-service/internal/entity"
	mock "github.com/zsandibe/command-runner-service/internal/service/mocks"
)

func TestService_GetCommands(t *testing.T) {
	type fields struct {
		Command      *mock.MockCommand
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		c   context.Context
		ids []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]entity.Command
		wantErr bool
		prepare func(args2 args, fields fields)
	}{
		{
			name: "test_1, error not expected",
			args: args{
				c:   context.Background(),
				ids: []string{"1", "2", "3", "4"},
			},
			want: &[]entity.Command{
				{
					Id:          int64(1),
					Script:      "echo Hello world!",
					Description: "simple script",
				},
				{
					Id:          int64(2),
					Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					Description: "combined script",
				},
				{
					Id:          int64(3),
					Script:      "echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					Description: "combined script with sleeping",
				},
				{
					Id:          int64(4),
					Script:      "ls",
					Description: "simple script for look files",
				},
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().GetCommands(
					context.Background(), []int64{1, 2, 3, 4}).Return(&[]entity.Command{
					{
						Id:          int64(1),
						Script:      "echo Hello world!",
						Description: "simple script",
					},
					{
						Id:          int64(2),
						Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
						Description: "combined script",
					},
					{
						Id:          int64(3),
						Script:      "echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
						Description: "combined script with sleeping",
					},
					{
						Id:          int64(4),
						Script:      "ls",
						Description: "simple script for look files",
					},
				}, nil)
				var i int64
				resps := []string{
					"echo Hello world!",
					"echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					"ls",
				}
				for i = 0; i < 4; i++ {
					fields.ScriptsCache.EXPECT().Set(i+1, resps[i]).Return(nil)
				}
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
			got, err := service.GetCommands(tt.args.c, tt.args.ids)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopRunner()
			require.Equal(t, tt.want, got)
		})
	}
}
