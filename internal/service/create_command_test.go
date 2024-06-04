package service

import (
	"context"
	"testing"

	"github.com/zsandibe/command-runner-service/internal/entity"
	mock "github.com/zsandibe/command-runner-service/internal/service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zsandibe/command-runner-service/internal/domain"
)

func TestService_CreateCommand(t *testing.T) {
	type fields struct {
		Command      *mock.MockCommand
		ScriptsCache *mock.MockCache
		ExecCmdCache *mock.MockCache
	}
	type args struct {
		c   context.Context
		inp *domain.CreateCommandInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.CreateCommandOutput
		wantErr bool
		prepare func(args2 args, fields fields)
	}{
		{
			name: "test_1, error not expected",
			args: args{
				c: context.Background(), inp: &domain.CreateCommandInput{
					Script:      "echo Hello world!",
					Description: "simple script",
				},
			},
			want: &domain.CreateCommandOutput{
				Id:          "1",
				Script:      "echo Hello world!",
				Description: "simple script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.inp.Script,
					Description: args2.inp.Description,
				}).Return(&entity.Command{
					Id:          1,
					Script:      args2.inp.Script,
					Description: args2.inp.Description,
				}, nil)
			},
		}, {
			name: "test_2, error not expected",
			args: args{
				c: context.Background(), inp: &domain.CreateCommandInput{
					Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
					Description: "combined script",
				},
			},
			want: &domain.CreateCommandOutput{
				Id:          "2",
				Script:      "echo \"The current directory is:\" && pwd && echo \"The user logged in is:\" && whoami",
				Description: "combined script",
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.inp.Script,
					Description: args2.inp.Description,
				}).Return(&entity.Command{
					Id:          2,
					Script:      args2.inp.Script,
					Description: args2.inp.Description,
				}, nil)
			},
		}, {
			name: "test_3, error not expected",
			args: args{
				c: context.Background(), inp: &domain.CreateCommandInput{
					Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
					Description: "combined script with sleeping",
				},
			},
			want: &domain.CreateCommandOutput{
				Id:          "3",
				Script:      "script\": \"echo \\\"Hello there!\\\" && sleep 8 && echo \\\"Oops! I fell asleep for a couple seconds!\\\"",
				Description: "combined script with sleeping",
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.inp.Script,
					Description: args2.inp.Description,
				}).Return(&entity.Command{
					Id:          3,
					Script:      args2.inp.Script,
					Description: args2.inp.Description,
				}, nil)
			},
		}, {
			name: "test_4, error not expected",
			args: args{
				c: context.Background(), inp: &domain.CreateCommandInput{
					Script:      "ls",
					Description: "simple script for look files",
				},
			},
			want: &domain.CreateCommandOutput{
				Id:          "4",
				Script:      "ls",
				Description: "simple script for look files",
			},
			prepare: func(args2 args, fields fields) {
				fields.Command.EXPECT().CreateCommand(context.Background(), &entity.Command{
					Script:      args2.inp.Script,
					Description: args2.inp.Description,
				}).Return(&entity.Command{
					Id:          4,
					Script:      args2.inp.Script,
					Description: args2.inp.Description,
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
			got, err := service.CreateCommand(tt.args.c, tt.args.inp)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			service.StopRunner()
			require.Equal(t, tt.want, got)
		})
	}
}
