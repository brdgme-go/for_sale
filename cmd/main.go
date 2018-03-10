package main

import (
	"os"

	"github.com/brdgme-go/cmd"
	"github.com/brdgme-go/for_sale"
)

func main() {
	cmd.Cli(&for_sale.Game{}, os.Stdin, os.Stdout)
}
