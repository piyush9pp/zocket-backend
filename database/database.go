package database

import (
    "context"
    "log"
    "os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
)

var DB *mongo.Database
var client *mongo.Client

func InitDB() error {
    // ✅ Load .env only in local development
    if os.Getenv("RENDER") == "" { 
        if err := godotenv.Load(); err != nil {
            log.Println("⚠️ Warning: .env file not found. Using system environment variables.")
        }
    }

    clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
    var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return err
    }

    if err = client.Ping(context.TODO(), nil); err != nil {
        return err
    }

    DB = client.Database("taskdb")
    log.Println("✅ Connected to MongoDB")
    return nil
}
