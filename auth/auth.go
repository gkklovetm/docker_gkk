package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const REGISTRY_SERVER = "https://registry.docker.io"

type AuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	rootPath string `json:-`
}

func NewAuthConfig(username, password, email, rootPath string) *AuthConfig {
	return &AuthConfig{
		Username: username,
		Password: password,
		Email:    email,
		rootPath: rootPath,
	}
}

func Login(authConfig *AuthConfig) (string, error) {
	storeConfig := false
	reqStatusCode := 0
	var status string
	var errMsg string
	var reqBody []byte
	jsonBody, err := json.Marshal(authConfig)
	if err != nil {
		errMsg = fmt.Sprintf("Config Error:%s", err)
		return "", errors.New(errMsg)
	}

	b := strings.NewReader(string(jsonBody))
	reql, err := http.Post(REGISTRY_SERVER+"/v1/users", "application/json; charset=utf-8", b)
	if err != nil {
		errMsg = fmt.Sprintf("Server Error: %s", err)
		return "", errors.New(errMsg)
	}

	reqStatusCode = reql.StatusCode
	defer reql.Body.Close()
	reqBody, err = ioutil.ReadAll(reql.Body)
	if err != nil {
		errMsg = fmt.Sprintf("Server Error: [%#v] %s", reqStatusCode, err)
		return "", errors.New(errMsg)
	}

	if reqStatusCode == 201 {
		status = "Accout Created\n"
		storeConfig = true
	} else if reqStatusCode == 400 {
		if string(reqBody) == "Username or email already exist" {
			client := &http.Client{}
			req, err := http.NewRequest("GET", REGISTRY_SERVER+"/v1/users", nil)
			req.SetBasicAuth(authConfig.Username, authConfig.Password)
			resp, err := client.Do(req)
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return "", err
			}
			if resp.StatusCode == 200 {
				status = "Lofin Succeeded\n"
				storeConfig = true
			} else {
				status = fmt.Sprintf("Login: %s", body)
				return "", errors.New(status)
			}
		} else {
			status := fmt.Sprintf("Registration: %s", string(reqBody))
			return "", errors.New(status)
		}
	} else {
		status = fmt.Sprintf("[%s] : %s ", reqStatusCode, string(reqBody))
		return "", errors.New(status)
	}
	if storeConfig {
		//authStr := EncodeAuth(authConfig)
		//save
		// TODO
		fmt.Println("nihao")
	}
	return status, nil
}
