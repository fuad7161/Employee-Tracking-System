package handlers

import (
	db "github.com/MahediSabuj/go-teams/db/sqlc"
	"github.com/MahediSabuj/go-teams/token"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	DB         *db.Queries
	tokenMaker token.Maker
}

func NewRepo(db *db.Queries, tokenMaker token.Maker) *Repository {
	return &Repository{
		DB:         db,
		tokenMaker: tokenMaker,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handles, home page UI from html template
func (m *Repository) Home(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(http.StatusOK, "home.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
		"admin":    session.Get("admin"),
	})
}

// GetLogin handles login  UI from html template
func (m *Repository) GetLogin(c *gin.Context) {
	session := sessions.Default(c)
	loggedIn := session.Get("loggedIn")
	admin := session.Get("admin")
	if loggedIn != nil || admin != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.HTML(http.StatusOK, "login.page.gohtml", nil)
}

// Login handles post login authentication and authorization
func (m *Repository) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Fetch user by email
	user, err := m.DB.GetUserByEmail(c, email)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.page.gohtml", gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Handle mismatched or other bcrypt errors
		c.HTML(http.StatusUnauthorized, "login.page.gohtml", gin.H{"error": "Invalid credentials"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()
	// TODO: set user role
	// Set session data
	session.Set("loggedIn", true)
	err = session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.page.gohtml", gin.H{"error": "Could not save session"})
		return
	}

	// Render the home page for logged-in users
	c.HTML(http.StatusOK, "home.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
		"msg":      "Welcome!",
	})

}

func (m *Repository) LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "home.page.gohtml", gin.H{"error": "Failed to save session"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func (m *Repository) GetRegistrationPage(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("loggedIn") != nil || session.Get("admin") != nil {
		c.Redirect(http.StatusFound, "/")
	}
	c.HTML(http.StatusOK, "registration.page.gohtml", nil)
}

func (m *Repository) GetProfile(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")

	// Fetch user by ID
	user, err := m.DB.GetUserByID(c, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "profile.page.gohtml", gin.H{
		"user":     user,
		"loggedIn": session.Get("loggedIn"),
	})
}

func (m *Repository) GetAboutPage(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(http.StatusOK, "about.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
	})
}

func (m *Repository) GetAdminPage(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("loggedIn") == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "admin.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
	})
}

// GetSuccess for success page
func (m *Repository) GetSuccess(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(http.StatusOK, "success.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
	})
	return
}

// GetContactPage for contact page
func (m *Repository) GetContactPage(c *gin.Context) {
	session := sessions.Default(c)
	c.HTML(http.StatusOK, "contact.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
	})
	return
}
