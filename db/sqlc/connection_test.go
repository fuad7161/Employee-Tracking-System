package db

import (
	"github.com/MahediSabuj/go-teams/token"
	"github.com/alexedwards/scs/v2"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	DB         *Queries
	TokenMaker token.Maker
	Session    *scs.SessionManager
}

func NewRepo(db *Queries, tokenMaker token.Maker, session *scs.SessionManager) *Repository {
	return &Repository{
		DB:         db,
		TokenMaker: tokenMaker,
		Session:    session,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
