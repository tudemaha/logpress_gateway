package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tudemaha/logpress_gateway/internal/compress/service"
	"github.com/tudemaha/logpress_gateway/internal/global/dto"
	globalService "github.com/tudemaha/logpress_gateway/internal/global/service"
	"github.com/tudemaha/logpress_gateway/pkg/logpress"
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
			var sr dto.ServerResponse
			var err error

			if !service.CheckConnection() {
				return
			}

			dbSize, _ := globalService.GetDBSize()

			filename, err := service.CreateDump()
			if err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}
			if err = service.CompressGZIP(filename); err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}

			if err = service.DeleteUncompressed(filename); err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}

			if sr, err = service.TransferCompressedDump(filename); err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}

			transferLog := fmt.Sprintf("%s,%f %s,%d ns,%d ns,%d ns,%d ns\n",
				sr.Data.TimestampSummary.StartTime[:19],
				dbSize, logpress.LoadLogpressConfig.ThresholdUnit,
				sr.Data.DurationSummary.TransferDuration,
				sr.Data.DurationSummary.DecompressDuration,
				sr.Data.DurationSummary.MergeDuration,
				sr.Data.DurationSummary.TotalDuration,
			)
			if err := globalService.AppendTransferLog(transferLog); err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}

			if err := service.DeleteOldData(); err != nil {
				log.Fatalf("ERROR CompressHandler fatal error: %v", err)
			}
		}()
	}
}
