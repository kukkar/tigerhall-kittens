package utils

import "time"

func ReverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

type RequestContext struct {
	SessionID     string
	RequestID     string
	ClientAppID   string
	UserID        string
	TransactionID string
	Method        string
	URI           string
	APIStartTime  time.Time
	IP            string
}
