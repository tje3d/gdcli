package gdlib

import (
	"encoding/json"
)

// UserAdd first step of authentication, Register or login
func UserAdd(mobile string) string {
	data := map[string]string{"mobile": mobile}
	byteData, _ := json.Marshal(data)

	options := RequestOptions{
		method: "GET",
		data:   string(byteData),
	}

	return sendRequest("user/add", options)
}

// User ...
func User(
	uuid string,
	appVersion string,
	company string,
	verificationCode string,
	mobile string,
) string {
	data := map[string]string{
		"device_uniqid":     uuid,
		"app_version":       appVersion,
		"company":           company,
		"verification_code": verificationCode,
		"mobile":            mobile,
	}

	byteData, _ := json.Marshal(data)

	options := RequestOptions{
		method: "POST",
		data:   string(byteData),
	}

	return sendRequest("user", options)
}
