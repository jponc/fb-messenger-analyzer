package csvgenerator

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jponc/fb-messenger-analyzer/internal/types"
	log "github.com/sirupsen/logrus"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateCSV(callRecords *[]types.CallRecord, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create outfile (%s): %v", outFile, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// headers
	headers := []string{"Start Time", "End Time", "Duration Seconds", "Duration Minutes"}
	err = writer.Write(headers)
	if err != nil {
		return fmt.Errorf("failed to write headers: %v", err)
	}

	for _, record := range *callRecords {
		startTime := time.Unix(int64(record.StartMS/1000), 0)
		endTime := time.Unix(int64(record.EndMS/1000), 0)
		minutes := record.DurationSeconds / 60

		err := writer.Write([]string{
			startTime.Format(time.RFC3339),
			endTime.Format(time.RFC3339),
			strconv.Itoa(record.DurationSeconds),
			strconv.Itoa(minutes),
		})
		if err != nil {
			return fmt.Errorf("failed to write csv data: %v", err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return fmt.Errorf("error in flushing csv writer: %v", err)
	}

	log.Info("successfully added entries to csv")

	return nil
}
