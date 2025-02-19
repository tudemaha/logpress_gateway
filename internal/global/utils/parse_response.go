package utils

import (
	"encoding/json"

	"github.com/tudemaha/logpress_gateway/internal/global/dto"
)

func ParseServerResponse(response []byte) (dto.ServerResponse, error) {
	var sr dto.ServerResponse

	err := json.Unmarshal(response, &sr)
	if err != nil {
		return sr, err
	}

	return sr, nil
}
