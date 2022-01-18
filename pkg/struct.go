package pkg

import "net/http"

type SMSAero interface {
	request(phone, message string)
	readStatus() (message string)
}

//Account структура
type Account struct {
	Email string
	Token string
	Chan  chan *http.Response
}
