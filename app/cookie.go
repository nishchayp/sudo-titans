package app

import (
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

/*
** contains all middleware fuctions to interact with cookies
 */

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func SetCookieHandler(stUserCookie *STUserCookie, w http.ResponseWriter) {

	if encoded, err := cookieHandler.Encode("stUserCookie", stUserCookie); err == nil {
		cookie := &http.Cookie{
			Name:    "stUserCookie",
			Value:   encoded,
			Path:    "/",
			Expires: time.Now().Add(12 * time.Hour),
		}
		http.SetCookie(w, cookie)
	}
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) (stUserCookie STUserCookie) {
	if cookie, err := r.Cookie("stUserCookie"); err == nil {
		if err = cookieHandler.Decode("stUserCookie", cookie.Value, &stUserCookie); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	return stUserCookie
}

func ClearCookieHandler(w http.ResponseWriter) {

	cookie := &http.Cookie{
		Name:   "stUserCookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
