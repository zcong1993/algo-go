package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zcong1993/algo-go/cmd/tags"

	"github.com/zcong1993/algo-go/pkg/leetcode"

	"github.com/zcong1993/algo-go/cmd/new"
	"github.com/zcong1993/algo-go/cmd/update"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("algo", "A command-line tool for algo-go repo.")

	updateCmd = app.Command("update", "Update readme.")

	newCmd = app.Command("new", "Init a new leetcode problem.")
	number = newCmd.Arg("number", "problem number").Required().String()

	metaCmd    = app.Command("meta", "Show problem meta by number.")
	metaNumber = metaCmd.Arg("number", "problem number").Required().String()

	tagsCmd = app.Command("tags", "Update tag toc files.")
)

func showMeta(number string) {
	meta, err := leetcode.GetMetaByNumber(number)
	if err != nil {
		log.Fatal(err)
	}
	meta.Content = ""
	meta.Code = ""
	fmt.Printf("%+v\n", meta)
}

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case updateCmd.FullCommand():
		update.Run()
	case newCmd.FullCommand():
		new.Run(*number)
	case metaCmd.FullCommand():
		showMeta(*metaNumber)
	case tagsCmd.FullCommand():
		tags.Run()
	}
}
