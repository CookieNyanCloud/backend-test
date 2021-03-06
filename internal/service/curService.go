package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
)

const baseCurURL = "http://api.exchangeratesapi.io/v1/latest?access_key=%s&symbols=RUB,%s"

type curService struct {
	ApiKey string
}

//init service
func NewCurService(apiKey string) *curService {
	return &curService{ApiKey: apiKey}
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

//api currency
func (u curService) GetCur(cur string, sum float64) (string, error) {
	query := fmt.Sprintf(baseCurURL, u.ApiKey, strings.ToUpper(cur))
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
	var jsondata domain.CurrencyResponse
	err = json.Unmarshal(data, &jsondata)
	if err != nil {
		return "", err
	}
	//convert
	var userEurCur float64
	switch i := jsondata.Rates[strings.ToUpper(cur)].(type) {
	case float64:
		userEurCur = i
	}
	var userEurRub float64
	switch i := jsondata.Rates["RUB"].(type) {
	case float64:
		userEurRub = i
	}
	balanceInCur := sum * userEurCur / userEurRub
	return fmt.Sprintf("%s%.2f", curSym[strings.ToUpper(cur)], balanceInCur), nil
}
