package date

import "time"

// GetCurrent returns the current date
func GetCurrent() time.Time {
	return time.Now().UTC()
}

// GetCurrentAsString returns the current date in the format of time.RFC3339
func GetCurrentAsString() string {
	return GetCurrent().Format(time.RFC3339)
}
