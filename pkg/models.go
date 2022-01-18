package pkg

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//request запрос API SMS Aero
//	phone - Номер телефона
//	message - Текст сообщения
func (m *Account) request(phone, message string) {
	value := url.Values{
		"number": {phone},
		"text":   {message},
		"sign":   {"SMS Aero"},
	}

	req, err := http.PostForm("https://"+m.Email+":"+m.Token+"@gate.smsaero.ru/v2/sms/send", value)
	if err != nil || req.StatusCode != 200 {
		log.Fatalf("ошибка запроса: %v", req.Status)
	}

	m.Chan <- req
}

//readStatus вывод в терминал о статусе запроса
func (m *Account) readStatus() (message string) {
	res := <-m.Chan
	return fmt.Sprintf("HTTP Status: %v", res.Status)
}

//writeMessage текст SMS для Абонента
//	 name - Имя абонента
func writeMessage(name string) (message string) {
	return fmt.Sprintf("Привет %v. Меня зовут Клыков Денис, если Вы получили это сообщения, пожалуйста дайте знать об этом. Пишу программу для отправки SMS... Спасибо большое!!!", name)
}
