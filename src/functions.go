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

// -----------------------------------------------------

func Log(message string, err ...error) {
	now := time.Now()
	dateTimeStr := now.Format("[01/02/2006 - 15:04:05] ")
	fmt.Print(dateTimeStr)
	fmt.Println(message)

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
