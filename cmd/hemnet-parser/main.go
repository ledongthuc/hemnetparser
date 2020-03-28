package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ledongthuc/hemnetparser"
	"github.com/ledongthuc/hemnetparser/googlesheet"
)

const (
	formatJSON        = "json"
	formatXML         = "xml"
	formatGoogleSheet = "googlesheet"
)

var formatArg *string
var googleConfiguration googlesheet.Configuration

func init() {
	formatArg = flag.String("format", "", "Format support json, excel")

	flag.StringVar(&googleConfiguration.GoogleCredentialPath, "credential", "", "Credential file path from https://developers.google.com/sheets/api/quickstart/go")
	flag.StringVar(&googleConfiguration.SpreadsheetID, "sheet-id", "", "https://developers.google.com/sheets/api/guides/concepts#spreadsheet_id")
	flag.StringVar(&googleConfiguration.SheetName, "sheet-name", "Sheet1", "Sheet name (tab) of sheet file")
	flag.IntVar(&googleConfiguration.NoSkipHeader, "skipheader", 0, "number of header will be skipped. Default=0")

}

func main() {
	flag.Parse()
	var format string
	if formatArg == nil {
		format = formatJSON
	} else {
		format = *formatArg
	}

	// Google Sheet handle
	if format == formatGoogleSheet {
		b, err := googlesheet.Process(googleConfiguration)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		return
	}

	if len(os.Args) < 1 {
		log.Fatal("Missing url. hemnet-parser <url>")
	}

	url := os.Args[len(os.Args)-1]
	output, err := hemnetparser.Parse(url)
	if err != nil {
		log.Fatal(err)
	}

	var outputB []byte
	switch format {
	case formatXML:
		outputB, err = output.ToXML()
		break
	case formatJSON:
		fallthrough
	default:
		outputB, err = output.ToJSON()
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(outputB))
}
