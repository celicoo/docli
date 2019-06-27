package cmd

import (
	"fmt"
	"log"

	"github.com/celicoo/docli"
)

type Docli struct {
	Create Create
	Add    Add
}

func (d *Docli) Doc() string {
	return `Docli is a declarative language for describing command-line interfaces
in Go programs. It cuts down the boilerplate to the bare minimum, so you can move
on to the reason why you’re making the app in the first place.

Usage: docli <command>

Available Commands:

  create     Create new CLI app
  add        Add command to existing CLI app

Use "docli <command> help" for more information about the <command>.`
}

func (d *Docli) Run() {
	d.Help()
}

func (d *Docli) Help() {
	fmt.Println(d.Doc())
}

func (d *Docli) Error(err error) {
	log.Fatal(err)
}

func Execute() {
	var d Docli
	args := docli.Args()
	args.Bind(&d)
}
