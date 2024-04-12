package Projet_GO_Reservation

import (
	"fmt"
	"strings"
	"time"
)

func log(message string, err ...error) {
	now := time.Now()
	dateTimeStr := now.Format("[01/02/2006 - 15:04:05] ")
	fmt.Print(dateTimeStr)
	fmt.Println(message)

	if err != nil {
		fmt.Println(err)
	}
}

func log(message string) string {
	now := time.Now()
	dateTimeStr := now.Format("[01/02/2006 - 15:04:05] " + message)
	return dateTimeStr
}

func (l *LogHelper) Error(message string, err ...error) {
	var result = log(message)

	fmt.Print("ERROR : " + result)

	if err != nil {
		fmt.Println(err)
	}
}

func (l *LogHelper) Infos(message string, err ...error) {
	var result = log(message)

	fmt.Print("INFOS : " + result)

	if err != nil {
		fmt.Println(err)
	}
}

func (l *LogHelper) Debug(message string, err ...error) {
	var result = log(message)

	fmt.Print("DEBUG : " + result)

	if err != nil {
		fmt.Println(err)
	}
}

// -----------------------------------------------------

func Println(message string) {
	fmt.Println(message)
}

// -----------------------------------------------------

func arrayToString(arr []string, noQuotes ...bool) string {
	if len(arr) == 0 {
		return ""
	}

	var sb strings.Builder
	for i, s := range arr {
		//sb.WriteString(s)

		_, err := strconv.Atoi(s)
		if err != nil && noQuotes == nil {
			// Cast to int failed
			sb.WriteString(`'` + s + `'`)
		} else {
			// Cast to int ok
			sb.WriteString(s)
		}

		if i < len(arr)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

// -----------------------------------------------------

func concatColumnWithValues(columns []string, values []string) string {

	if len(columns) == 0 || len(values) == 0 {
		Log.Error("Plz columns and values string array must have at least one key each")
		return ""
	}

	if len(columns) != len(values) {
		Log.Error("Plz columns and values string array must have the same length")
		return ""
	}

	var sb strings.Builder
	for i, s := range values {
		//sb.WriteString(s)

		_, err := strconv.Atoi(s)
		if err != nil {
			// Cast to int failed
			sb.WriteString(columns[i] + `='` + s + `'`)
		} else {
			// Cast to int ok
			sb.WriteString(columns[i] + "=" + s)
		}

		if i < len(columns)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}
