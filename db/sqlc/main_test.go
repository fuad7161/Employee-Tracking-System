package db

import (
	"context"
	"github.com/MahediSabuj/go-teams/token"
	"github.com/alexedwards/scs/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connPool, err := pgxpool.New(context.Background(), "postgresql://postgres:@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	database := New(connPool)

	//set up session
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	// set up token maker
	tokenMaker, _ := token.NewJWTMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))

	repo := NewRepo(database, tokenMaker, session)
	NewHandlers(repo)

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	code := m.Run()

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

	os.Exit(code)
}
