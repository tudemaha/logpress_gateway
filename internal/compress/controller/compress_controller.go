package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tudemaha/logpress_gateway/internal/compress/service"
	"github.com/tudemaha/logpress_gateway/internal/global/dto"
)

func CompressHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		var response dto.Response

		response.DefaultOK()
		json.NewEncoder(w).Encode(response)

		go func() {
			filename := service.CreateDump()
			service.CompressGZIP(filename)
			service.DeleteUncompressed(filename)
		}()
	}
}
