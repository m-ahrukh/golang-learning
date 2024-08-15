package main

import (
	"goLangLearning/smartystreets/processor"
	"log"
	"net/http"
	"os"
)

func main() {
	client := processor.NewAuthenticationClient(
		http.DefaultClient, "https", "us-street.api.smartystreets.com",
		"52bfb315-cec6-3fe1-682e-1900977e24f6", "RfDPLeLpY06mChUSTpbB")

	pipeline := processor.NewPipeline(os.Stdin, os.Stdout, client, 8)

	if err := pipeline.Process(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
