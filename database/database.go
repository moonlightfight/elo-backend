package database

import (
	"context"
	"fmt"
	"log"
	"time"

	c "github.com/moonlightfight/elo-backend/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConfigDB() (*mongo.Client, error) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if cancel != nil {
		log.Println(cancel)
	}
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.ucnph.mongodb.net/%s?retryWrites=true&w=majority", configuration.Database.DBUser, configuration.Database.DBPass, configuration.Database.DBName))
	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}
	return db, err
}
