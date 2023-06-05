package main

import (
	"bot-for-academic-attendance/app"
	"bot-for-academic-attendance/auth"
	"bot-for-academic-attendance/utils"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func main() {
	username := "NIM"
	password := "PASSWORD"

	client := &http.Client{}
	jar, err := cookiejar.New(nil)
	utils.CheckError(err)

	client.Jar = jar

	fmt.Printf("Logging in as %s...\n", username)
	auth.Login(client, username, password)

	for {
		utils.ClearConsole()
		fmt.Printf("---------------------------------- Automatic Attendance ----------------------------------\n")
		app.GetAttendance(client)
		time.Sleep(10 * time.Minute)
	}
}
