package handlers

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	db "github.com/MahediSabuj/go-teams/db/sqlc"
	"github.com/MahediSabuj/go-teams/token"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewRepo(t *testing.T) {
	mockDB := &db.Queries{}
	mockTokenMaker, _ := token.NewJWTMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	mockSession := scs.New()
	repo := NewRepo(mockDB, mockTokenMaker, mockSession)

	assert.NotNil(t, repo)
	assert.Equal(t, mockDB, repo.DB)
	assert.Equal(t, mockTokenMaker, repo.tokenMaker)
	assert.Equal(t, mockSession, repo.Session)
}

func TestNewHandlers(t *testing.T) {
	mockDB := &db.Queries{}
	mockTokenMaker, _ := token.NewJWTMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	mockSession := scs.New()
	repo := NewRepo(mockDB, mockTokenMaker, mockSession)
	NewHandlers(repo)

	assert.NotNil(t, Repo)
	assert.Equal(t, mockDB, Repo.DB)
	assert.Equal(t, mockTokenMaker, Repo.tokenMaker)
	assert.Equal(t, mockSession, Repo.Session)
}

func TestHome(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router and apply session middleware
	store := memstore.NewStore([]byte("secret"))
	r := gin.New()
	r.Use(sessions.Sessions("test-session", store))

	// Define and parse the HTML template
	tmpl := template.Must(template.New("home.page.gohtml").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <title>Home</title>
  </head>
  <body>
    {{if .loggedIn}}Logged in: {{.loggedIn}}{{end}}
    {{if .admin}}Admin: {{.admin}}{{else}}Admin: false{{end}}
  </body>
</html>`))
	r.SetHTMLTemplate(tmpl)

	// Register the Home handler
	repo := &Repository{}
	r.GET("/", func(c *gin.Context) {
		// Set up session values for this test route
		session := sessions.Default(c)
		session.Set("loggedIn", true)
		session.Set("admin", false)
		session.Save()

		// Call the actual handler
		repo.Home(c)
	})

	// Create a test HTTP request
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)

	// Create a response recorder to capture the output
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logged in: true")
	assert.Contains(t, w.Body.String(), "Admin: false")
}
