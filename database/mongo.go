package database

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/maikpro/saver/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB(ctx context.Context) (*mongo.Client, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file:", err)
		return nil, err
	}

	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	if connectionString == "" {
		log.Println("MONGODB_CONNECTION_STRING environment variable is not set")
		return nil, err
	}

	// Use the connection string to establish the MongoDB connection
	// Example code to connect to MongoDB using the connection string
	log.Println("Connecting to MongoDB:", connectionString)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

func getEnvString(str string) (string, error) {
	envString := os.Getenv(str)
	if envString == "" {
		log.Printf("%s environment variable is not set or cannot be found!", str)
		return "", errors.New("environment variable is not set or cannot be found!")
	}
	return envString, nil
}

func getCollection(client *mongo.Client) (*mongo.Collection, error) {
	databaseString, err := getEnvString("MONGODB_DATABASE_STRING")
	if err != nil {
		return nil, err
	}
	collectionString, err := getEnvString("MONGODB_COLLECTION_STRING")
	if err != nil {
		return nil, err
	}
	return client.Database(databaseString).Collection(collectionString), nil
}

func Save(client *mongo.Client, ctx context.Context, image models.Image) error {
	collection, err := getCollection(client)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(ctx, image)
	if err != nil {
		return err
	}

	log.Println(`saved image to MongoDB database`, image)
	return nil
}

/* func SaveImage(client *mongo.Client, ctx context.Context, imageData []byte) error {
	// Encode the image data as Base64
	// base64Data := base64.StdEncoding.EncodeToString(imageData)

	collection, err := getCollection(client)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(ctx, bson.M{"imageData": imageData})
	if err != nil {
		return err
	}

	log.Printf(`saved fileData %s to MongoDB database`, base64Data)
	return nil
}

func ReadImage(client *mongo.Client) error {
	collection, err := getCollection(client)
	if err != nil {
		return err
	}

	var result mongo.SingleResult
	err = collection.FindOne(context.Background(), bson.M{}).Decode(&result)
	if err != nil {
		log.Fatal("Error querying image from MongoDB:", err)
	}

	// Decode the Base64-encoded image data
	imageData, err := base64.StdEncoding.DecodeString(result.)
	if err != nil {
		log.Fatal("Error decoding Base64 data:", err)
	}

	fmt.Println("Image data:", string(imageData))
	return err
} */
