package session

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/hunterjsb/sqlbase/internal/redis"
)

const (
	sessionCookieName = "session_token"
	sessionExpiration = 24 * time.Hour
)

func generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func CreateSession(w http.ResponseWriter, userID string) {
	sessionID := generateSessionID()

	// Set session in Redis
	err := redis.Client.Set(redis.Ctx, sessionID, userID, sessionExpiration).Err()
	if err != nil {
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:    sessionCookieName,
		Value:   sessionID,
		Expires: time.Now().Add(sessionExpiration),
	})
}

func GetUserID(r *http.Request) (string, error) {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return "", err
	}

	userID, err := redis.Client.Get(redis.Ctx, cookie.Value).Result()
	if err != nil {
		return "", err
	}

	return userID, nil
}

func DestroySession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(sessionCookieName)
	if err != nil {
		return
	}

	// Delete session from Redis
	redis.Client.Del(redis.Ctx, cookie.Value)

	// Delete session cookie
	http.SetCookie(w, &http.Cookie{
		Name:    sessionCookieName,
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}
