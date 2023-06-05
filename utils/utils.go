package utils

import (
	"fmt"
	"net/http"
	"os"
)

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func HandleResponse(resp *http.Response, err error) {
	if err != nil {
		CheckError(err)
	}

	// defer resp.Body.Close()
}
