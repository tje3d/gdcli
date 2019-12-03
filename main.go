package main

import (
	"encoding/json"
	"errors"
	"gdlib/gdlib"

	"github.com/manifoldco/promptui"
	"github.com/op/go-logging"
	uuid "github.com/satori/go.uuid"
)

var log = logging.MustGetLogger("default")

func main() {
	var err error

	logging.SetLevel(logging.DEBUG, "default")

	prompt1 := promptui.Prompt{Label: "Country Code", Default: "98"}
	countryCode, err := prompt1.Run()

	if err != nil {
		return
	}

	prompt2 := promptui.Prompt{Label: "Mobile"}
	mobile, err := prompt2.Run()

	if err != nil {
		return
	}

	fullMobile := "+" + countryCode + mobile

	err = userAdd(fullMobile)

	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("User Add Success")

	err = verifyCode(fullMobile)

	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Verify Code Success")
}

func userAdd(mobile string) (err error) {
	resultUserAdd := gdlib.UserAdd(mobile)

	var data map[string]interface{}

	err = json.Unmarshal([]byte(resultUserAdd), &data)

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

	log.Debug(resultUserAdd)
	return
}

func verifyCode(mobile string) (err error) {
	const label = "Verify Code"
	const version = "1.0.0"
	const company = "GAP Cli"

	prompt3 := promptui.Prompt{Label: label}
	verifyCode, err := prompt3.Run()

	if err != nil {
		return err
	}

	uuid := uuid.Must(uuid.NewV4())
	resultUser := gdlib.User(uuid.String(), version, company, verifyCode, mobile)

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
