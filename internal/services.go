package internal

import (
	"fiber-user-auth-session/config"
	"fiber-user-auth-session/internal/user"
)

// AppServices struct will contain all repositories as fields
type AppServices struct {
	UserService  *user.UserService
	SessionStore session.SessionStore
	// // Add other repositories as needed, e.g.:
	// // CompanyRepository *company.CompanyRepository
	// TenantService *tenant.TenantService
	// OrgService    *company.OrgService
}

// NewAppServices initializes all repositories and returns AppServices
func NewAppServices(appConfig *config.AppConfig) *AppServices {

	// Initialize user repository and service
	userRepo := user.NewUserRepository(appConfig.Database)
	userService := user.NewUserService(userRepo)

	return &AppServices{
		UserService:  userService,
		// SessionStore: appConfig.SessionStore, // session store
		// // Initialize other repositories similarly
		// TenantService: tenant.NewTenantService(appConfig),
		// OrgService:    company.NewOrganizationService(appConfig),
	}
}