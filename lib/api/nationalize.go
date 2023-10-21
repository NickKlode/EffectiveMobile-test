package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type userCountry struct {
	Country []country `json:"country"`
}
type country struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

func GetNationality(name string) (string, error) {
	nat := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	var cArr userCountry

	// if name == "" {
	// 	return "", errors.New("enter user data")
	// }

	d, err := http.Get(nat)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(d.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &cArr)
	if err != nil {
		return "", err
	}
	// первый элемент т.к. они уже отсортированы по убыванию и берется наибольшая вероятность
	cA := cArr.Country[0]

	return cA.CountryID, nil
}
