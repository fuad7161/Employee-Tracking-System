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
	"log"
	"net/http"
)

// GetSbu for all users to view their sbu
func (m *Repository) GetSbu(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("loggedIn") == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	sbu, err := m.DB.ListSBUs(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	updatedSbu := make([]config.SBU, len(sbu))
	userName := make(map[int]interface{})
	for i, sbu := range sbu {
		user, err := m.DB.GetUserByID(c, sbu.SbuHeadUserID.Int64)
		if err != nil {
			log.Println("Error fetching user: ", err)
			c.Redirect(http.StatusBadRequest, "/404")
			return
		}
		userName[int(sbu.SbuHeadUserID.Int64)] = user.Firstname + " " + user.Lastname
		fmt.Println(user.Firstname, user.Lastname)
		sbuModel := config.SBU{
			ID:            sbu.ID,
			SbuName:       sbu.SbuName.String,
			SbuHeadUserID: user.ID,
			CreatedAt:     util.HumanDate(sbu.CreatedAt),
		}
		updatedSbu[i] = sbuModel
	}

	fmt.Println(updatedSbu)
	fmt.Println(userName)
	c.HTML(http.StatusOK, "sbu.page.gohtml", gin.H{
		"sbu":      updatedSbu,
		"loggedIn": session.Get("loggedIn"),
		"userName": userName,
	})
}

// GetCreateSBU for admin to create new sbu
func (m *Repository) GetCreateSBU(c *gin.Context) {
	session := sessions.Default(c)
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	users, err := m.DB.ListUsers(c)
	if err != nil {
		log.Println("Error fetching users: ", err)
		c.Redirect(http.StatusBadRequest, "/404")
		return
	}
	usersMap := make(map[int]interface{})
	for _, user := range users {
		usersMap[int(user.ID)] = user.Firstname + " " + user.Lastname
	}
	c.HTML(http.StatusOK, "create.sbu.page.gohtml", gin.H{
		"loggedIn": session.Get("loggedIn"),
		"Users":    usersMap,
	})
	return
}

// PostCreateSBU for admin response to create new sbu
func (m *Repository) PostCreateSBU(c *gin.Context) {
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	name := c.PostForm("name")
	headUserID := util.StringToInt(c.PostForm("sbu_head_user_id"))

	arg := db.CreateSBUParams{
		SbuName: pgtype.Text{
			String: name,
			Valid:  true,
		},
		SbuHeadUserID: pgtype.Int8{
			Int64: headUserID,
			Valid: true,
		},
	}

	sbu, err := m.DB.CreateSBU(c, arg)

	// Retrieve the user by ID from the database
	user, err := m.DB.GetUserByID(c, headUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user by ID"})
		return
	}

	// Parse form data
	userID := user.ID
	firstname := user.Firstname
	lastname := user.Lastname
	email := user.Email
	userRoleID := user.UserRoleID
	sbuID := sbu.ID
	if err := m.DB.UpdateUserInformation(c, db.UpdateUserInformationParams{
		Firstname:  firstname,
		Lastname:   lastname,
		Email:      email,
		UserRoleID: userRoleID,
		SbuID:      pgtype.Int8(sql.NullInt64{Int64: int64(sbuID), Valid: true}),
		ID:         userID,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user information"})
		return
	}

	if err != nil {
		log.Println("Error creating SBU: ", err)
		c.Redirect(http.StatusBadRequest, "/404")
		return
	}

	c.Redirect(http.StatusFound, "/success")
	return
}

// EditSBUPage for admin to edit sbu
func (m *Repository) EditSBUPage(c *gin.Context) {

	session := sessions.Default(c)
	if isAdmin(c) == false {
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// Get SBU ID from the request
	sbuID := util.StringToInt(c.Param("id"))

	// Fetch the SBU from the database
	sbu, err := m.DB.GetSBU(c, sbuID)
	if err != nil {
		log.Println("Error fetching SBU: ", err)
		c.Redirect(http.StatusNotFound, "/404")
		return
	}
	// Fetch SBU head user details
	user, err := m.DB.GetUserByID(c, sbu.SbuHeadUserID.Int64)
	if err != nil {
		log.Println("Error fetching user: ", err)
		c.Redirect(http.StatusNotFound, "/404")
		return
	}

	// Fetch SBU head user details
	users, err := m.DB.ListUsers(c)
	if err != nil {
		log.Println("Error fetching user: ", err)
		c.Redirect(http.StatusNotFound, "/404")
		return
	}

	userlst := make(map[int64]interface{})
	for _, user := range users {
		userlst[user.ID] = user.Firstname + " " + user.Lastname
	}

	// Serve the response
	c.HTML(http.StatusOK, "edit-sbu.page.gohtml", gin.H{
		"subName":  sbu.SbuName.String,
		"sbuHead":  user.Firstname + " " + user.Lastname,
		"sbu":      sbu,
		"userlst":  userlst,
		"loggedIn": session.Get("loggedIn"),
	})
	return
}

func (m *Repository) UpdateSBUPagePost(c *gin.Context) {
	sbuId := util.StringToInt(c.Param("id"))
	sbuName := c.PostForm("sbuName")
	subHeadUserID := pgtype.Int8{Int64: util.StringToInt(c.PostForm("sbu_head_id")), Valid: true}
	m.DB.UpdateSbuByID(c, db.UpdateSbuByIDParams{
		ID:            sbuId,
		SbuHeadUserID: subHeadUserID,
		SbuName:       pgtype.Text{String: sbuName, Valid: true},
	})

	// Retrieve the user by ID from the database
	user, err := m.DB.GetUserByID(c, util.StringToInt(c.PostForm("sbu_head_id")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user by ID"})
		return
	}

	// Parse form data
	userID := user.ID
	firstname := user.Firstname
	lastname := user.Lastname
	email := user.Email
	userRoleID := user.UserRoleID
	sbuID := sbuId
	if err := m.DB.UpdateUserInformation(c, db.UpdateUserInformationParams{
		Firstname:  firstname,
		Lastname:   lastname,
		Email:      email,
		UserRoleID: userRoleID,
		SbuID:      pgtype.Int8(sql.NullInt64{Int64: int64(sbuID), Valid: true}),
		ID:         userID,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user information"})
		return
	}

	if err != nil {
		log.Println("Error creating SBU: ", err)
		c.Redirect(http.StatusBadRequest, "/404")
		return
	}

	c.Redirect(http.StatusFound, "/admin/sbu")
}

func (m *Repository) DeleteSBU(c *gin.Context) {
	id := util.StringToInt(c.Param("id"))

	err := m.DB.DeleteSBU(c, id)
	if err != nil {
		log.Println("Error deleting SBU: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete SBU"})
		return
	}

	c.JSON(http.StatusOK, true)
}
