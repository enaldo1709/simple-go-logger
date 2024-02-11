package golog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"
)

var logLevels = map[string][]any{
	"info":  {Green, "INFO "},
	"warn":  {YellowLight, "WARN "},
	"error": {Red, "ERROR"},
}

func Info(elem ...any) {
	printLog("info", Reset, elem)
}

func Warn(elem ...any) {
	printLog("warn", Reset, elem)
}

func Error(elem ...any) {
	printLog("error", Reset, elem)
}

func Err(err error) string {
	return fmt.Sprintf("\n%s-> %+v %s", Red, err, Reset)
}

func Any(v any) string {
	content := new(bytes.Buffer)
	if err := json.NewEncoder(content).Encode(v); err != nil {
		return fmt.Sprintf("\n%s%v%s", Cyan, v, Reset)
	}
	return fmt.Sprintf("\n%s%s%s", Cyan, content.String(), Reset)
}

func printLog(level, nextColor string, elem []any) {
	if !checked {
		check()
	}
	headers := logLevels[level]
	pc, _, no, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok {
		headers = append(headers, Purple, fmt.Sprintf("[%s:%d] -", formatLocation(details.Name()), no))
	}
	fmt.Print(GrayDark)
	headers = append(headers, nextColor)
	elem = append(headers, elem...)
	log.Println(append(elem, Reset)...)

}

func formatLocation(location string) string {
	splitted := strings.Split(location, "/")
	var toBeComplete []string
	toStrip := []string{}
	if len(splitted) > 2 {
		toStrip = splitted[:len(splitted)-2]
		toBeComplete = splitted[len(splitted)-2:]
	} else {
		toBeComplete = splitted
	}

	stripped := []string{}
	for _, st := range toStrip {
		stripped = append(stripped, string(st[0]))
	}

	return strings.Join(append(stripped, toBeComplete...), "/")
}
