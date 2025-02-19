package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tudemaha/logpress_gateway/internal/compress/service"
	"github.com/tudemaha/logpress_gateway/internal/global/dto"
)

func CompressHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		var response dto.Response

		if r.Method != "POST" {
			response.DefaultNotAllowed()
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}

		response.DefaultOK()
		json.NewEncoder(w).Encode(response)

		go func() {
			filename := service.CreateDump()
			if err := service.CompressGZIP(filename); err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}

			if err := service.DeleteUncompressed(filename); err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}

			if err := service.TransferCompressedDump(filename); err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}
		}()
	}
}
