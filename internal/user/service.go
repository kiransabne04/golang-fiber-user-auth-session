package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *UserRepository
}

type DashboardData struct {
	OrgName  string `json:"org_name"`
	Greeting string `json:"greeting"`
}

type GetMeResponse struct {
	User     UserResponse `json:"user"`
	OrgName  string       `json:"org_name"`
	Greeting string       `json:"greeting"`
}

// NewUserService creates a new instance of userService.
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

// RegisterUser registears a new user with hashed password.
func (s *UserService) RegisterUser(ctx context.Context, name, email, password string) (int, error) {
	// Validate input
	if name == "" || email == "" || password == "" {
		return 0, errors.New("name, email, and password are required")
	}

	// Hash the password
	_, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// Create user struct
	user := &User{
		//Name:         name,
		Email: email,
		//PasswordHash: string(hashedPassword),
	}

	// Save user to the database
	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) ListAllUser() (User, error) {
	return User{ID: 1, FirstName: "Kiran", LastName: "Sabne"}, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// verify the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
