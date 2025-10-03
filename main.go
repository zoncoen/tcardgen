package main

import (
	"log"

	"github.com/zoncoen/tcardgen/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
