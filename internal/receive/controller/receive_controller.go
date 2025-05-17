package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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

	log.Println(string(body))

	sensorData, err := parseSensorData(string(body))
	if err != nil {
		response.DefaultInternalError()
		w.WriteHeader(response.Code)
		response.Data = []string{err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	log.Println(sensorData)

	if service.CheckNullValue(&sensorData) {
		response.DefaultBadRequest()
		w.WriteHeader(response.Code)
		response.Data = []string{"co, humid, lpg, smoke, or temp must not 0"}
		json.NewEncoder(w).Encode(response)
		return
	}

	service.TransformSensorData(&sensorData)
	log.Println(sensorData)

	db := database.DatabaseConnection()
	defer db.Close()

	id := uuid.New().String()
	gatewayID := os.Getenv("GATEWAY_ID")

	stmt := `INSERT INTO ` + os.Getenv("TABLE_NAME") +
		` (id, timestamp, node_id, gateway_id, temp, humid, soil_ph, soil_moisture, gas, gps)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(stmt,
		id,
		sensorData.Timestamp,
		sensorData.NodeID,
		gatewayID,
		sensorData.Temp,
		sensorData.Humid,
		sensorData.SoilPH,
		sensorData.SoilMoisture,
		sensorData.Gas,
		sensorData.GPS)
	if err != nil {
		response.DefaultInternalError()
		w.WriteHeader(response.Code)
		response.Data = []string{err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	stmtPersist := `INSERT INTO persist_sensors
		(id, timestamp, node_id, gateway_id, temp, humid, soil_ph, soil_moisture, gas, gps)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(stmtPersist,
		id,
		sensorData.Timestamp,
		sensorData.NodeID,
		gatewayID,
		sensorData.Temp,
		sensorData.Humid,
		sensorData.SoilPH,
		sensorData.SoilMoisture,
		sensorData.Gas,
		sensorData.GPS)
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

func parseSensorData(body string) (receiveDto.NewSensorData, error) {
	var result receiveDto.NewSensorData
	var err error

	bodyStr := strings.Split(body, ";")

	result.Timestamp, err = time.Parse("2006-01-02T15:04:05Z", bodyStr[0])
	if err != nil {
		return result, err
	}
	result.NodeID = bodyStr[1]
	result.Temp, _ = strconv.ParseFloat(bodyStr[2], 64)
	result.Humid, _ = strconv.ParseFloat(bodyStr[3], 64)
	result.SoilPH, _ = strconv.ParseFloat(bodyStr[4], 64)
	result.SoilMoisture, _ = strconv.ParseFloat(bodyStr[5], 64)
	result.Gas, _ = strconv.ParseFloat(bodyStr[6], 64)
	result.GPS = bodyStr[7]

	return result, nil
}
