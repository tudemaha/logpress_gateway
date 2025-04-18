package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	globalDto "github.com/tudemaha/logpress_gateway/internal/global/dto"
	receiveDto "github.com/tudemaha/logpress_gateway/internal/receive/dto"
	"github.com/tudemaha/logpress_gateway/internal/receive/service"
	"github.com/tudemaha/logpress_gateway/pkg/database"
)

func ReceiveHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		var response globalDto.Response

		if r.Method != "POST" {
			response.DefaultNotAllowed()
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		receiveData(w, r)
	}
}

func receiveData(w http.ResponseWriter, r *http.Request) {
	var response globalDto.Response

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.DefaultInternalError()
		w.WriteHeader(response.Code)
		response.Data = []string{err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var sensorData receiveDto.SensorData
	err = json.Unmarshal(body, &sensorData)
	if err != nil {
		response.DefaultInternalError()
		w.WriteHeader(response.Code)
		response.Data = []string{err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	log.Println(sensorData)

	transformSensorData(&sensorData)
	log.Println(sensorData)

	db := database.DatabaseConnection()
	defer db.Close()

	id := uuid.New().String()

	stmt := `INSERT INTO sensors 
		(id, timestamp, device_id, co, humid, temp, lpg, smoke, light, motion)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(stmt,
		id,
		sensorData.Timestamp,
		sensorData.DeviceID,
		sensorData.CO,
		sensorData.Humid,
		sensorData.Temp,
		sensorData.LPG,
		sensorData.Smoke,
		sensorData.Light,
		sensorData.Motion)
	if err != nil {
		response.DefaultInternalError()
		w.WriteHeader(response.Code)
		response.Data = []string{err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	response.DefaultOK()
	json.NewEncoder(w).Encode(response)
}

func transformSensorData(data *receiveDto.SensorData) {
	data.CO = service.TransformFloat(data.CO, 14)
	data.Humid = service.TransformFloat(data.Humid, 14)
	data.Temp = service.TransformFloat(data.Temp, 14)
	data.LPG = service.TransformFloat(data.LPG, 14)
	data.Smoke = service.TransformFloat(data.Smoke, 14)
}
