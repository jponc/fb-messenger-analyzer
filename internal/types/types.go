package types

import "time"

type FBParticipant struct {
	Name string `json:"name"`
}

type FBMessage struct {
	TimestampMS         int `json:"timestamp_ms"`
	CallDurationSeconds int `json:"call_duration,omitempty"`
}

type FBFileData struct {
	Participants []FBParticipant `json:"participants"`
	Messages     []FBMessage
}

type CallRecord struct {
	StartMS         int
	EndMS           int
	DurationSeconds int
}

type CSVRow struct {
	StartTime       time.Time `csv:"start_time"`
	EndTime         time.Time `csv:"end_time"`
	DurationSeconds int       `csv:"duration_seconds"`
	DurationMinutes int       `csv:"duration_minutes"`
}
