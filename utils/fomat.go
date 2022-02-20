package utils

import "time"

func FormatDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
