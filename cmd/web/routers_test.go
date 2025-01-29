package main

import (
	db "github.com/MahediSabuj/go-teams/db/sqlc"
	"github.com/MahediSabuj/go-teams/token"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock handler methods to test routes
type MockRepo struct {
	DB         *db.Queries
	tokenMaker token.Maker
	Session    *scs.SessionManager
}

func (m *MockRepo) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Home"})
}
func (m *MockRepo) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateUser"})
}
func (m *MockRepo) GetLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetLogin"})
}
func (m *MockRepo) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login"})
}
func (m *MockRepo) LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout"})
}
func (m *MockRepo) GetRegistrationPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetRegistrationPage"})
}
func (m *MockRepo) GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetProfile"})
}
func (m *MockRepo) GetAboutPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAboutPage"})
}
func (m *MockRepo) GetTaskPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetTaskPage"})
}
func (m *MockRepo) GetAdminPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAdminPage"})
}
func (m *MockRepo) GetUsersPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUsersPage"})
}
func (m *MockRepo) GetCreateClient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetCreateClient"})
}
func (m *MockRepo) PostCreateClient(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PostCreateClient"})
}
func (m *MockRepo) GetCreateProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetCreateProject"})
}
func (m *MockRepo) PostCreateProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "PostCreateProject"})
}
func (m *MockRepo) GetSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetSuccess"})
}

func TestRoutes(t *testing.T) {
	// Initialize the Gin router
	r := gin.Default()

	// Mock the handlers

	// Set up the routes
	routes(r)

	// Test the routes
	tests := []struct {
		method       string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/", http.StatusOK, `{"message":"Home"}`},
		{"POST", "/users", http.StatusOK, `{"message":"CreateUser"}`},
		{"GET", "/login", http.StatusOK, `{"message":"GetLogin"}`},
		{"POST", "/login", http.StatusOK, `{"message":"Login"}`},
		{"GET", "/logout", http.StatusOK, `{"message":"Logout"}`},
		{"GET", "/registration", http.StatusOK, `{"message":"GetRegistrationPage"}`},
		{"GET", "/profile", http.StatusOK, `{"message":"GetProfile"}`},
		{"GET", "/about-us", http.StatusOK, `{"message":"GetAboutPage"}`},
		{"GET", "/tasks", http.StatusOK, `{"message":"GetTaskPage"}`},
		{"GET", "/admin/dashboard", http.StatusOK, `{"message":"GetAdminPage"}`},
		{"GET", "/admin/users", http.StatusOK, `{"message":"GetUsersPage"}`},
		{"GET", "/create/client", http.StatusOK, `{"message":"GetCreateClient"}`},
		{"POST", "/create/client", http.StatusOK, `{"message":"PostCreateClient"}`},
		{"GET", "/create/project", http.StatusOK, `{"message":"GetCreateProject"}`},
		{"POST", "/create/project", http.StatusOK, `{"message":"PostCreateProject"}`},
		{"GET", "/success", http.StatusOK, `{"message":"GetSuccess"}`},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			// Create a new HTTP request
			req, _ := http.NewRequest(tt.method, tt.url, nil)
			w := httptest.NewRecorder()

			// Serve the HTTP request
			r.ServeHTTP(w, req)
		})
	}
}
