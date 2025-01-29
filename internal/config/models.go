package config

type Task struct {
	ID             int64  `json:"id"`
	TaskTitle      string `json:"task_title"`
	Progress       int64  `json:"progress"`
	ProjectID      int64  `json:"project_id"`
	AssignedUserID int64  `json:"assigned_user_id"`
	CreatedAt      string `json:"created_at"`
}

type User struct {
	ID         int64  `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserRoleID int64  `json:"user_role_id"`
	SbuID      int64  `json:"sbu_id"`
	CreatedAt  string `json:"created_at"`
}

type Project struct {
	ID          int64  `json:"id"`
	ProjectName string `json:"project_name"`
	ClientID    int64  `json:"client_id"`
	CreatedAt   string `json:"created_at"`
}

type Client struct {
	ID         int64  `json:"id"`
	ClientName string `json:"client_name"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}

type SBU struct {
	ID            int64  `json:"id"`
	SbuName       string `json:"sbu_name"`
	SbuHeadUserID int64  `json:"sbu_head_user_id"`
	CreatedAt     string `json:"created_at"`
}
