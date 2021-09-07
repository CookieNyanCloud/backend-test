# Тестовое задание на позицию стажера-бекендера

## Запуск
### Локально
- Общие данные берутся из файла конфигураций /configs/main.yml

- Пароль к базе дынных и ключь к api для выдачи курса валют в .env файле
```dotenv
POSTGRES_PASSWORD=
API_KEY=
host=db
 ```
- для локального запуска `make run` (go run cmd/main.go -local флаг позволяет отличить хост и его источник)

### Docker

## Описание, вопросы и решения
- в качестве id использован int, что, возможно, не лучшая идея, но тестить будет легче, чем uuid
- decimal и float64 в бд и структурах соответственно и по таким же соображениям
# начисление средств
- POST `http://localhost:8090/api/v1/operation/transaction`
- на входе обязательные id (целое число больше 0),сумма (десятичная дробь), вариативно описание(до 20 символов)
```json
{
    "id":1,
    "sum":888.50,
    "description":"test"
}
```
- так как в базе данных не должно быть данных изначально, про отсутвии id запись создается
- возвращает сообщение об успехе операции
```json
{
    "message": "удачная транзакция"
}
```

# перевод
- POST `http://localhost:8090/api/v1/operation/remittance`
- на входе обязательные id отправителя и получателя (целое число больше нуля),сумма (положительная), вариативно описание до 20 символов
```json
{
    "id_from":1,
    "id_to":4,
    "sum":1000000
}
```
- создается только получатель, при отсутвии отправителя выдается ошибка об уходе в минус
- возвращает сообщение об успехе операции
```json
{
    "message": "недостаточно средств"
}
```


# баланс
- GET `http://localhost:8090/api/v1/operation/balance`
- GET `http://localhost:8090/api/v1/operation/balance?currency=USD`
- на входе обязательно id(целое число), в параметрах указывается валюта(по умолчанию рубли)
```json
{
    "id":9
}
```
- возвращает баланс и валюту
- так как https://exchangeratesapi.io/ в качестве бесплатной базовой валюты предоставляет только евро, перевод в курсы к рублям осуществлен отношением ЕВРО-НЕОБХОДИМАЯ/ЕВРО-РУБЛИ

```json
{
    "balanceResponse": "₽4588.50",
    "cur": "RUB"
}
```
```json
{
    "balanceResponse": "$63.12",
    "cur": "USD"
}
```

# транзакции
- GET `http://localhost:8090/api/v1/operation/transactionsList`
- GET `http://localhost:8090/api/v1/operation/transactionsList?sort=sum&dir=asc&page=1`
- на входе обязательно id(целое число), в параметрах указывается поле сортировки, направление и страница(по умолчанию дата по возрастанию, страница включает в себя 5 записей в json)
```json
{
    "id":40
}
```
- возвращает список транзакций

```json
[
    {
        "id": 40,
        "operation": "remittance",
        "sum": 50.5,
        "date": "2021-09-06T16:37:55.075447Z",
        "id_to": "41"
    },
    {
        "id": 40,
        "operation": "transaction",
        "sum": 100,
        "date": "2021-09-06T16:36:47.147248Z",
        "description": "test"
    },
    {
        "id": 40,
        "operation": "remittance",
        "sum": 100,
        "date": "2021-09-06T16:37:34.783359Z",
        "description": "test3",
        "id_to": "41"
    },
    {
        "id": 40,
        "operation": "remittance",
        "sum": 100.5,
        "date": "2021-09-06T16:37:46.350523Z",
        "description": "test4",
        "id_to": "41"
    },
    {
        "id": 40,
        "operation": "transaction",
        "sum": 100.6,
        "date": "2021-09-06T16:36:55.398285Z",
        "description": "test2"
    }
]
```

## схема
![Schema](https://i.ibb.co/YbMpDPy/avito-schema.png)



## Для связи
- [telegram - t.me/cookienyancloud](t.me/cookienyancloud)
- [почта - emil8yunusov@gmail.com](emil8yunusov@gmail.com)





 
 
 