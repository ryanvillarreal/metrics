package file

import (
	// internals
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	// ext libs
	"github.com/ryanvillarreal/metrics/core/db"
	"github.com/tomsteele/go-nmap"
)

// splitting this into ReadXMLFiles allows us to perform preprocessing on the xml before ingestion
func ReadXMLFiles(dir string) error {
	if dir == "" {
		dir = "./data/"
	}
	files, err := os.ReadDir(dir)

	// handle no dir
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// handle no xml files
	if len(files) <= 0 {
		return fmt.Errorf("no XML files present")
	}

	// parse each
	for _, file := range files {
		// maybe someday we will support other extensions but for now .xml
		if filepath.Ext(file.Name()) == ".xml" {
			filePath := filepath.Join(dir, file.Name())
			fmt.Printf("Processing file: %s\n", filePath)

			f, err := os.Open(filePath)

			// handle file io errs
			if err != nil {
				log.Printf("failed to open file %s: %v", filePath, err)
				continue
			}
			defer f.Close()

			// nmap.Parse needs byte array
			content, err := io.ReadAll(f)
			if err != nil {
				log.Printf("failed to read file %s: %v", filePath, err)
				continue
			}

			// try the nmap.parse with the contents of the current
			nmapRun, err := nmap.Parse(content)
			//TODO: better error handling of the nmapRun - we could probably to more here
			if err != nil {
				log.Printf("failed to parse XML content from %s: %v", filePath, err)
				continue
			}

			// send the full nmapRun obj to the db ingestor
			err = db.StoreMetrics(nmapRun)
			if err != nil {
				log.Fatal("ruh roh raggy")
			}
		}
	}

	return nil
}
