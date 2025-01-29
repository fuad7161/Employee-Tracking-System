package handlers

import (
	"database/sql"
	"fmt"
	db "github.com/MahediSabuj/go-teams/db/sqlc"
	"github.com/MahediSabuj/go-teams/internal/config"
	"github.com/MahediSabuj/go-teams/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func (m *Repository) CreateUser(c *gin.Context) {
	password := c.PostForm("password")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	arg := db.CreateUserParams{
		Firstname: c.PostForm("firstname"),
		Lastname:  c.PostForm("lastname"),
		Email:     c.PostForm("email"),
		Password:  string(hashedPassword),
	}
	_, err := m.DB.CreateUser(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// Step 4: Respond with the created user
	c.Redirect(http.StatusFound, "/login")
}

func (m *Repository) GetUsersPage(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("loggedIn") == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	// Fetch users from the database
	users, err := m.DB.ListUsers(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to fetch users"})
		return
	}

	updatedUsers := make([]config.User, len(users))

	// change task created_at to human date
	for i, user := range users {
		userModel := config.User{
			ID:         user.ID,
			Firstname:  user.Firstname,
			Lastname:   user.Lastname,
			Email:      user.Email,
			UserRoleID: user.UserRoleID.Int64,
			SbuID:      user.SbuID.Int64,
			CreatedAt:  util.HumanDate(user.CreatedAt),
		}
		updatedUsers[i] = userModel
	}

	fmt.Println(updatedUsers)

	// Fetch SBU list from the database
	sbus, err := m.DB.ListSBUs(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to fetch SBUs"})
		return
	}

	for i := range sbus {
		fmt.Println(sbus[i].SbuName.String)
	}

	//Convert SBU list to id-name format
	sbuList := make(map[int]interface{})

	for _, sbu := range sbus {
		sbuList[int(sbu.ID)] = sbu.SbuName.String
	}
	fmt.Println(sbuList)

	// Fetch user roles from the database
	userRoles, err := m.DB.ListUserRole(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to fetch user roles"})
		return
	}
	userRoleslst := make(map[int]interface{})
	for i := range userRoles {
		userRoleslst[int(userRoles[i].ID)] = userRoles[i].RoleName
	}
	fmt.Println(userRoleslst)

	// Pass SBU and user role lists to the rendering context
	c.HTML(http.StatusOK, "users.page.gohtml", gin.H{
		"users":     updatedUsers,
		"loggedIn":  session.Get("loggedIn"),
		"sbuList":   sbuList,
		"userRoles": userRoleslst,
	})
}

// EditUserPage to update all information user id
func (m *Repository) EditUserPage(c *gin.Context) {

	id := c.Param("id")
	user, err := m.DB.GetUserByID(c, util.StringToInt(id))
	sbu, err := m.DB.ListSBUs(c)
	roles, err := m.DB.ListUserRole(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user roles"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user information"})
		return
	}

	session := sessions.Default(c)
	c.HTML(http.StatusOK, "edit-user.page.gohtml", gin.H{
		"loggedIn":     session.Get("loggedIn"),
		"user":         user,
		"currentSbu":   user.SbuID,
		"currentRoles": user.UserRoleID,
		"Sbus":         sbu,
		"Roles":        roles,
	})
}

func (m *Repository) UpdateUserProfile(c *gin.Context) {
	// Parse form data
	userID := c.Param("id")
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	email := c.PostForm("email")
	userRoleID := c.PostForm("user_role_id")
	sbuID := c.PostForm("sbu_id")

	userIDInt, err := strconv.Atoi(userID)
	userRoleIDInt, err := strconv.Atoi(userRoleID)
	sbuIDInt, err := strconv.Atoi(sbuID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	fmt.Println(firstname, lastname, email, userRoleIDInt, sbuIDInt, userIDInt)
	if err := m.DB.UpdateUserInformation(c, db.UpdateUserInformationParams{
		Firstname:  firstname,
		Lastname:   lastname,
		Email:      email,
		UserRoleID: pgtype.Int8(sql.NullInt64{Int64: int64(userRoleIDInt), Valid: true}),
		SbuID:      pgtype.Int8(sql.NullInt64{Int64: int64(sbuIDInt), Valid: true}),
		ID:         int64(userIDInt),
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user information"})
		return
	}
	//session := sessions.Default(c)
	//c.HTML(200, "users.page.gohtml", gin.H{
	//	"loggedIn": session.Get("loggedIn"),
	//})
	c.Redirect(http.StatusFound, "/admin/users")
}

func (m *Repository) DeleteUser(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	id := c.Param("id")
	if err := m.DB.DeleteUser(c, util.StringToInt(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, true)
}

func (m *Repository) UpdateUserSBUById(c *gin.Context, id int, sbuId int) error {

	// Retrieve the user by ID from the database
	user, err := m.DB.GetUserByID(c, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user by ID"})
		return err
	}

	// Respond with the retrieved user
	c.JSON(http.StatusOK, gin.H{"user": user})
	// Parse form data
	userID := user.ID
	firstname := user.Firstname
	lastname := user.Lastname
	email := user.Email
	userRoleID := user.UserRoleID
	sbuID := sbuId

	fmt.Println(firstname, lastname, email, userRoleID, sbuID, userID)
	if err := m.DB.UpdateUserInformation(c, db.UpdateUserInformationParams{
		Firstname:  firstname,
		Lastname:   lastname,
		Email:      email,
		UserRoleID: userRoleID,
		SbuID:      pgtype.Int8(sql.NullInt64{Int64: int64(sbuID), Valid: true}),
		ID:         userID,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user information"})
		return err
	}
	return nil
}
