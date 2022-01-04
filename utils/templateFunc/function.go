package templateFunc

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"html/template"
	"time"
)

func GetFunctions() template.FuncMap {
	return template.FuncMap{
		"html":             noescape,
		"inc":              inc,
		"shorten_number":   intToString,
		"number_separator": numberComas,
		"time":             timeFormat,
	}
}

func noescape(str string) template.HTML {
	return template.HTML(str)
}

func timeFormat(t time.Time) string {
	return t.Format("02/01/2006 15:01")
}

func inc(i int) int {
	return i + 1
}

func intToString(i uint64) string {
	if i < 1000 {
		return fmt.Sprintf("%d", i)
	} else if i < 1000000 {
		return fmt.Sprintf("%.1fk", float64(i)/1000.0)
	}
	return fmt.Sprintf("%.1fm", float64(i)/1000000.0)
}

func numberComas(i uint64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d\n", i)
}
