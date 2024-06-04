package v1

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock "github.com/zsandibe/command-runner-service/internal/delivery/http/v1/mocks"
	"github.com/zsandibe/command-runner-service/internal/domain"
)

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)

	return ctx
}

func TestHandler_CreateCommand(t *testing.T) {
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
				res:    "{\"id\":\"1\",\"script\":\"ls\",\"description\":\"look files\"}",
				status: http.StatusOK,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/command", strings.NewReader(`{"script": "ls", "description": "look files"}`))
				c.Request.Header.Set("Content-Type", "application/json")
				fields.Service.EXPECT().CreateCommand(c.Request.Context(), &domain.CreateCommandInput{
					Script:      "ls",
					Description: "look files",
				}).Return(&domain.CreateCommandOutput{
					Id:          "1",
					Script:      "ls",
					Description: "look files",
				}, nil)
			},
		}, {
			name:    "test_2, StatusBadRequest",
			wantErr: false,
			want: want{
				res:    "{\"error\":\"script must not be empty\"}",
				status: http.StatusBadRequest,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/command", strings.NewReader(`{"script": "", "description": "look files"}`))
				c.Request.Header.Set("Content-Type", "application/json")
			},
		}, {
			name:    "test_3, StatusBadRequest",
			wantErr: false,
			want: want{
				res:    "{\"error\":\"description for script must not be empty\"}",
				status: http.StatusBadRequest,
			},
			prepare: func(fields fields, recorder *httptest.ResponseRecorder, c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodPost, "/command", strings.NewReader(`{"script": "ls", "description": ""}`))
				c.Request.Header.Set("Content-Type", "application/json")
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

			handler.CreateCommand(c)

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
