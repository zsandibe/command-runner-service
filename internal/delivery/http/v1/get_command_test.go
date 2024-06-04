package v1

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock "github.com/zsandibe/command-runner-service/internal/delivery/http/v1/mocks"
	"github.com/zsandibe/command-runner-service/internal/entity"
)

func TestHandler_GetCommandById(t *testing.T) {
	type want struct {
		res    string
		status int
	}
	type fields struct {
		Service *mock.MockService
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
		want    want
		prepare func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context)
	}{
		{
			name:    "test_1, StatusOK",
			wantErr: false,
			want: want{
				res:    "{\"id\":1,\"script\":\"ls\",\"description\":\"look files\"}",
				status: http.StatusOK,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				route := "/command/{id}"
				value := "1"
				c.Request = httptest.NewRequest(http.MethodGet, route, nil)
				c.Params = []gin.Param{{
					Key:   "id",
					Value: value,
				}}
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().GetCommandById(c.Request.Context(), value).Return(&entity.Command{
					Id:          1,
					Script:      "ls",
					Description: "look files",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Service: mock.NewMockService(ctrl),
			}
			handler := NewHandler(f.Service)

			w := httptest.NewRecorder()
			c := GetTestGinContext(w)
			tt.prepare(f, w, c)
			handler.GetCommandById(c)

			resp := w.Result()

			body, err := ioutil.ReadAll(resp.Body)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want.status, resp.StatusCode)
			require.Equal(t, tt.want.res, string(body))

		})
	}
}
