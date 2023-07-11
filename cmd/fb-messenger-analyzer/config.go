package main

import "flag"

type Config struct {
	SenderName string
	InFile     string
	OutFile    string
}

func NewConfig() *Config {
	senderName := flag.String("sender-name", "", "sender name to analyze")
	inFile := flag.String("in-file", "", "facebook JSON input file")
	outFile := flag.String("out-file", "", "output file")

	flag.Parse()

	return &Config{
		SenderName: *senderName,
		InFile:     *inFile,
		OutFile:    *outFile,
	}
}
