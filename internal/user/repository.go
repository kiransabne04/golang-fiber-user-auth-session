package user

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignUpInput struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

type PersonSession struct {
	SessionID   uint64 `json:"session_i"`
	PersonLid   int    `json:"person_id"`
	Active      bool   `json:"is_active"`
	DeviceInfo  string `json:"device_info"`
	IPAddress   string `json:"ip_address"`
	UserAgent   string `json:"user_agent"`
	TempUrlPath string `json:"temp_url_path"`
	Login       string `json:"login"`
}

type Tenant struct {
	ID          int    `json:"id"`
	SlugName    string `json:"slug_name"`
	Description string `json:"description"`
	Active      bool   `json:"bool"`
}

type UserResponse struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
}

// userRepository handles database operations for users.
type UserRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository creates a new instance of userRepository.
func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser implements Repository.
func (r *UserRepository) CreateUser(ctx context.Context, user *User) (int, error) {
	// query := `
	// 	INSERT INTO users (name, email, password_hash)
	// 	VALUES ($1, $2, $3)
	// 	RETURNING id, created_at
	// `
	// err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.PasswordHash).
	// 	Scan(&user.ID, &user.CreatedAt)
	// if err != nil {
	// 	return err
	// }
	return 1, nil
}

// FindByEmail implements query returning user data for email id.
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := `SELECT id, email, first_name, last_name, password, active, created_at FROM person WHERE email = $1`
	err := r.db.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.Active, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ValidateTenant(slugName string) (bool, error) {
	var slugNameValue string
	log.Println("validateTenant slugName -> ", slugName)
	sqlStmt := `select slug_name from tenant where slug_name = $1`
	err := r.db.QueryRow(context.Background(), sqlStmt, slugName).Scan(&slugNameValue)
	if err != nil {
		log.Println("validateTenant -> ", err)

		return false, err
	}
	return true, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, email, first_name, last_name, password, active, created_at, updated_at
        FROM person
        WHERE email = $1 AND active = true`
	var user User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password,
		&user.Active, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
