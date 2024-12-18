package auth

import (
	"context"
	"fiber-user-auth-session/internal/session"
	"fiber-user-auth-session/internal/user"
	"fiber-user-auth-session/pkg"
	"log"
)

type AuthService struct {
	UserRepo    *user.UserRepository
	SessionRepo *session.SessionRepository
	SecretKey   []byte
}

func NewAuthService(userRepo *user.UserRepository, sessionRepo *session.SessionRepository, secretKey []byte) *AuthService {
	return &AuthService{UserRepo: userRepo, SessionRepo: sessionRepo, SecretKey: secretKey}
}

func (s *AuthService) LoginService(ctx context.Context, email, password string) (string, string, error) {
	//verify user creds
	user, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}
	log.Println("user -> ", user)

	// vreate session
	// _, err = s.SessionRepo.CreateSession(ctx, user.ID, accessToken, refreshToken, "", "", "")
	// if err != nil {
	// 	return "", "", err
	// }
	sessionID, err := s.SessionRepo.CreateNewSession(ctx, user.ID, "", "", "")
	if err != nil {
		return "", "", err
	}
	log.Println("sessionID created -> ", sessionID)

	//generate access token
	accessToken, err := pkg.GenerateAccessToken(sessionID, user.ID, s.SecretKey)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := pkg.GenerateRefreshToken(s.SecretKey)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
