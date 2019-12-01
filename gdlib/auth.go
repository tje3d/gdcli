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
