package types

import "time"

type CallRecord struct {
	StartTime time.Time
	EndTime   time.Time
	Sender    string
	Duration  time.Duration
}

type FBParticipant struct {
	Name string `json:"name"`
}

type FBMessageCall struct {
	Type         string `json:"type"`
	SenderName   string `json:"sender_name"`
	CallDuration string `json:"call_duration"`
	TimestampMS  int    `json:"timestamp_ms"`
	Content      string `json:"content"`
}

type FBData struct {
	Participants []FBParticipant `json:"participants"`
	MessageCalls []FBMessageCall `json:"messages"`
}
