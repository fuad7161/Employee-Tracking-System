package handlers

import (
	db "github.com/MahediSabuj/go-teams/db/sqlc"
	"github.com/MahediSabuj/go-teams/internal/config"
	"github.com/MahediSabuj/go-teams/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// PostCreateProject for admin response to create new project
func (m *Repository) PostCreateProject(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	name := c.PostForm("name")
	clientID := c.PostForm("client_id")

	arg := db.CreateProjectParams{
		ProjectName: name,
		ClientID:    util.StringToInt(clientID),
	}

	_, err := m.DB.CreateProject(c, arg)
	if err != nil {
		log.Println("Error creating project: ", err)
		c.Redirect(http.StatusBadRequest, "/404")
		return
	}
	c.Redirect(http.StatusFound, "/success")
	return
}

// GetCreateProject for admin to create new project
func (m *Repository) GetCreateProject(c *gin.Context) {
	session := sessions.Default(c)
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// Fetch clients from the database
	clients, err := m.DB.ListClients(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch clients"})
		return
	}

	// Convert client list to id-name format
	clientList := make(map[int]interface{})
	for _, client := range clients {
		clientList[int(client.ID)] = client.ClientName
	}

	// Output the client list to the rendering context
	c.HTML(http.StatusOK, "create.project.page.gohtml", gin.H{
		"loggedIn":   session.Get("loggedIn"),
		"clientList": clientList,
	})
	return
}

// GetProjects for all users to view their projects
func (m *Repository) GetProjects(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("loggedIn") == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	userId := session.Get("user_id")
	log.Println("userId: ", userId)

	projects, err := m.DB.ListProjects(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	updatedProjects := make([]config.Project, len(projects))

	// change task created_at to human date
	for i, project := range projects {
		projectModel := config.Project{
			ID:          project.ID,
			ProjectName: project.ProjectName,
			ClientID:    project.ClientID,
			CreatedAt:   util.HumanDate(project.CreatedAt),
		}
		updatedProjects[i] = projectModel
	}

	clients, err := m.DB.ListClients(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch clients",
		})
		return
	}

	clientlst := make(map[int64]string)

	for _, client := range clients {
		clientlst[client.ID] = client.ClientName
	}

	c.HTML(http.StatusOK, "project.page.gohtml", gin.H{
		"projects":  updatedProjects,
		"loggedIn":  session.Get("loggedIn"),
		"clientlst": clientlst,
	})
}

// DeleteProjectByID allows admins to delete a project by its ID
func (m *Repository) DeleteProjectByID(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	projectID := c.Param("id")

	err := m.DB.DeleteProject(c, util.StringToInt(projectID))
	if err != nil {
		log.Println("Error deleting project: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to delete project"})
		return
	}

	c.JSON(http.StatusOK, true)
}

// EditProjectPage allows admins to edit a project by its ID
func (m *Repository) EditProjectPage(c *gin.Context) {
	session := sessions.Default(c)
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	projectID := c.Param("id")
	project, err := m.DB.GetProjectByID(c, util.StringToInt(projectID))
	if err != nil {
		log.Println("Error fetching project: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch project"})
		return
	}

	clients, err := m.DB.ListClients(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch clients"})
		return
	}

	clientList := make(map[int64]interface{})
	for _, client := range clients {
		clientList[client.ID] = client.ClientName
	}

	c.HTML(http.StatusOK, "edit-project.page.gohtml", gin.H{
		"project":    project,
		"clientList": clientList,
		"loggedIn":   session.Get("loggedIn"),
	})
}

// UpdateProject allows admins to update a project by its ID
func (m *Repository) UpdateProject(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	projectID := c.Param("id")
	name := c.PostForm("project")
	clientID := c.PostForm("client_id")

	arg := db.UpdateProjectByIDParams{
		ID:          util.StringToInt(projectID),
		ProjectName: name,
		ClientID:    util.StringToInt(clientID),
	}

	err := m.DB.UpdateProjectByID(c, arg)
	if err != nil {
		log.Println("Error updating project: ", err)
		c.Redirect(http.StatusBadRequest, "/404")
		return
	}

	c.Redirect(http.StatusFound, "/projects")
}
