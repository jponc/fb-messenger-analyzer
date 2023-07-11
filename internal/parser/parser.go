package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/jponc/fb-messenger-analyzer/internal/types"
)

type ParserClient struct{}

func NewParserclient() *ParserClient {
	return &ParserClient{}
}

func (p *ParserClient) ParseCalls(inDir string) (*[]types.CallRecord, error) {
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read files from directory %s: %v", inDir, err)
	}

	records := []types.CallRecord{}

	for _, file := range files {
		isJSONFile := strings.Contains(file.Name(), "json")
		if !isJSONFile {
			continue
		}

		path := fmt.Sprintf("%s/%s", inDir, file.Name())

		jsonFile, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open %s: %v", path, err)
		}
		defer jsonFile.Close()

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %v", path, err)
		}

		var fbFileData types.FBFileData
		err = json.Unmarshal(byteValue, &fbFileData)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal fb file data: %v", err)
		}

		for _, message := range fbFileData.Messages {
			if message.CallDurationSeconds == 0 {
				continue
			}

			startMS := message.TimestampMS - (message.CallDurationSeconds * 1000)

			records = append(records, types.CallRecord{
				StartMS:         startMS,
				EndMS:           message.TimestampMS,
				DurationSeconds: message.CallDurationSeconds,
			})
		}
	}

	// Sort the records
	sort.Slice(records, func(i, j int) bool {
		return records[i].StartMS < records[j].StartMS
	})

	return &records, nil
}
