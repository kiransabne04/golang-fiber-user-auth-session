package session

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Session struct {
	SessionID    string
	PersonID     int
	IPAddress    string
	UserAgent    string
	LastActivity time.Time
	IsActive     bool
}

type SessionRepository struct {
	db *pgxpool.Pool
}

func NewSessionRepository(db *pgxpool.Pool) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) GetSessionByID(ctx context.Context, sessionID string) (*Session, error) {
	query := `
        SELECT session_id, person_id, ip_address, last_activity, is_active
        FROM person_session
        WHERE session_id = $1
    `
	var session Session
	err := r.db.QueryRow(ctx, query, sessionID).Scan(&session.SessionID, &session.PersonID, &session.IPAddress, &session.LastActivity, &session.IsActive)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepository) InvalidateSession(ctx context.Context, sessionID string) error {
	query := `UPDATE person_session SET is_active = FALSE WHERE session_id = $1`
	_, err := r.db.Exec(ctx, query, sessionID)
	return err
}

func (r *SessionRepository) UpdateLastActivity(ctx context.Context, sessionID string, lastActivity time.Time) error {
	query := `UPDATE person_session SET last_activity = $1 WHERE session_id = $2`
	_, err := r.db.Exec(ctx, query, lastActivity, sessionID)
	return err
}

func (r *SessionRepository) CreateSession(ctx context.Context, personID int, accessTokenID, refreshTokenID int64, deviceInfo, ipAddress, userAgent string) (string, error) {
	var sessionID string
	query := `
        INSERT INTO person_session (person_id, access_token_id, refresh_token_id, start_time, device_info, ip_address, user_agent)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING session_id
    `
	err := r.db.QueryRow(ctx, query, personID, accessTokenID, refreshTokenID, time.Now(), deviceInfo, ipAddress, userAgent).Scan(&sessionID)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}