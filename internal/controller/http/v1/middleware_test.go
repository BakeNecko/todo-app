package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
	"todo-app/internal/usecase/mocks"
)

func TestHandlers_authMiddleware(t *testing.T) {
	type mockBehavior func(s *mocks.MockITokenManager, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mocks.MockITokenManager, token string) {
				s.EXPECT().Parse(token).Return("1", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "No Header",
			headerName:           "",
			mockBehavior:         func(s *mocks.MockITokenManager, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"empty auth header"}`,
		},
		{
			name:                 "Invalid Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			mockBehavior:         func(s *mocks.MockITokenManager, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"error":"token is empty"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)

			tokenManager := mocks.NewMockITokenManager(c)
			testCase.mockBehavior(tokenManager, testCase.token)

			// Test server
			r := gin.New()

			deps := &Deps{
				UseCases:        nil,
				JwtTokenManager: tokenManager,
			}

			onlyAuthAPI := r.Group("/", deps.authMiddleware)
			onlyAuthAPI.GET("/protected", func(context *gin.Context) {
				id, _ := context.Get(userCtx)
				if id != nil {
					context.String(200, id.(string))
				}
			})

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			// Make Request
			r.ServeHTTP(w, req)

			// Asserts
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
