только на Go
# Тестовое задание на Python Backend разработчика 

Бекенд на fastapi/drf для новостных статей:
- Автор (обьект некий) будет создавать статьи
  - В таблице авторов будет хранится только уникальный username и пароль, авторизация в сервис реализуется по jwt токенам
- Статья
  - В таблице статей будет хранится title (заголовок), content (сам основной контент) и сам автор

## Детали
Эндпоинт `/all` где должна выводится вся информация об авторах и их статьях в такой структуре сериализатора:
```json
[
  {
    "username": "michael",
    "articles": [
      {
        "Title": "hello",
        "Content": "i love rust"
      }
    ]
  },
  {
    "username": "kamran",
    "articles": [
      {
        "Title": "Uzbekistan",
        "Content": "tashkent is the capital of uzbekistan"
      },
      {
        "Title": "London",
        "Content": "london is the capital of Great Britain"
      }
    ]
  }
]
```

__Ограничения: Заголовок должен быть больше 3 символов и меньше 100, и при этом в заголовке и в контенте должны быть только *БУКВЫ*, без цифр и без др спец символов__

В имени пользователя должны быть только латинские буквы. Если нарушено ограничение сервер пусть выдает ошибку с причиной этой ошибки с правильным статус кодом для этих ошибок, ты должен сам понять каким для таких случаев

### Дополнительные требования
- Закрыть все это дело тестами, ассерт статус кода и работы json ответа

__Данное тестовое задание нацеленно на правильный подход, хорошему коду и хорошей структуре__



# JWT Auth Gybernaty

## Требования
- Docker
- Docker Compose
- Postman или любой другой аналог

## Установка и настройка

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/Seven11Eleven/jwt_auth_gybernaty.git
   cd jwt_auth_gybernaty
   ```

2. Соберите и запустите docker-контейнеры
  ```bash
docker compose up --build
  ```

## проект сам поднимется, поднимется база данных с уже готовыми таблицами для работы
## http сервис поднимется на 8080 порту

## Как опробовать?

## API Эндпоинты
# 1. Регистрация пользователя

    Эндпоинт: POST /signup

    Тело запроса:

```json
{
    "Username": "kamran",
    "Password": "seven1337"
}
```
Команда cURL:

```bash

    curl -X POST http://localhost:8080/signup \
    -H "Content-Type: application/json" \
    -d '{"Username": "kamran", "Password": "seven1337"}'
```
# 2. Логин пользователя

    Эндпоинт: POST /login

    Тело запроса:

```json

{
    "Username": "kamran",
    "Password": "seven1337"
}
```
Команда cURL:

```bash

curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"Username": "kamran", "Password": "seven1337"}'
```
Ответ:

```json

    {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImthbXJhbiIsImlkIjoiNzNkMDk3MTMtMjZlZS00ZGNmLWExZTctY2YxMjcwYzNmZDIwIiwiZXhwIjoxNzIzNjUzMjc4fQ.aXYCv-3eREqH_2rAqxMJbTXBEe1gVEhc4O5xhVYbKs0",
        "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjczZDA5NzEzLTI2ZWUtNGRjZi1hMWU3LWNmMTI3MGMzZmQyMCIsImV4cCI6MTcyNDM2NjA3OH0.FIgIRUdITIGsHgmgkGiOEaUfD60iLR9hNvJ07Jj_nNo"
    }
```
# 3. Получение всех авторов и статей

    Эндпоинт: GET /all

    Заголовок:
        Authorization: Bearer <accessToken>

    Команда cURL:

```bash

    curl -X GET http://localhost:8080/all \
    -H "Authorization: Bearer <accessToken>"
```
# 4. Создание новой статьи

    Эндпоинт: POST /article

    Заголовок:
        Authorization: Bearer <accessToken>

    Тело запроса:

``` json

{
    "title": "how to write python way",
    "content": "idk i am go dev"
}
```
Команда cURL:

```bash

curl -X POST http://localhost:8080/article \
-H "Authorization: Bearer <accessToken>" \
-H "Content-Type: application/json" \
-d '{"title": "how to write python way", "content": "idk i am go dev"}'
```
