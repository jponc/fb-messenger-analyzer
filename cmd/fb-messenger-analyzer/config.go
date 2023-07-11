package main

import "flag"

type Config struct {
	InDir   string
	OutFile string
}

func NewConfig() *Config {
	inDir := flag.String("in-dir", "", "directory which contains the FB JSON files")
	outFile := flag.String("out-csv", "", "output CSV file")

	flag.Parse()

	return &Config{
		InDir:   *inDir,
		OutFile: *outFile,
	}
}
