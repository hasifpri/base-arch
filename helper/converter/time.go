package helperconverter

import "time"

func ConvertTimeToString(t *time.Time) string {
	if t == nil {
		return ""
	}
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}
