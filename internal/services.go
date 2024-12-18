package internal

import (
	"fiber-user-auth-session/internal/auth"
	"fiber-user-auth-session/internal/session"
	"fiber-user-auth-session/internal/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

// AppServices struct will contain all repositories as fields
type AppServices struct {
	UserService    *user.UserService
	SessionService *session.SessionService
	AuthService    *auth.AuthService
	// // Add other repositories as needed, e.g.:
	// // CompanyRepository *company.CompanyRepository
	// TenantService *tenant.TenantService
	// OrgService    *company.OrgService
}

// NewAppServices initializes all repositories and returns AppServices
func NewAppServices(db *pgxpool.Pool, secretKey string) *AppServices {

	// Initialize user repository and service
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	// Initailize session repo & Service
	sessionRepo := session.NewSessionRepository(db)
	sessionService := session.NewSessionService(sessionRepo)

	authService := auth.NewAuthService(userRepo, sessionRepo, []byte(secretKey))

	return &AppServices{
		UserService:    userService,
		SessionService: sessionService,
		AuthService:    authService,
		// SessionStore: appConfig.SessionStore, // session store
		// // Initialize other repositories similarly
		// TenantService: tenant.NewTenantService(appConfig),
		// OrgService:    company.NewOrganizationService(appConfig),
	}
}
