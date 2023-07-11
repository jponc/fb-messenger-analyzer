package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jponc/fb-messenger-analyzer/internal/types"
	log "github.com/sirupsen/logrus"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(inFile string, outFile string, senderName string) error {
	jsonFile, err := os.Open(inFile)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %v", err)
	}
	defer jsonFile.Close()

	log.Infof("succesfully opened %s", inFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var fbData types.FBData

	err = json.Unmarshal(byteValue, &fbData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal fb data: %v", err)
	}

	return nil
}
