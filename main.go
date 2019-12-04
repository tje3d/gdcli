package main

import (
	"gdlib/gdlib"

	"github.com/manifoldco/promptui"
	"github.com/op/go-logging"
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

	err = gdlib.DoUserAdd(fullMobile)

	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("User Add Success")

	err = gdlib.DoUser(fullMobile)

	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Info("Verify Code Success")
}
