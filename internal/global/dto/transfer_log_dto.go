package dto

import "time"

type Log struct {
	Timestamp    time.Time
	DatabaseSize float64
	Transfer     uint
	Decompress   uint
	Merge        uint
	Total        uint
}
