package Projet_GO_Reservation

import (
	"fmt"
	"time"
)

type LogHelper struct {
}

// -----------------------------------------------------

func ILog(message string, err ...error) {
	now := time.Now()
	dateTimeStr := now.Format("[01/02/2006 - 15:04:05] ")
	fmt.Print(dateTimeStr)
	fmt.Println(message)

	if err != nil {
		fmt.Println(err)
	}
}

func log() string {
	now := time.Now()
	dateTimeStr := now.Format("[01/02/2006 - 15:04:05] ")
	return dateTimeStr
}

func (l *LogHelper) Error(message string, err ...error) {
	var result = log()

	fmt.Printf("\033[1;31m%s ERROR : \033[0m%s\n", result, message)

	if err != nil {
		fmt.Println(err)
	}
}

func (l *LogHelper) Infos(message string, err ...error) {
	var result = log()

	fmt.Println(result + "INFOS : " + message)

	if err != nil {
		fmt.Println(err)
	}
}

func (l *LogHelper) Debug(message string, err ...error) {
	var result = log()

	fmt.Printf("\033[1;32m%s DEBUG : \033[0m%s\n", result, message)

	if err != nil {
		fmt.Println(err)
	}
}

// -----------------------------------------------------

func Println(message string) {
	fmt.Println(message)
}
