package dto

import "github.com/tudemaha/logpress_gateway/pkg/logpress"

type DashboardData struct {
	Username       string
	DBSize         float64
	LogpressConfig logpress.LogpressConfig
	TransferLogs   [][]string
}
