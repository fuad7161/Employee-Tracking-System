package main

import (
	"github.com/MahediSabuj/go-teams/internal/handlers"
	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) {
	// Define routes for CRUD operations
	r.LoadHTMLGlob("../../template/*")
	r.Static("/static", "../../static")
	r.GET("/", handlers.Repo.Home)
	r.POST("/users", handlers.Repo.CreateUser)

	r.GET("/login", handlers.Repo.GetLogin)
	r.POST("/login", handlers.Repo.Login)
	r.GET("/logout", handlers.Repo.LogoutHandler)

	r.GET("/registration", handlers.Repo.GetRegistrationPage)
	r.GET("/profile", handlers.Repo.GetProfile)

	r.GET("/about-us", handlers.Repo.GetAboutPage)
	r.GET("/contact-us", handlers.Repo.GetContactPage)
	r.GET("/tasks", handlers.Repo.GetTaskPage)

	r.GET("/admin/dashboard", handlers.Repo.GetAdminPage)
	r.GET("/admin/users", handlers.Repo.GetUsersPage)
	r.GET("/admin/users/update/:id", handlers.Repo.EditUserPage)
	r.POST("/admin/update-profile/:id", handlers.Repo.UpdateUserProfile)
	r.DELETE("/admin/users/delete/:id", handlers.Repo.DeleteUser)

	r.GET("/create/client", handlers.Repo.GetCreateClient)
	r.POST("/create/client", handlers.Repo.PostCreateClient)
	r.GET("/admin/client/update/:id", handlers.Repo.EditClientPage)
	r.POST("/admin/client/update/:id", handlers.Repo.UpdateClient)
	r.DELETE("/admin/client/delete/:id", handlers.Repo.DeleteClientByID)

	r.GET("/create/project", handlers.Repo.GetCreateProject)
	r.POST("/create/project", handlers.Repo.PostCreateProject)
	r.GET("/admin/project/update/:id", handlers.Repo.EditProjectPage)
	r.POST("/admin/project/update/:id", handlers.Repo.UpdateProject)
	r.DELETE("/admin/project/delete/:id", handlers.Repo.DeleteProjectByID)

	r.GET("/create/task", handlers.Repo.GetCreateTask)
	r.POST("/create/task", handlers.Repo.PostCreateTask)
	r.DELETE("/admin/task/delete/:id", handlers.Repo.DeleteTask)
	r.GET("/admin/task/update/:id", handlers.Repo.UpdateTaskPage)
	r.POST("/admin/update-task/:id", handlers.Repo.PostUpdateTask)

	r.GET("/create/sbu", handlers.Repo.GetCreateSBU)
	r.POST("/create/sbu", handlers.Repo.PostCreateSBU)
	r.GET("/sbu/update/:id", handlers.Repo.EditSBUPage)
	r.POST("/sbu/update/:id", handlers.Repo.UpdateSBUPagePost)
	r.DELETE("/admin/sbu/delete/:id", handlers.Repo.DeleteSBU)

	r.GET("/projects", handlers.Repo.GetProjects)

	r.GET("/clients", handlers.Repo.GetClients)

	r.GET("/success", handlers.Repo.GetSuccess)

	r.GET("/admin/sbu", handlers.Repo.GetSbu)

	return
}
