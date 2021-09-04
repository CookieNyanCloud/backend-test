package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseCurURL = "http://api.exchangeratesapi.io/v1/latest?access_key=%s&symbols=RUB,%s"


type CurService struct {
	ApiKey string
}

type CurrencyResponse struct {
	Success bool `json:"success"`
	Timestamp int64 `json:"timestamp"`
	Base string `json:"base"`
	Date string `json:"date"`
	Rates map[string]interface{} `json:"rates"`

}

type rates struct {
	RUB float64 `json:"rub"`
	NEW map[string]interface{} `json:"-"`
}



type Currency  interface {
	GetCur(ctx context.Context, cur string, sum float64) (string, error)
}

func (u CurService) GetCur (ctx context.Context, cur string, sum float64) (string, error) {

	querry:= fmt.Sprintf(baseCurURL,u.ApiKey,cur)
	res, err:=http.Get(querry)
	if err!=nil{
		return "",err
	}
	data,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		return "",err
	}
	err = res.Body.Close()
	if err!=nil{
		return "",err
	}

	var jsondata CurrencyResponse
	err = json.Unmarshal(data, &jsondata)
	if err!=nil{
		return "",err
	}

	fmt.Printf("%#v\n",jsondata.Rates[cur])

	return "", nil
}