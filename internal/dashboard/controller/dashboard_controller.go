package controller

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	dashboardDto "github.com/tudemaha/logpress_gateway/internal/dashboard/dto"
	globalDto "github.com/tudemaha/logpress_gateway/internal/global/dto"
	"github.com/tudemaha/logpress_gateway/internal/global/service"
	"github.com/tudemaha/logpress_gateway/internal/global/utils"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
)

func DashboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, auth := utils.GetSession(w, r)
		if !auth {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if r.Method == "GET" {
			renderDashboard(w, username)
			return
		}

		var response globalDto.Response
		w.Header().Add("Content-Type", "application/json")
		response.DefaultNotAllowed()
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
	}
}

func renderDashboard(w http.ResponseWriter, username string) {
	var response globalDto.Response

	templ, err := template.ParseFiles("./public/templates/index.gohtml")
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		response.DefaultInternalError()
		response.Data = []string{err.Error()}
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	dbSize, _ := service.GetDBSize()
	transferLogs, _ := service.ReadTransferLog()

	data := dashboardDto.DashboardData{
		Username:       username,
		LogpressConfig: logpress.LoadLogpressConfig,
		DBSize:         dbSize,
		TransferLogs:   transferLogs,
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

func UpdateConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var response globalDto.Response
		w.Header().Add("Content-Type", "application/json")
		oldConfig := logpress.LoadLogpressConfig

		if r.Method != "PUT" {
			response.DefaultNotAllowed()
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		_, auth := utils.GetSession(w, r)
		if !auth {
			response.DefaultForbidden()
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.DefaultInternalError()
			w.WriteHeader(response.Code)
			response.Data = []string{err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		var newConfig logpress.LogpressConfig
		err = json.Unmarshal(body, &newConfig)
		if err != nil {
			response.DefaultInternalError()
			w.WriteHeader(response.Code)
			response.Data = []string{err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		oldConfig.CronInterval = newConfig.CronInterval
		oldConfig.CronUnit = newConfig.CronUnit
		oldConfig.Threshold = newConfig.Threshold
		oldConfig.ThresholdUnit = newConfig.ThresholdUnit

		err = logpress.WriteConfig(oldConfig)
		if err != nil {
			response.DefaultInternalError()
			w.WriteHeader(response.Code)
			response.Data = []string{err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		logpress.ReadConfig()

		response.DefaultOK()
		json.NewEncoder(w).Encode(response)
	}
}
