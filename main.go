package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type SMSAero interface {
	request(phone, message string)
	writeStatus() (message string)
}

//Account структура
type Account struct {
	Email string
	Token string
	Chan  chan *http.Response
}

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

//writeStatus вывод на экран
func (m *Account) writeStatus() (message string) {
	res := <-m.Chan
	return fmt.Sprintf("HTTP Status: %v", res.Status)
}

/*
	SMSAero API: https://smsaero.ru/api/
	Отправка SMS по телефону
*/
func main() {
	var SMS = &Account{
		Email: "...", // Логин учетной записи
		Token: "...", // Токен учетной записи
		Chan:  make(chan *http.Response),
	}

	arrayPhone := [4]string{"79963567212", "79044222171", "79033158691", "79047795408"}

	for _, ph := range arrayPhone {
		go SMS.request(ph, "Привет! Меня зовут Клыков Денис, если Вы получили это сообщения, пожалуйста дайте мне знать об этом. Пишу программу отправка SMS... Спасибо большое!!!")
	}

	fmt.Println(SMS.writeStatus())
}
