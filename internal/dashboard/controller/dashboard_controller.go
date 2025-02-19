package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	dashboardDto "github.com/tudemaha/logpress_gateway/internal/dashboard/dto"
	globalDto "github.com/tudemaha/logpress_gateway/internal/global/dto"
	"github.com/tudemaha/logpress_gateway/internal/global/service"
	"github.com/tudemaha/logpress_gateway/internal/global/utils"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
)

func DashboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := utils.GetSession(w, r)
		// if !auth {
		// 	http.Redirect(w, r, "/login", http.StatusSeeOther)
		// 	return
		// }

		if r.Method == "GET" {
			renderDashboard(w, username)
			return
		}
	}
}

func renderDashboard(w http.ResponseWriter, username string) {
	var response globalDto.Response

	templ, err := template.ParseFiles("./internal/dashboard/template/index.gohtml")
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		response.DefaultInternalError()
		response.Data = []string{err.Error()}
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	dbSize, _ := service.GetDBSize()

	data := dashboardDto.DashboardData{
		Username:       username,
		LogpressConfig: logpress.LoadLogpressConfig,
		DBSize:         dbSize,
	}

	err = templ.Execute(w, data)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		response.Data = []string{err.Error()}
		response.DefaultInternalError()
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}
}
