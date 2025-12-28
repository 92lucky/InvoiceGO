package auth

import (
	"encoding/json"

	"invoice-go/config"
	"invoice-go/internal/repository"

	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var Store *sessions.CookieStore

func InitSession() {
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		panic("SESSION_KEY belum di-set atau gagal terbaca dari .env")
	}

	Store = sessions.NewCookieStore([]byte(sessionKey))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   false,
	}
}

var googleOauthConfig *oauth2.Config

func InitOAuthConfig() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func RegisterAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", HandleLogin)
	mux.HandleFunc("/callback", HandleCallback)
	mux.HandleFunc("/logout", HandleLogout)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	returnTo := r.URL.Query().Get("returnTo")
	if returnTo == "" {
		returnTo = "/index"
	}

	session, _ := Store.Get(r, "session")
	session.Values["returnTo"] = returnTo
	session.Save(r, w)

	url := googleOauthConfig.AuthCodeURL("state-random")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Kode tidak ditemukan di callback", http.StatusBadRequest)
		return
	}

	// Tukar kode dengan token
	token, err := googleOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Gagal menukar kode dengan token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := googleOauthConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Gagal mendapatkan user info dari Google: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Gagal decode user info", http.StatusInternalServerError)
		return
	}

	email, ok := userInfo["email"].(string)
	if !ok || email == "" {
		http.Error(w, "Email tidak tersedia", http.StatusInternalServerError)
		return
	}

	// ✅ Set session authenticated
	session, _ := Store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["email"] = email
	session.Save(r, w)

	// ✅ Cek profil user
	profile, err := repository.GetUserEmail(config.DB, email)
	if err != nil || profile.NamaPT == "" {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}

	// Redirect ke index
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}


func HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
