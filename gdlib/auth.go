package gdlib

import (
	"encoding/json"
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/op/go-logging"
	uuid "github.com/satori/go.uuid"
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

// DoUser send user request and parse response
func DoUser(mobile string) (err error) {
	const label = "Verify Code"
	const version = "1.0.0"
	const company = "GAP Cli"

	prompt3 := promptui.Prompt{Label: label}
	verifyCode, err := prompt3.Run()

	if err != nil {
		return
	}

	uuid := uuid.Must(uuid.NewV4())
	resultUser := User(uuid.String(), version, company, verifyCode, mobile)

	var data map[string]interface{}

	err = json.Unmarshal([]byte(resultUser), &data)

	if err != nil {
		return err
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

	log.Debug(resultUser)
	return
}
