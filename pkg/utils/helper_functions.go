package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GooglePayload struct {
	SUB           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
	Error	      string `json:"error"`
}

func ConvertToken(accessToken string) (*GooglePayload, error) {
    resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
    if httpErr != nil {
        return nil, httpErr
    }

    defer resp.Body.Close()
    
    var data GooglePayload
    dataerr := json.NewDecoder(resp.Body).Decode(&data)
	if dataerr != nil {
		return nil, dataerr
	}

	if data.Error != "" {
		return nil, errors.New("invalid token")
	}

	return &data,nil
}