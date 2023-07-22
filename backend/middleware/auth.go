package middleware

import (
	"database/sql"
	"net/http"

	"code-database/config"

	"github.com/gorilla/sessions"
)

// AdminAuth アドミン画面への認証用ミドルウェア
func AdminAuth(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store := sessions.NewCookieStore([]byte(config.SessionKey))
		session, _ := store.Get(r, "cookie-name")
		if isAuth, ok := session.Values["authenticated"].(bool); ok && isAuth {
			db, err := sql.Open("mysql", config.SQLEnv)
			if err != nil {
				panic(err.Error())
			}
			defer db.Close()
			fn(w, r)
			return
		}
		http.Redirect(w, r, config.AdminLoginPath, http.StatusFound)
		return
	}
}

// UserAuth ユーザー画面への認証用ミドルウェア
func UserAuth(fn func(w http.ResponseWriter, r *http.Request, auth bool)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store := sessions.NewCookieStore([]byte(config.SessionKey))
		session, _ := store.Get(r, "cookie-name")
		auth := false
		if isAuth, ok := session.Values["authenticated"].(bool); ok && isAuth {
			auth = true
		}
		db, err := sql.Open("mysql", config.SQLEnv)
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		fn(w, r, auth)
	}
}
