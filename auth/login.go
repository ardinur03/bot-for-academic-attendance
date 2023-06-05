package auth

import (
	"bot-for-academic-attendance/utils"
	"net/http"
	"net/url"
	"strings"
)

func Login(client *http.Client, username, password string) {
	loginURL := "https://akademik.polban.ac.id/laman/login"
	loginData := url.Values{
		"username": {username},
		"password": {password},
		"submit":   {""},
	}
	req, err := createLoginRequest(loginURL, loginData)
	utils.CheckError(err)

	resp, err := client.Do(req)
	utils.HandleResponse(resp, err)

	print("Login success!")
}

func createLoginRequest(loginURL string, loginData url.Values) (*http.Request, error) {
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(loginData.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}
