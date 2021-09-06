package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//прослойка для связи с сервисом получения курса

const baseCurURL = "http://api.exchangeratesapi.io/v1/latest?access_key=%s&symbols=RUB,%s"

type CurService struct {
	ApiKey string
}

//структура ответа
type CurrencyResponse struct {
	Success   bool   `json:"success"`
	Timestamp int64  `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date"`
	//заранее неизвестны поля курса
	Rates map[string]interface{} `json:"rates"`
}

//возможные популярные валюты и их символы
var (
	curSym = map[string]string{
		"RUB": "₽",
		"USD": "$",
		"EUR": "€",
		"GBP": "£",
		"JPY": "¥",
	}
)

type Currency interface {
	GetCur(cur string, sum float64) (string, error)
}

func (u CurService) GetCur(cur string, sum float64) (string, error) {
	//создание запроса
	querry := fmt.Sprintf(baseCurURL, u.ApiKey, cur)
	res, err := http.Get(querry)
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
	//запись ответа в структуру
	var jsondata CurrencyResponse
	err = json.Unmarshal(data, &jsondata)
	if err != nil {
		return "", err
	}
	//вывод значений интерфейса по ключу в виде float64
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
	//нахождение необходимого курса по отношению двух курсов базовой валюты
	balanceInCur := sum * userEurCur / userEurRub
	return fmt.Sprintf("%s%.2f", curSym[cur], balanceInCur), nil
}
