package main

import (
	"encoding/json"
	"errors"
	"gdcli/gdlib"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/op/go-logging"
	uuid "github.com/satori/go.uuid"
)

var log = logging.MustGetLogger("default")

func main() {
	var err error

	logging.SetLevel(logging.DEBUG, "default")

	countryCode := askCountryCode()

	mobile := askMobile()

	fullMobile := "+" + countryCode + mobile

	errStr := gdlib.DoUserAdd(fullMobile)

	if errStr != "" {
		log.Error(errStr)
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

func askCountryCode() string {
	prompt := promptui.Prompt{Label: "Country Code", Default: "98"}
	code, err := prompt.Run()

	if err != nil {
		os.Exit(0)
		return ""
	}

	return code
}

func askMobile() string {
	prompt2 := promptui.Prompt{Label: "Mobile"}
	mobile, err := prompt2.Run()

	if err != nil {
		os.Exit(0)
		return ""
	}

	if len(mobile) != 10 {
		log.Info("Mobile length should be 10")
		return askMobile()
	}

	return mobile
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
