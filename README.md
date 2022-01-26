# Commercial SMS messages
Отправка SMS на телефон... Образец кода. SMS Aero https://smsaero.ru/

***

### 1. Описание флагов:
- `--help` Подробное описания парметров (Обязательно к просмотру)
- `--user` Email от учетной записи https://smsaero.ru
- `--token` Token от учетной записи https://smsaero.ru/cabinet/settings/apikey/

### 2. Запуск утилиты:
- `go run cmd/report/main.go --user=Почта_Пользователя --token=Токен`

***

### 3. Build Docker:
- `sudo docker build -t sms .`

### 3. Run Docker:
- `sudo docker run -t -i sms ./app --user=Почта_Пользователя --token=Токен`

