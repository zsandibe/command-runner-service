package v1

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock "github.com/zsandibe/command-runner-service/internal/delivery/http/v1/mocks"
)

func TestHandler_StopCommandById(t *testing.T) {
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
				res:    "{\"message\":\"command stopped successfully\"}",
				status: http.StatusOK,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				route := "/command/{id}"
				c.Request = httptest.NewRequest(http.MethodPost, route, nil)
				c.AddParam("id", "1")
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().StopCommandById("1").Return(nil)
			},
		}, {
			name:    "test_2, StatusInternalServerError",
			wantErr: false,
			want: want{
				res:    "{\"error\":\"not found value\"}",
				status: http.StatusInternalServerError,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				route := "/command/{id}"
				c.Request = httptest.NewRequest(http.MethodPost, route, nil)
				c.AddParam("id", "1")
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().StopCommandById("1").Return(errors.New("not found value"))
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
			handler.StopCommandById(c)

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
