package main

import (
	"log"
	"os"

	"github.com/amPerl/pilw/cmd/pilw/commands"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func main() {
	apiKey := os.Getenv("PILW_API_KEY")
	if len(apiKey) < 1 {
		log.Fatal("Missing PILW_API_KEY from env")
	}
	viper.Set("key", apiKey)

	if err := commands.PilwCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
