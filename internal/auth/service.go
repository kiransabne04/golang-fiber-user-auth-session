package auth

import (
	"context"
	"fiber-user-auth-session/internal/session"
	"fiber-user-auth-session/internal/user"
	"log"
)

type AuthService struct {
	UserRepo    *user.UserRepository
	SessionRepo *session.SessionRepository
}

func NewAuthService(userRepo *user.UserRepository, sessionRepo *session.SessionRepository) *AuthService {
	return &AuthService{UserRepo: userRepo, SessionRepo: sessionRepo}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, string, error) {
	//verify user creds
	user, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}
	log.Println("user -> ", user)
	//generate tokens
	// Generate tokens
	// accessToken, err := pkg.GenerateAccessToken(user.ID, user.ID)
	// if err != nil {
	// 	return "", "", err
	// }
	// refreshToken, err := pkg.GenerateRefreshToken()
	// if err != nil {
	// 	return "", "", err
	// }

	// // Create session
	// _, err = s.SessionRepo.CreateSession(ctx, user.ID, accessToken, refreshToken)
	// if err != nil {
	// 	return "", "", err
	// }

	return "", "", nil
}