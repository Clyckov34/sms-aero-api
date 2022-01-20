package main

import "WWW/internal"

/*
	SMSAero API: https://smsaero.ru/api/
	Отправка SMS по телефону
*/
func main() {
	dictPhone := make(map[string]string)
	dictPhone["79963567212"] = "Денис"
	dictPhone["79996299940"] = "Маша"
	dictPhone["79996256594"] = "Света"
	dictPhone["79610857087"] = "Дима"
	dictPhone["79963568358"] = "Петя"

	internal.Run(dictPhone)
}
