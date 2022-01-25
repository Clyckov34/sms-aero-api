package main

import (
	"flag"
	"log"
	"mysmsproject/internal"
	"os"
)

var user internal.User

func init() {
	flag.StringVar(&user.Email, "user", "", "Email от учетной записи https://smsaero.ru")
	flag.StringVar(&user.Token, "token", "", "Token от учетной записи https://smsaero.ru/cabinet/settings/apikey/")
}

/*
	SMSAero API: https://smsaero.ru/api/
	Отправка SMS по телефону
*/
func main() {
	listPhone := []int{79963567212, 79996299940, 79996256594, 79610857087, 79963568358}

	if len(os.Args) == 1 {
		log.Fatalln("ошибка: Не указаны флаги")
	} else {
		flag.Parse()
		internal.SMSRun(listPhone, &user)
	}
}
