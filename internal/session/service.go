package session

import (
	"context"
	"time"
)

type SessionService struct {
	repo *SessionRepository
}

func NewSessionService(repo *SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) CreateSession(ctx context.Context, personID int, ipAddress, userAgent string) (string, error) {
	return s.repo.CreateSession(ctx, personID, "", "", ipAddress, ipAddress, userAgent)
}

func (s *SessionService) GetSessionByID(ctx context.Context, sessionID string) (*Session, error) {
	return s.repo.GetSessionByID(ctx, sessionID)
}

func (s *SessionService) InvalidateSession(ctx context.Context, sessionID string) error {
	return s.repo.InvalidateSession(ctx, sessionID)
}

func (s *SessionService) UpdateLastActivity(ctx context.Context, sessionID string) error {
	return s.repo.UpdateLastActivity(ctx, sessionID, time.Now())
}
