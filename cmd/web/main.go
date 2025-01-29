package main

import (
	"context"
	db "github.com/MahediSabuj/go-teams/db/sqlc"
	"github.com/MahediSabuj/go-teams/internal/handlers"
	"github.com/MahediSabuj/go-teams/token"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connPool, err := pgxpool.New(context.Background(), "postgresql://postgres:@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	database := db.New(connPool)
	// set up token maker
	tokenMaker, _ := token.NewJWTMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))

	repo := handlers.NewRepo(database, tokenMaker)
	handlers.NewHandlers(repo)

	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   60 * 15,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	r.Use(sessions.Sessions("mysession", store))
	routes(r)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
