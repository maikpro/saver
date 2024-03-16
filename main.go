package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/maikpro/saver/database"
	"github.com/maikpro/saver/models"
	"github.com/maikpro/saver/services"
)

func main() {
	log.Println("Saver started... :)")

	ctx := context.TODO()
	client, err := database.ConnectToMongoDB(ctx)
	if err != nil {
		log.Println("Error connecting to mongo database:", err)
		return
	}
	defer client.Disconnect(ctx)

	var imageUrl string
	flag.StringVar(&imageUrl, "imageUrl", "", "Specify the url from the image to download")

	// Parse command-line flags
	flag.Parse()

	if !isValidURL(imageUrl) {
		log.Fatalf("provided url is invalid: '%s'", imageUrl)
		return
	}

	// Access flag values
	log.Println("image url is:", imageUrl)

	// Downloading image
	log.Println("downloading the image...")
	imageData, err := services.GetFileData(imageUrl)
	if err != nil {
		log.Println("Error downloading imageData:", err)
		return
	}

	// Get the extension from image
	// Get the file extension from the URL
	ext := filepath.Ext(imageUrl)

	// Remove the leading dot from the extension
	ext = ext[1:]
	path := "./upload"
	imageName := fmt.Sprintf("img_%s.%s", uuid.New(), ext)

	fullpath, err := services.SaveFile(path, imageName, imageData)
	if err != nil {
		log.Println("Error saving image locally:", err)
		return
	}

	image := models.Image{
		Name:     imageName,
		Fullpath: fullpath,
	}

	// save filename and path to mongoDb
	err = database.Save(client, ctx, image)
	if err != nil {
		log.Println("Error saving to mongo db database:", err)
		return
	}
}

func isValidURL(str string) bool {
	// Parse the URL
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	return true
}
