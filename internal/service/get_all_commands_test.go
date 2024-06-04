package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zsandibe/command-runner-service/internal/entity"
	mock "github.com/zsandibe/command-runner-service/internal/service/mocks"
)

func TestService_GetAllCommands(t *testing.T) {
	type fields struct {
		Command      *mock.MockCommand
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		c context.Context
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
				c: context.Background(),
			},
			want: &[]entity.Command{
				{
					Id:          1,
					Script:      "echo Hello world!",
					Description: "simple script",
				},
				{
					Id:          2,
					Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					Description: "combined script",
				},
				{
					Id:          3,
					Script:      "echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					Description: "combined script with sleeping",
				},
				{
					Id:          4,
					Script:      "ls",
					Description: "simple script for look files",
				},
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().GetAllCommands(
					context.Background()).Return(&[]entity.Command{
					{
						Id:          1,
						Script:      "echo Hello world!",
						Description: "simple script",
					},
					{
						Id:          2,
						Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
						Description: "combined script",
					},
					{
						Id:          3,
						Script:      "echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
						Description: "combined script with sleeping",
					},
					{
						Id:          4,
						Script:      "ls",
						Description: "simple script for look files",
					},
				}, nil)
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
			got, err := service.GetAllCommands(tt.args.c)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopRunner()
			require.Equal(t, tt.want, got)
		})
	}
}
