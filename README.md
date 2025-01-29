# Employee-Tracking-System

[![codecov](https://codecov.io/github/MahediSabuj/go-teams/graph/badge.svg?token=8KG0CEAU74)](https://codecov.io/github/MahediSabuj/go-teams)

The **Employee Tracking System** is a web application designed to manage employees, SBUs (Strategic Business Units),
tasks, clients, and projects within a company. This system supports user authentication and role-based access, allowing
users such as admins and normal employees to interact with the system effectively.

## Features

### 1. Home Page

- A simple landing page with a navigation bar.
- Navigation links to other pages such as **Home**, **User**, **SBU**, **Task**, **Clients**, **Projects**, and *
  *Profile/Login**.

### 2. User Management

- Displays a list of all users in the system.
- Admin features:
    - Update user information.
    - Delete user accounts.

### 3. SBU Management

- View the list of SBUs in the company.
- Assign a user as the head of an SBU.
- Admin features:
    - Create, update, and delete SBUs.

### 4. Task Management

- Users:
    - View tasks assigned to them.
- Admin features:
    - Assign tasks to other users.
    - Update or delete tasks.

### 5. Client Management

- View a list of clients.
- Admin features:
    - Add new clients.
    - Update client information.
    - Delete clients.

### 6. Project Management

- Admin features:
    - Add new projects.
    - Assign clients and employees to projects.
    - Update or delete projects.

### 7. Profile/Login

- Users need to log in to access system features.
- If not logged in, the navigation bar displays the **Login** option.
- If logged in, the navigation bar displays the **Profile** page.

---

## Tech Stack

### Backend

- **Language**: Go (Golang)
- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: PostgreSQL

### Frontend

- **Languages**: HTML, CSS, JavaScript

### Key Go Packages

- `github.com/alexedwards/scs/v2`: Session management in Golang.
- `github.com/dgrijalva/jwt-go`: JSON Web Token (JWT) authentication.
- `github.com/gin-contrib/sessions`: Session middleware for Gin framework.
- `github.com/gin-gonic/gin`: Lightweight and fast web framework.
- `github.com/google/uuid`: UUID generation.
- `github.com/jackc/pgx/v5`: PostgreSQL driver and toolkit.
- `github.com/joho/godotenv`: Load environment variables from `.env` file.
- `github.com/stretchr/testify`: Testing utilities.
- `golang.org/x/crypto`: Cryptographic libraries.

---

## Routes

Here are the routes defined in the system:

### Static Files and Templates

- `r.LoadHTMLGlob("../../template/*")`: Load HTML templates.
- `r.Static("/static", "../../static")`: Serve static files like CSS and JavaScript.

### User Routes

- `GET /users`: Fetch user information.
- `POST /users`: Create a new user.
- `GET /profile`: User profile page.
- `GET /login`: Login page.
- `POST /login`: Authenticate user.
- `GET /logout`: Logout handler.
- `GET /registration`: User registration page.

### Admin User Management

- `GET /admin/users`: View list of users.
- `GET /admin/users/update/:id`: Update user form.
- `POST /admin/update-profile/:id`: Update user information.
- `DELETE /admin/users/delete/:id`: Delete a user.

### Task Management

- `GET /tasks`: View tasks for logged-in user.
- Admin Routes:
    - `GET /create/task`: Display form to create a task.
    - `POST /create/task`: Submit form to add a new task.
    - `GET /admin/task/update/:id`: Update task form.
    - `POST /admin/update-task/:id`: Submit updated task.
    - `DELETE /admin/task/delete/:id`: Delete a task.

### SBU Management

- `GET /admin/sbu`: View list of SBUs.
- Admin Routes:
    - `GET /create/sbu`: Create SBU form.
    - `POST /create/sbu`: Add SBU to the database.
    - `GET /sbu/update/:id`: Update SBU form.
    - `POST /sbu/update/:id`: Submit updated SBU.
    - `DELETE /admin/sbu/delete/:id`: Delete an SBU.

### Client Management

- `GET /clients`: View list of clients.
- Admin Routes:
    - `GET /create/client`: Create client form.
    - `POST /create/client`: Add client to the database.
    - `GET /admin/client/update/:id`: Update client form.
    - `POST /admin/client/update/:id`: Submit updated client information.
    - `DELETE /admin/client/delete/:id`: Delete a client.

### Project Management

- `GET /projects`: View list of projects.
- Admin Routes:
    - `GET /create/project`: Create project form.
    - `POST /create/project`: Add project to the database.
    - `GET /admin/project/update/:id`: Update project form.
    - `POST /admin/project/update/:id`: Submit updated project details.
    - `DELETE /admin/project/delete/:id`: Delete a project.

### Miscellaneous Routes

- `GET /`: Home page.
- `GET /about-us`: About Us page.
- `GET /contact-us`: Contact Us page.
- `GET /success`: Confirmation page.

---

## Prerequisites

1. **Go**: Install [Go 1.23+](https://golang.org/).
2. **PostgreSQL**: Setup a PostgreSQL database.
3. **Node.js and npm**: Install [Node.js](https://nodejs.org/) (required for frontend assets and dependency management).

---

## Getting Started

### 1. Clone the Repository

```bash
git clone <repository-url>
cd Employee-Tracking-System
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=your-database-host
DB_PORT=your-database-port
DB_USER=your-database-username
DB_PASSWORD=your-database-password
DB_NAME=your-database-name
JWT_SECRET=your-jwt-secret
```

### 3. Run the Application

```bash
go run main.go
```

### 4. Access the Application

Open your browser and navigate to `http://localhost:8080`.

---

## Directory Structure

```plaintext
Employee-Tracking-System/
├── cmd/                    # Application entry points
├── handlers/               # Route handlers
├── models/                 # Database models
├── routes/                 # Route definitions
├── static/                 # Static files (CSS, JS, images)
├── templates/              # HTML templates
├── .env                    # Environment variables
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies checksum
└── main.go                 # Main program
```

---

## License

This project is licensed under the MIT License.
