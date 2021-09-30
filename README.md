# Тестовое задание

## Запуск
- Общие данные(для базы данных указаны postgres postgres postgres основные поля) берутся из файла конфигураций /configs/main.yml
- создание таблиц происходит на стороне сервера в соответствии с схемой
- ниже представленные запросы выполнялись в Postman
- не совсем механизм миграций, инициализация бд с пустыми таблицами в соответствии с схемой происходит на стороне сервера при инициализации бд
### Локально
- Пароль к базе дынных и ключ к api для выдачи курса валют в .env файле
```dotenv
POSTGRES_PASSWORD=
API_KEY=
host=db
 ```
- для локального запуска `make run` (go run cmd/main.go -local флаг позволяет отличить хост и его источник)
### Docker
- запись хоста в .env выше для докера, как и пароль, остальные данные из файла конфигураций
- `make up` (docker-compose up backend-test) для запуска докера

## Описание, вопросы и решения
- в качестве id использован uuid ради уникальности, хоть и оверкил
- decimal и float64 в бд и структурах соответственно
# начисление средств
- POST `http://localhost:8090/api/v1/operation/transaction`
- на входе обязательные id (uuid),сумма (десятичная дробь), вариативно описание(до 20 символов)
```json
{
    "id":"a8887f18-b68e-4999-9c33-cc8ecbdf8c2c",
    "sum":100.5,
    "description":"tests"
}
```
- так как в базе данных не должно быть данных изначально, при отсутствии id запись создается
- возвращает сообщение об успехе операции
```json
{
    "message": "удачная транзакция"
}
```

# перевод
- POST `http://localhost:8090/api/v1/operation/remittance`
- на входе обязательные id отправителя и получателя (uuid),сумма (положительная), вариативно описание до 20 символов
```json
{
    "id_from":"a8887f18-b68e-4999-9c33-cc8ecbdf8c2c",
    "id_to":"bc5f99f1-6808-4631-9eb0-e99f51e69bc8",
    "sum":200000,
    "description":"tests"
}
```
- создается только получатель, при отсутствии отправителя выдается ошибка об уходе в минус
- возвращает сообщение об успехе операции
```json
{
    "message": "недостаточно средств"
}
```


# баланс
- GET `http://localhost:8090/api/v1/operation/balance`
- GET `http://localhost:8090/api/v1/operation/balance?currency=USD`
- на входе обязательно id(uuid), в параметрах указывается валюта(по умолчанию рубли)
```json
{
    "id":"a8887f18-b68e-4999-9c33-cc8ecbdf8c2c"
}
```
- возвращает баланс и валюту
- так как https://exchangeratesapi.io/ в качестве бесплатной базовой валюты предоставляет только евро, перевод в курсы к рублям осуществлен отношением ЕВРО-НЕОБХОДИМАЯ/ЕВРО-РУБЛИ

```json
{
    "balanceResponse": "₽520.50",
    "cur": "RUB"
}
```
```json
{
    "balanceResponse": "$7.10",
    "cur": "USD"
}
```

# транзакции
- GET `http://localhost:8090/api/v1/operation/transactionsList`
- GET `http://localhost:8090/api/v1/operation/transactionsList?sort=sum&dir=asc&page=1`
- на входе обязательно id(uuid), в параметрах указывается поле сортировки, направление и страница(по умолчанию дата по возрастанию, страница включает в себя до 5 записей в json)
```json
{
    "id":"a8887f18-b68e-4999-9c33-cc8ecbdf8c2c"
}
```
- возвращает список транзакций

```json
[
    {
        "id": "a8887f18-b68e-4999-9c33-cc8ecbdf8c2c",
        "operation": "transaction",
        "sum": -110,
        "date": "2021-09-08T17:28:32.033618Z"
    },
    {
        "id": "a8887f18-b68e-4999-9c33-cc8ecbdf8c2c",
        "operation": "transaction",
        "sum": 10,
        "date": "2021-09-08T17:28:25.044972Z"
    },
    {
        "id": "a8887f18-b68e-4999-9c33-cc8ecbdf8c2c",
        "operation": "remittance",
        "sum": 10,
        "date": "2021-09-08T17:29:03.243682Z",
        "id_to": "bc5f99f1-6808-4631-9eb0-e99f51e69bc8"
    },
    {
        "id": "a8887f18-b68e-4999-9c33-cc8ecbdf8c2c",
        "operation": "transaction",
        "sum": 45,
        "date": "2021-09-08T17:40:11.272648Z",
        "description": "test"
    },
    {
        "id": "a8887f18-b68e-4999-9c33-cc8ecbdf8c2c",
        "operation": "transaction",
        "sum": 110,
        "date": "2021-09-08T17:28:28.052397Z"
    }
]
```

## схема
![Schema](https://i.ibb.co/WKy1r5w/avito.png)



## Для связи
- [telegram - t.me/cookienyancloud](t.me/cookienyancloud)
- [почта - emil8yunusov@gmail.com](emil8yunusov@gmail.com)




 
 
 
