package handlers

import (
	"fmt"
	db "github.com/MahediSabuj/go-teams/db/sqlc"
	"github.com/MahediSabuj/go-teams/internal/config"
	"github.com/MahediSabuj/go-teams/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (m *Repository) GetTaskPage(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("loggedIn") == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	userId := session.Get("user_id")
	log.Println("userId: ", userId)

	tasks, err := m.DB.ListTasksByUserID(c, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	updatedTasks := make([]config.Task, len(tasks))

	// change task created_at to human date
	for i, task := range tasks {
		taskModel := config.Task{
			ID:             task.ID,
			TaskTitle:      task.TaskTitle,
			Progress:       task.Progress,
			ProjectID:      task.ProjectID,
			AssignedUserID: task.AssignedUserID,
			CreatedAt:      util.HumanDate(task.CreatedAt),
		}
		updatedTasks[i] = taskModel
	}

	// get all projects
	projects, err := m.DB.ListProjects(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch projects"})
		return
	}

	projectlst := make(map[int]interface{})
	for _, project := range projects {
		projectlst[int(project.ID)] = project.ProjectName
	}
	c.HTML(http.StatusOK, "task.page.gohtml", gin.H{
		"tasks":      updatedTasks,
		"loggedIn":   session.Get("loggedIn"),
		"projectlst": projectlst,
	})
}

// GetCreateTask for admin to create new task
func (m *Repository) GetCreateTask(c *gin.Context) {
	session := sessions.Default(c)
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	projects, err := m.DB.ListProjects(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch projects"})
		return
	}

	users, err := m.DB.ListUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}

	c.HTML(http.StatusOK, "create.task.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
		"projects": projects,
		"users":    users,
	})
	c.HTML(http.StatusOK, "create.task.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
	})
	return
}

// PostCreateTask for admin response to create new task
func (m *Repository) PostCreateTask(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	taskTile := c.PostForm("title")
	progress := c.PostForm("progress")
	projectId := c.PostForm("project_id")
	userId := c.PostForm("user_id")

	arg := db.CreateTaskParams{
		TaskTitle:      taskTile,
		Progress:       util.StringToInt(progress),
		ProjectID:      util.StringToInt(projectId),
		AssignedUserID: util.StringToInt(userId),
	}

	_, err := m.DB.CreateTask(c, arg)
	if err != nil {
		log.Println("Error creating task: ", err)
		c.Redirect(http.StatusBadRequest, "/404")
		return
	}
	c.Redirect(http.StatusFound, "/success")
	return
}

func isAdmin(c *gin.Context) bool {
	session := sessions.Default(c)
	if session.Get("user_id") == nil {
		return false
	}
	userId := session.Get("user_id")
	user, err := Repo.DB.GetUserByID(c, userId.(int64))
	if err != nil {
		return false
	}
	userRoleID := user.UserRoleID.Int64

	// get role from roles table using userRoleID
	role, err := Repo.DB.GetUserRoleByID(c, userRoleID)
	if err != nil {
		return false
	}

	if role.RoleName == "admin" {
		return true
	}

	return false
}

func (m *Repository) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := m.DB.DeleteTask(c, util.StringToInt(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}
	c.JSON(http.StatusOK, true)
	return
}

func (m *Repository) UpdateTaskPage(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := m.DB.GetTaskByID(c, int64(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		return
	}
	session := sessions.Default(c)

	projectID := task.ProjectID // Assuming `task` contains the project ID
	project, err := m.DB.GetProjectByID(c, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch project"})
		return
	}
	fmt.Println(task)
	c.HTML(http.StatusOK, "edit-task.page.gohtml", gin.H{
		"loggedIn":    session.Get("loggedIn"),
		"task":        task,
		"projectName": project.ProjectName,
	})

}

func (m *Repository) PostUpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	m.DB.UpdateTaskTile(c, db.UpdateTaskTileParams{
		ID:        int64(id),
		TaskTitle: c.PostForm("task_title"),
	})
	m.DB.UpdateTaskProgress(c, db.UpdateTaskProgressParams{
		ID:       int64(id),
		Progress: util.StringToInt(c.PostForm("progress")),
	})
	c.Redirect(http.StatusFound, "/tasks")
}
