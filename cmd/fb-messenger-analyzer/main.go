package main

import (
	"flag"

	"github.com/jponc/fb-messenger-analyzer/internal/csvgenerator"
	"github.com/jponc/fb-messenger-analyzer/internal/parser"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := NewConfig()
	flag.Parse()

	parserClient := parser.NewParserclient()
	generatorClient := csvgenerator.NewGenerator()

	callRecords, err := parserClient.ParseCalls(config.InDir)
	if err != nil {
		log.Fatalf("failed to parse the message calls: %v", err)
	}

	err = generatorClient.GenerateCSV(callRecords, config.OutFile)
	if err != nil {
		log.Fatalf("failed to generate csv: %v", err)
	}
}
