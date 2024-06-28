package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/deemount/goFetchNasaEpicImage/internal/epic"
)

func main() {
	dateStr := flag.String("date", "", "Date in YYYY-MM-DD format (optional)")
	targetFolder := flag.String("target", "", "Target folder to store images")
	flag.Parse()

	if *targetFolder == "" {
		log.Fatal("Target folder is required")
	}

	// Fetch the API key from the environment variable
	apiKey := os.Getenv("NASA_EPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("NASA_EPIC_API_KEY environment variable is not set")
	}

	// Set the API key in the epic package
	epic.SetAPIKey(apiKey)

	var date time.Time
	var err error
	if *dateStr == "" {
		date, err = epic.GetLatestDate()
		if err != nil {
			log.Fatalf("Failed to get latest date: %v", err)
		}
	} else {
		date, err = time.Parse("2006-01-02", *dateStr)
		if err != nil {
			log.Fatalf("Invalid date format: %v", err)
		}
	}

	subfolder := filepath.Join(*targetFolder, date.Format("2006-01-02"))
	err = os.MkdirAll(subfolder, 0755)
	if err != nil {
		log.Fatalf("Failed to create subfolder: %v", err)
	}

	err = epic.DownloadImages(date, subfolder)
	if err != nil {
		log.Fatalf("Failed to download images: %v", err)
	}

	fmt.Printf("Images downloaded to %s\n", subfolder)
}
