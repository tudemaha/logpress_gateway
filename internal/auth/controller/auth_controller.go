package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	authDto "github.com/tudemaha/logpress_gateway/internal/auth/dto"
	globalDto "github.com/tudemaha/logpress_gateway/internal/global/dto"
	"github.com/tudemaha/logpress_gateway/internal/global/utils"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, auth := utils.GetSession(w, r)
		if auth {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		if r.Method == "GET" {
			renderLoginPage(w)
			return
		}

		if r.Method == "POST" {
			validateLogin(w, r)
			return
		}

		var response globalDto.Response
		w.Header().Add("Content-Type", "application/json")
		response.DefaultNotAllowed()
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

func renderLoginPage(w http.ResponseWriter) {
	var response globalDto.Response

	templ, err := template.ParseFiles("./internal/auth/template/login.gohtml")
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		response.DefaultInternalError()
		response.Data = []string{err.Error()}
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = templ.Execute(w, nil)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		response.DefaultInternalError()
		response.Data = []string{err.Error()}
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}
}

func validateLogin(w http.ResponseWriter, r *http.Request) {
	var response globalDto.Response
	var loginInfo authDto.LoginInfo

	if err := r.ParseForm(); err != nil {
		w.Header().Add("Content-Type", "application/json")
		response.DefaultInternalError()
		response.Data = []string{err.Error()}
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	loginInfo.Username = r.PostForm.Get("username")
	loginInfo.Password = r.PostForm.Get("password")

	errUsername := logpress.LoadLogpressConfig.Username != loginInfo.Username

	errPass := bcrypt.CompareHashAndPassword([]byte(logpress.LoadLogpressConfig.Password),
		[]byte(loginInfo.Password))
	if errPass != nil || errUsername {
		data := authDto.LoginErrorDto{
			Error: true,
		}

		templ, err := template.ParseFiles("./public/templates/login.gohtml")
		if err != nil {
			w.Header().Add("Content-Type", "securecookieapplication/json")
			response.DefaultInternalError()
			response.Data = []string{err.Error()}
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		err = templ.Execute(w, data)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			response.DefaultInternalError()
			response.Data = []string{err.Error()}
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		return
	}

	utils.CreateSession(w, r, loginInfo.Username)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
