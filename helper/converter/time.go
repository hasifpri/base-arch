package helperconverter

import "time"

func ConvertTimeToString(t *time.Time) string {
	if t == nil {
		return ""
	}
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}

func ConvertStringToDate(tString string) (time.Time, error) {
	layout := "2006-01-02"

	t, err := time.Parse(layout, tString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func ConvertStringToTime(tString string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"

	t, err := time.Parse(layout, tString)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
