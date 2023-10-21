package api

import (
	"emobletest/internal/storage/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAge(name string) (int, error) {
	age := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	var a model.User
	// if name == "" {
	// 	return 0, errors.New("enter user data")
	// }
	d, err := http.Get(age)
	if err != nil {
		return 0, err
	}
	body, err := io.ReadAll(d.Body)
	if err != nil {
		return 0, err
	}
	_ = json.Unmarshal(body, &a)
	if err != nil {
		return 0, err
	}

	return a.Age, nil
}
