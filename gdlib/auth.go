package gdlib

import (
	"encoding/json"
	"errors"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("default")

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

// DoUserAdd send user add request and handle response
func DoUserAdd(mobile string) (err error) {
	resultUserAdd := UserAdd(mobile)

	var data map[string]interface{}

	err = json.Unmarshal([]byte(resultUserAdd), &data)

	if err != nil {
		return
	}

	if data["status"] == "success" {
		return
	}

	msg, ok := data["message"].(string)

	if ok {
		err = errors.New(msg)
	} else {
		err = errors.New("Failed")
	}

	log.Debug(resultUserAdd)
	return
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
