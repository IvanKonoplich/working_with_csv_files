# О проекте
___
Телеграм бот, который принимает сообщения и отправляет их же обратно.
Проект создан, чтобы научиться взаимодействовать с телеграм api.
`t.me/EchoBotMadeByIvanKonoplich_bot.`

## Запуск

Проект запускается через:
```
go run main.go
```
По умолчанию работает на порту 8080.
В проекте не подключена БД, вместо нее сервис который имитирует термометр и
выдает случайную температуру.
Значение кэшируется и обновляется каждые 10 секунд.
## Запрос через Postman
* Узнать текущую температуру:
  ```
  GET
  ```
  ```
  http://localhost:8080/get_weather
  ```
  Пример ответа (json):
  ```json
  {
  "temperature": 91.446754
  }
  ```

## Contact

Ivan Konoplich - konoplich_i@mail.ru

Project Link: [https://github.com/IvanKonoplich/Weather-Service.git](https://github.com/IvanKonoplich/Weather-Service.git)

