package main

import (
	"fmt"
	"os"
	"traveldeck/app"
)

func main() {
	path := "traveldeck.db"
	a, e := app.New(path)
	if e != nil {
		panic(e)
	}
	defer a.Store.Close()
	msg, e := a.Demo()
	if e != nil {
		panic(e)
	}
	fmt.Println(msg)
	_ = os.Stdout
}
