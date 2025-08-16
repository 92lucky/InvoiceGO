package handlers

import (
	"html/template"
	"invoice-go/auth"
	"invoice-go/config"
	"invoice-go/model"
	"invoice-go/service"
	"net/http"
)

// Struct khusus untuk template (supaya tidak lempar model mentah)
type setupView struct {
	Email           string
	NamaPT          string
	NamaBank        string
	NoRekening      string
	PenanggungJawab string
	Alamat          string
	Kabupaten       string
}

func HandleSetup(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.GetSession(r)
		email := session.Values["email"].(string)

		switch r.Method {
		case http.MethodGet:
			profile, err := service.LoadProfileByEmail(config.DB, email)
			if err != nil {
				// Jika tidak ditemukan, isi struct kosong (Email tetap diisi agar tidak null di form)
				profile = &model.AppProfile{Email: email}
			}

			// mapping profile → setupView (hanya field yang perlu ditampilkan)
			viewData := setupView{
				Email:           profile.Email,
				NamaPT:          profile.NamaPT,
				NamaBank:        profile.NamaBank,
				NoRekening:      profile.NoRekening,
				PenanggungJawab: profile.PenanggungJawab,
				Alamat:          profile.Alamat,
				Kabupaten:       profile.Kabupaten,
			}

			if err := tmpl.ExecuteTemplate(w, "setup.html", viewData); err != nil {
				http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
			}

		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Invalid form data", http.StatusBadRequest)
				return
			}

			profile := model.AppProfile{
				Email:           email,
				NamaPT:          r.FormValue("nama_pt"),
				NamaBank:        r.FormValue("nama_bank"),
				NoRekening:      r.FormValue("no_rekening"),
				PenanggungJawab: r.FormValue("penanggung_jawab"),
				Alamat:          r.FormValue("alamat"),
				Kabupaten:       r.FormValue("kabupaten"),
			}

			if err := service.UpdateProfile(config.DB, profile); err != nil {
				http.Error(w, "Gagal simpan profil: "+err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/index", http.StatusSeeOther)
		}
	}
}
