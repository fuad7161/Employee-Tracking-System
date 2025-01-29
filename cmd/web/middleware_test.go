package main

import (
	"github.com/MahediSabuj/go-teams/token"
	"github.com/stretchr/testify/mock"
)

// MockTokenMaker is a mock implementation of the token.Maker interface.
type MockTokenMaker struct {
	mock.Mock
}

func (m *MockTokenMaker) VerifyToken(token string) (*token.Payload, error) {
	return nil, nil
}

//
//func TestAuthMiddleware(t *testing.T) {
//	testCases := []struct {
//		name                     string
//		setupAuthHeader          func() string
//		buildTokenVerifierMock   func() token.Maker
//		expectedStatusCode       int
//		expectedPayloadInContext bool
//	}{
//		{
//			name: "ValidAuthorizationHeader",
//			setupAuthHeader: func() string {
//				return "bearer valid-token"
//			},
//			buildTokenVerifierMock: func() token.Maker {
//				mockMaker := &MockTokenMaker{}
//				mockMaker.On("VerifyToken", "valid-token").Return(&token.Payload{Username: "testuser"}, nil)
//				return nil
//			},
//			expectedStatusCode:       http.StatusOK,
//			expectedPayloadInContext: true,
//		},
//		{
//			name: "NoAuthorizationHeader",
//			setupAuthHeader: func() string {
//				return ""
//			},
//			buildTokenVerifierMock: func() token.Maker {
//				return nil
//			},
//			expectedStatusCode:       http.StatusUnauthorized,
//			expectedPayloadInContext: false,
//		},
//		{
//			name: "InvalidAuthorizationHeaderFormat",
//			setupAuthHeader: func() string {
//				return "invalidheader"
//			},
//			buildTokenVerifierMock: func() token.Maker {
//				return nil
//			},
//			expectedStatusCode:       http.StatusUnauthorized,
//			expectedPayloadInContext: false,
//		},
//		{
//			name: "UnsupportedAuthorizationType",
//			setupAuthHeader: func() string {
//				return "basic sometoken"
//			},
//			buildTokenVerifierMock: func() token.Maker {
//				return nil
//			},
//			expectedStatusCode:       http.StatusUnauthorized,
//			expectedPayloadInContext: false,
//		},
//		{
//			name: "InvalidToken",
//			setupAuthHeader: func() string {
//				return "bearer invalid-token"
//			},
//			buildTokenVerifierMock: func() token.Maker {
//				mockMaker := &MockTokenMaker{}
//				mockMaker.On("VerifyToken", "invalid-token").Return(nil, errors.New("invalid token"))
//				return nil
//			},
//			expectedStatusCode:       http.StatusUnauthorized,
//			expectedPayloadInContext: false,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			// Initialize Gin in test mode
//			gin.SetMode(gin.TestMode)
//			router := gin.New()
//
//			// Create the token verifier mock
//			tokenMaker := tc.buildTokenVerifierMock()
//
//			// Add the middleware to the router
//			router.Use(authMiddleware(tokenMaker))
//
//			// Define a test route
//			router.GET("/test", func(ctx *gin.Context) {
//				if payload, exists := ctx.Get(authorizationPayloadKey); exists {
//					ctx.JSON(http.StatusOK, gin.H{"payload": payload})
//				} else {
//					ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
//				}
//			})
//
//			// Prepare the test request
//			recorder := httptest.NewRecorder()
//			request, err := http.NewRequest(http.MethodGet, "/test", nil)
//			require.NoError(t, err)
//
//			// Set the Authorization header
//			authHeader := tc.setupAuthHeader()
//			if len(strings.TrimSpace(authHeader)) > 0 {
//				request.Header.Set(authorizationHeaderKey, authHeader)
//			}
//
//			// Perform the request
//			router.ServeHTTP(recorder, request)
//
//			// Check the response status code
//			require.Equal(t, tc.expectedStatusCode, recorder.Code)
//
//			// Verify if payload is set in the context properly
//			if tc.expectedPayloadInContext {
//				var response map[string]interface{}
//				require.Contains(t, response, "payload")
//			}
//		})
//	}
//}
