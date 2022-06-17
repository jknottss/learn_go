package main

import "unicode/utf8"

func main() {
	//по умолчанию создастся пустая строка
	var str string

	//со спец символами
	var hello string = "Привет\n\t"

	//без спец символов
	var world string = `Мир\n\t`

	//UTF-8 из коробки
	hi := "你好, 世界"

	//для байт одинарные кавычки (uint8)
	var rewBinary byte = '\x27'

	//rune (uint32) для UTF-8 символов
	var someChinese rune = '界'

	helloWorld := "Привет Мир"
	//конкатенация строк
	andGoodMorning := helloWorld + "и доброе утро!"

	//строки неизменяемы
	//helloWorld[0] = 72 - не сработает

	//получение длины строки
	byteLen := len(helloWorld)                    //19 байт
	symbols := utf8.RuneCountInString(helloWorld) //10 рун

	//получение подстроки в байтах, не символах
	hello := helloWorld[:12] //Привет, 0-11 байт
	H := helloWorld[0]       //byte 72

	//конвертация в слайс байт и обратно
	byteString := []byte(helloWorld)
	helloWorld = string(byteString)
}
