# Тестовое задание на позицию стажера-бекендера

## Запуск
### Локально
- Общие данные берутся из файла конфигураций /configs/main.yml

- Пароль к базе дынных и ключь к api для выдачи курса валют в .env файле
```dotenv
POSTGRES_PASS=
API_KEY=
 ```
- для локального запуска `make run` 

### Docker

## Описание, вопросы и решения

# начисление средств
- POST `http://localhost:8090/api/v1/operation/transaction`
- на входе обязательные id (целое число),сумма (с точностью до сотых), вариативно описание
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
- на входе обязательные id отправителя и получателя (целое число),сумма (с точностью до сотых), вариативно описание
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
- GET `http://localhost:8090/api/v1/operation/transactionsList?sort=sum&dir=asc&page=1`
- GET `http://localhost:8090/api/v1/operation/transactionsList?sort=sum&dir=asc&page=1`
- на входе обязательно id(целое число), в параметрах указывается поле сортировки, направление и страница(по умолчанию дата по возрастанию, страница включает в себя 5 записей в json)
```json
{
    "id":1
}
```
- возвращает список транзакций

```json
[
    {
        "id": 1,
        "operation": "remittance",
        "sum": 100,
        "date": "2021-09-05T21:50:34.249344Z",
        "description": "",
        "id_to": -1
    },
    {
        "id": 1,
        "operation": "remittance",
        "sum": 100,
        "date": "2021-09-05T21:50:07.371658Z",
        "description": "",
        "id_to": 21
    },
    {
        "id": 1,
        "operation": "remittance",
        "sum": 100,
        "date": "2021-09-05T21:49:47.189342Z",
        "description": "aaa",
        "id_to": 20
    },
    {
        "id": 1,
        "operation": "transaction",
        "sum": 100,
        "date": "2021-09-05T18:21:04.096842Z",
        "description": "",
        "id_to": -1
    },
    {
        "id": 1,
        "operation": "transaction",
        "sum": 200,
        "date": "2021-09-05T18:21:07.805747Z",
        "description": "",
        "id_to": -1
    }
]
```

## схема
![Schema](https://ibb.co/R7mzhYB)



## Для связи
- [telegram - t.me/cookienyancloud](t.me/cookienyancloud)
- [почта - emil8yunusov@gmail.com](emil8yunusov@gmail.com)





 
 
 