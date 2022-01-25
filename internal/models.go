package internal

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//request запрос API SMS Aero
//	phone - Номер телефона
//	message - Текст сообщения
func (m *WaitGroup) request(phone, message string) {
	value := url.Values{
		"number": {phone},
		"text":   {message},
		"sign":   {"SMS Aero"},
	}

	req, err := http.PostForm("https://"+m.Account.Email+":"+m.Account.Token+"@gate.smsaero.ru/v2/sms/send", value)
	if err != nil || req.StatusCode != 200 {
		log.Fatalln("ошибка запроса:", req.Status)
	}

	defer func() {
		err := req.Body.Close()
		if err != nil {
			log.Println(err)
		}

		m.Wg.Done()
	}()

	status(req.Status)
}

//status выводит статус запроса SMS
func status(status string) string {
	return fmt.Sprintln("HTTP Status:", status)
}
