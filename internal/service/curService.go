package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseCurURL = "http://api.exchangeratesapi.io/v1/latest?access_key=%s&symbols=RUB,%s"

type CurService struct {
	ApiKey string
}

func NewCurService(apiKey string) ICurrency {
	return &CurService{ApiKey: apiKey}
}

type CurrencyResponse struct {
	Success   bool                   `json:"success"`
	Timestamp int64                  `json:"timestamp"`
	Base      string                 `json:"base"`
	Date      string                 `json:"date"`
	Rates     map[string]interface{} `json:"rates"`
}

var (
	curSym = map[string]string{
		"RUB": "₽",
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"JPY": "¥",
	}
)

type ICurrency interface {
	GetCur(cur string, sum float64) (string, error)
}

func (u CurService) GetCur(cur string, sum float64) (string, error) {
	query := fmt.Sprintf(baseCurURL, u.ApiKey, cur)
	res, err := http.Get(query)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	err = res.Body.Close()
	if err != nil {
		return "", err
	}
	var jsondata CurrencyResponse
	err = json.Unmarshal(data, &jsondata)
	if err != nil {
		return "", err
	}
	//convert
	var userEurCur float64
	switch i := jsondata.Rates[cur].(type) {
	case float64:
		userEurCur = i
	}
	var userEurRub float64
	switch i := jsondata.Rates["RUB"].(type) {
	case float64:
		userEurRub = i
	}
	balanceInCur := sum * userEurCur / userEurRub
	return fmt.Sprintf("%s%.2f", curSym[cur], balanceInCur), nil
}
