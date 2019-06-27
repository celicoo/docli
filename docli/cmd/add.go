package cmd

import (
	"fmt"
	"log"
)

type Add struct{
	Name,
	Parent string
}

func (a *Add) Doc() string {
	return `Add command to existing CLI app

Usage: docli add --name=<name> [options]

Options:

  -n, --name     The name of the command
  -p, --parent   The name of the parent command [default: root]`
}

func (a *Add) Run() {
}

func (a *Add) Help() {
	fmt.Println(a.Doc())
}

func (a *Add) Error(err error) {
	log.Fatal(err)
}

