package main

import "fmt"

type HeadPhones struct {
}

func (hp *HeadPhones) InsertMiniJack(conn MiniJackConnector) {
	fmt.Println("Подключаю наушники через штекер Mini Jack")
	conn.PlugMiniJack()
}

type MiniJackConnector interface {
	PlugMiniJack()
}

type MiniJack struct{}

func (mjc *MiniJack) PlugMiniJack() {
	fmt.Println("Штекер Mini Jack подключен к разъёму Mini Jack")
}

type Jack struct{}

func (mjc *Jack) PlugJack() {
	fmt.Println("Штекер Jack подключен к разъёму Jack")
}

type JackAdapter struct {
	Jack *Jack
}

func (mjc *JackAdapter) PlugMiniJack() {
	fmt.Println("Штекер Mini Jack подключен к адаптеру Jack")
	mjc.Jack.PlugJack()
}

func main() {
	headPhones := HeadPhones{}
	miniJackConn := MiniJack{}
	jackConn := Jack{}

	jackAdapter := JackAdapter{Jack: &jackConn}

	fmt.Println("Пробуем подключить наушники к Mini Jack напрямую...")
	headPhones.InsertMiniJack(&miniJackConn)

	print("\n")

	fmt.Println("Пробуем подключить наушники к Jack через адаптер...")
	headPhones.InsertMiniJack(&jackAdapter)
}
