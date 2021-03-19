package config

import (
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

func Parse(c interface{}) interface{} {
	p := flags.NewParser(c, flags.Default)

	if _, err := p.Parse(); err != nil {
		log.Println("[ERROR] Error while parsing config options:", err)
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}

		os.Exit(1)
	}

	return c
}
