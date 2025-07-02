
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create Slack client
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID := os.Getenv("CHANNEL_ID")
	fileArr := []string{"ZIPL.pdf"}

	for _, filePath := range fileArr {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			log.Fatalf("Error getting file info: %v", err)
		}

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()

		// Upload file
			params := slack.UploadFileV2Parameters{
			Channel:  channelID,
			File:     filePath,
			Filename: fileInfo.Name(),
			FileSize: int(fileInfo.Size()),
		}
		uploadedFile, err := api.UploadFileV2(params)
		if err != nil {
			log.Fatalf("Error uploading file: %v", err)
		}

		fmt.Printf("Uploaded File -> Name: %s, ID: %s\n", uploadedFile.Title, uploadedFile.ID)
	}
}
