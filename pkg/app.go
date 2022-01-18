package pkg

import (
	"fmt"
	"net/http"
)

func Run(dictPhone map[string]string) {
	var SMS = &Account{
		Email: "...", // Логин учетной записи
		Token: "...", // Токен учетной записи
		Chan:  make(chan *http.Response),
	}

	for phone, name := range dictPhone {
		go SMS.request(phone, SMS.messageSMS(name))
	}

	fmt.Println(SMS.writeStatus())
}
