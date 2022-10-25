package templates

import (
	"html/template"
	"time"
)

const DateTimeFormat = "2006/01/02 15:04:05"

func dateTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

func noEscape(str string) template.HTML {
	return template.HTML(str)
}
