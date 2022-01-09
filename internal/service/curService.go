package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cookienyancloud/avito-backend-test/internal/domain"
)

//go:generate mockgen -source=curService.go -destination=mocks/curServiceMock.go


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

//currency interface
type ICurrency interface {
	GetCur(cur string, sum float64) (string, error)
}

//api currency
func (u curService) GetCur(cur string, sum float64) (string, error) {
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
	var jsondata domain.CurrencyResponse
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
