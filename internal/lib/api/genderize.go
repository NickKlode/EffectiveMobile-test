package api

import (
	"emobletest/internal/storage/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetGender(name string) (string, error) {
	gender := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	var g model.User
	// if name == "" {
	// 	return "", errors.New("enter user data")
	// }
	d, err := http.Get(gender)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(d.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &g)
	if err != nil {
		return "", err
	}
	return g.Gender, nil
}
