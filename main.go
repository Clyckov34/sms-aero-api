package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//SMSAero
type SMSAero interface {
	request(phone, message string) (resp *http.Response, err error)
}

//Account структура
type Account struct {
	Email string
	Token string
}

//request запрос API SMS Aero
//	phone - Номер телефона
//	message - Текст сообщения
func (m *Account) request(phone, message string) (resp *http.Response, err error) {
	value := url.Values{
		"number": {phone},
		"text":   {message},
		"sign":   {"SMS Aero"},
	}

	req, err := http.PostForm("https://"+m.Email+":"+m.Token+"@gate.smsaero.ru/v2/sms/send", value)
	if err != nil || req.StatusCode != 200 {
		return nil, fmt.Errorf("ошибка запроса: %v - %v", req.StatusCode, err)
	}

	return req, nil
}

/*
	SMSAero API: https://smsaero.ru/api/
	Отправка SMS по телефону
*/
func main() {
	var SMS SMSAero = &Account{
		Email: "...", // Логин учетной записи
		Token: "...", // Токен учетной записи
	}

	resp, err := SMS.request("79963567212", "Привет! Меня зовут Денис")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("HTTP Status:", resp.StatusCode)
}
