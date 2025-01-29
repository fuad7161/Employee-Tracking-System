package handlers

import (
	db "github.com/MahediSabuj/go-teams/db/sqlc"
	"github.com/MahediSabuj/go-teams/internal/config"
	"github.com/MahediSabuj/go-teams/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetClients for all users to view their clients
func (m *Repository) GetClients(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("loggedIn") == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	//userId := session.Get("user_id")

	clients, err := m.DB.ListClients(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	updatedClients := make([]config.Client, len(clients))

	// change task created_at to human date
	for i, client := range clients {
		clientModel := config.Client{
			ID:         client.ID,
			ClientName: client.ClientName,
			Status:     client.Status,
			CreatedAt:  util.HumanDate(client.CreatedAt),
		}
		updatedClients[i] = clientModel
	}

	c.HTML(http.StatusOK, "client.page.gohtml", gin.H{
		"clients":  updatedClients,
		"loggedIn": session.Get("loggedIn"),
	})
}

// GetCreateClient for admin to create new client
func (m *Repository) GetCreateClient(c *gin.Context) {
	session := sessions.Default(c)
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	c.HTML(http.StatusOK, "create.client.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
	})
	return
}

// PostCreateClient for admin response to create new client
func (m *Repository) PostCreateClient(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	name := c.PostForm("name")
	status := c.PostForm("status")

	arg := db.CreateClientParams{
		ClientName: name,
		Status:     status,
	}

	_, err := m.DB.CreateClient(c, arg)
	if err != nil {
		log.Println("Error creating client: ", err)
		c.Redirect(http.StatusBadRequest, "/404")
		return
	}
	c.Redirect(http.StatusFound, "/success")

	return
}

func (m *Repository) DeleteClientByID(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid client ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Client ID must be a valid integer",
		})
		return
	}
	err = m.DB.DeleteClient(c, int64(idInt))
	if err != nil {
		log.Println("Error deleting client: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "This client has active project. Please delete all projects before deleting client",
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// EditClientPage for admin to edit client
func (m *Repository) EditClientPage(c *gin.Context) {
	session := sessions.Default(c)
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid client ID: ", err)
		c.Redirect(http.StatusTemporaryRedirect, "/404")
		return
	}

	client, err := m.DB.GetClientByID(c, int64(idInt))
	if err != nil {
		log.Println("Error getting client: ", err)
		c.Redirect(http.StatusTemporaryRedirect, "/404")
		return
	}

	clientModel := config.Client{
		ID:         client.ID,
		ClientName: client.ClientName,
		Status:     client.Status,
	}

	c.HTML(http.StatusOK, "edit-client.page.gohtml", gin.H{
		"client":   clientModel,
		"loggedIn": session.Get("loggedIn"),
	})
}

// UpdateClient for admin to update client
func (m *Repository) UpdateClient(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid client ID: ", err)
		c.Redirect(http.StatusTemporaryRedirect, "/404")
		return
	}

	name := c.PostForm("clientname")
	status := c.PostForm("status")

	arg := db.UpdateClientParams{
		ID:         int64(idInt),
		ClientName: name,
		Status:     status,
	}

	err = m.DB.UpdateClient(c, arg)
	if err != nil {
		log.Println("Error updating client: ", err)
		c.Redirect(http.StatusTemporaryRedirect, "/404")
		return
	}

	c.Redirect(http.StatusFound, "/clients")
}
