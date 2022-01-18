package main

import "WWW/pkg"

/*
	SMSAero API: https://smsaero.ru/api/
	Отправка SMS по телефону
*/
func main() {
	dictPhone := make(map[string]string)
	dictPhone["79963567212"] = "Денис"
	dictPhone["79044222171"] = "Андрей"
	dictPhone["79033158691"] = "Сергей"
	dictPhone["79047795408"] = "Даша"

	pkg.Run(dictPhone)
}
