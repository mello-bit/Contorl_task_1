# Финальный проект №1

## Описание

Этот проект реализует веб-сервис для вычисление простых математических выражений. Пользователи могут отправлять HTTP запросы с математическими выражениями в теле запроса, а сервис будет отправлять ответ, либо показывать что выражение некорректно или метод запроса был выбран неправильно(важно чтобы метод был именно POST)

## Структура проекта

```
|
|- services
|    |-- server_test.go
|    |-- server.go
|- utils
|    |-- calculation_test.go
|    |-- calculation.go
|    |-- types.go
|- go.mod
|- main.go
|- Readme.md
```

## Установка

    1. Копируете репозиторий
    ```
    git clone https://github.com/mello-bit/Contorl_task_1
    ```
    2. Переходите в папку с проектом
    ```
    cd Contorl_task_1
    ```
    3. Установка зависимостей
    ```
    go mod download
    ```

## Запуск проекта

    ```
    go run main.go
    ```

## Использование

1. Пример 1: Обычное выражение

```
curl -i -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d '{"expression":"2+2*2"}'
```

Ожидаемый ответ:

```
HTTP/1.1 200 OK
Date: Mon, 30 Dec 2024 09:38:04 GMT
Content-Length: 20
Content-Type: text/plain; charset=utf-8

{"result":"6.0000"}
```

2. Выражение со скобками:

```
curl -i -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d '{"expression":"(2+2)*2"}'
```
Ожидаемый ответ: 

```
HTTP/1.1 200 OK
Date: Mon, 30 Dec 2024 09:39:57 GMT
Content-Length: 20
Content-Type: text/plain; charset=utf-8

{"result":"8.0000"}
```

3. Выражение с пробелами:

```
curl -i -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d '{"expression":"( 2 + 2 ) * 20 "}'
```

Ожидаемый ответ:

```
HTTP/1.1 200 OK
Date: Mon, 30 Dec 2024 09:41:31 GMT
Content-Length: 21
Content-Type: text/plain; charset=utf-8

{"result":"80.0000"}
```

4. Выражение с деленим на 0:

```
curl -i -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d '{"expression":"( 2 + 2 ) / (2 - 2) "}'
```

Ожидаемый ответ:

```
HTTP/1.1 422 Unprocessable Entity
Date: Mon, 30 Dec 2024 09:44:07 GMT
Content-Length: 36
Content-Type: text/plain; charset=utf-8

{"error":"Expression is not valid"}
```

5. Запрос не метода POST:

```
curl -i -X GET http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d '{"expression":"( 2 + 2 ) / (2 - 2) "}'
```

Ожидаемый ответ:

```
HTTP/1.1 500 Internal Server Error
Date: Mon, 30 Dec 2024 09:45:29 GMT
Content-Length: 34
Content-Type: text/plain; charset=utf-8

{"error":"Internal server error"}
```

## Тестирование программы
** Для тестирования работы запросов:

```
go test ./services
```

** Для тестирования работы калькулятора:

```
go test ./utils
```