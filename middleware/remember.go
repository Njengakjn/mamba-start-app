package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kjnjenga/webapp/models"
)

func (m *Middleware) CheckRemember(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !m.mamba.Session.Exists(r.Context(), "userID") {
			cookie, err := r.Cookie(fmt.Sprintf("_%s_remember", m.mamba.Title))
			if err != nil {
				// no cookie
				next.ServeHTTP(w, r)
			} else {
				key := cookie.Value

				var u models.User

				if len(key) > 0 {
					split := strings.Split(key, "|")
					uid, hash := split[0], split[1]
					id, _ := strconv.Atoi(uid)

					validHash := u.CheckForRememberToken(id, hash)
					if !validHash {
						deleteRememberCookie(m, w, r)
						m.mamba.Session.Put(r.Context(), "error", "You've been logged out from another device")
						next.ServeHTTP(w, r)
					} else {
						user, _ := u.Get(id)
						m.mamba.Session.Put(r.Context(), "userID", user.ID)
						m.mamba.Session.Put(r.Context(), "remember_token", hash)
						next.ServeHTTP(w, r)
					}
				} else {
					// key length is zero, so it's probably a leftover cookie (user has not closed browser).
					deleteRememberCookie(m, w, r)
					next.ServeHTTP(w, r)
				}
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func deleteRememberCookie(m *Middleware, w http.ResponseWriter, r *http.Request) {
	_ = m.mamba.Session.RenewToken(r.Context())

	newCookie := http.Cookie{
		Name:     fmt.Sprintf("_%s_remember", m.mamba.Title),
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-100 * time.Hour),
		HttpOnly: true,
		Domain:   m.mamba.Session.Cookie.Domain,
		MaxAge:   -1,
		Secure:   m.mamba.Session.Cookie.Secure,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &newCookie)

	m.mamba.Session.Remove(r.Context(), "userID")
	m.mamba.Session.Destroy(r.Context())

	_ = m.mamba.Session.RenewToken(r.Context())
}
