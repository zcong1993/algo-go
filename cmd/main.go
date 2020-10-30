package main

import (
	"os"

	"github.com/zcong1993/algo-go/cmd/new"
	"github.com/zcong1993/algo-go/cmd/update"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("algo", "A command-line tool for algo-go repo.")

	updateCmd = app.Command("update", "Update readme.")

	newCmd = app.Command("new", "Init a new leetcode problem.")
	number = newCmd.Arg("number", "problem number").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case updateCmd.FullCommand():
		update.Run()
	case newCmd.FullCommand():
		new.Run(*number)
	}
}
