package admin

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	c "github.com/moonlightfight/elo-backend/config"
	m "github.com/moonlightfight/elo-backend/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

func DoPasswordsMatch(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err != nil
}

func HashPassword(password string) string {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	var base64PasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64PasswordHash
}

func CreateAdminEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var admin m.Admin
	_ = json.NewDecoder(request.Body).Decode(&admin)
	// encrypt user password
	admin.Password = HashPassword(admin.Password)
	collection := client.Database("").Collection("Admin")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, admin)
	json.NewEncoder(response).Encode(result)
}

func AdminLoginEndpoint(response http.ResponseWriter, request *http.Request) {
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
	response.Header().Set("content-type", "application/json")
	var loginData m.LoginData
	// retrieve request args
	_ = json.NewDecoder(request.Body).Decode(&loginData)
	var user m.Admin
	// retrieve user if exists
	collection := client.Database("").Collection("Admin")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	dbErr := collection.FindOne(ctx, m.Admin{Email: loginData.Email}).Decode(&user)
	if dbErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	// check if passwords match with bcrypt
	passwordsMatch := DoPasswordsMatch(user.Password, loginData.Password)
	if !passwordsMatch {
		response.WriteHeader(http.StatusNotAcceptable)
		response.Write([]byte(`{ "message": "Invalid Password"}`))
		return
	}
	// generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"_id":      user.ID,
		"iat":      time.Now().Unix(),
	})
	tokenString, tokenErr := token.SignedString([]byte(configuration.Server.Secret))
	if tokenErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		io.WriteString(response, `{"error":"token_generation_failed"}`)
		return
	}
	// format response data
	type ResData struct {
		ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		token string             `json:"token,omitempty" bson:"token,omitempty"`
	}

	resData := ResData{
		ID:    user.ID,
		token: tokenString,
	}

	// return data
	json.NewEncoder(response).Encode(resData)
}
