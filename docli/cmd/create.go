package cmd

import (
	"fmt"
	"log"
	"os"
)

type Create struct {
	Path string
}

func (c *Create) Doc() string {
	s, err := os.Getwd()
	if err != nil {
		c.Error(err)
	}

	return fmt.Sprintf(`Create new CLI app

Usage: docli create --path=<path>

Options:

  -p, --path     The path where the CLI should be created [default: %s]`, s)
}

func (c *Create) Run() {

}

func (c *Create) Help() {
	fmt.Println(c.Doc())
}

func (c *Create) Error(err error) {
	log.Fatal(err)
}
