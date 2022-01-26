# Commercial SMS messages
Отправка SMS на телефон... Образец кода. SMS Aero https://smsaero.ru/

***

### 1. Описание флагов:
- `--help` Подробное описания флагов (Обязательно к просмотру)
- `--user` Email от учетной записи https://smsaero.ru
- `--token` Token от учетной записи https://smsaero.ru/cabinet/settings/apikey/

### 2. Запуск утилиты:
- `go run cmd/report/main.go --user=Почта_Пользователя --token=Токен`

***

### 3. Docker build and run:
- `sudo docker build sms .`
- `sudo docker run sms ./app --user=Почта_Пользователя --token=Токен`

