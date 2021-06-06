package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	c "github.com/moonlightfight/elo-backend/config"
	"github.com/moonlightfight/elo-backend/database"
	m "github.com/moonlightfight/elo-backend/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func DoPasswordsMatch(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))

	return err == nil
}

func HashPassword(password string) string {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hashedPasswordBytes)
}

func CreateAdminEndpoint(response http.ResponseWriter, request *http.Request) {
	client, err := database.ConfigDB()
	if err != nil {
		log.Println(err)
	}
	response.Header().Set("content-type", "application/json")
	var admin m.Admin
	jsonErr := json.NewDecoder(request.Body).Decode(&admin)
	if jsonErr != nil {
		log.Println(err)
	}
	// encrypt user password
	admin.Password = HashPassword(admin.Password)
	collection := client.Database("test").Collection("Admin")
	ctx, ctxErr := context.WithTimeout(context.Background(), 5*time.Second)

	if ctxErr != nil {
		log.Println(ctxErr)
	}
	result, resErr := collection.InsertOne(ctx, admin)
	if resErr != nil {
		log.Println(resErr)
	}
	json.NewEncoder(response).Encode(result)
}

func AdminLoginEndpoint(response http.ResponseWriter, request *http.Request) {
	client, err := database.ConfigDB()
	if err != nil {
		log.Println(err)
	}
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("../..")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	viperErr := viper.Unmarshal(&configuration)
	if viperErr != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	response.Header().Set("content-type", "application/json")
	var loginData m.LoginData
	// retrieve request args
	_ = json.NewDecoder(request.Body).Decode(&loginData)
	var user m.Admin
	// retrieve user if exists
	collection := client.Database("test").Collection("Admin")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	dbErr := collection.FindOne(ctx, m.Admin{Email: loginData.Email}).Decode(&user)
	if dbErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + dbErr.Error() + `" }`))
		return
	}
	// check if passwords match with bcrypt
	passwordsMatch := DoPasswordsMatch(user.Password, loginData.Password)
	tokenSecret := []byte(configuration.Server.Secret)
	if !passwordsMatch {
		response.WriteHeader(http.StatusNotAcceptable)
		response.Write([]byte(`{ "message": "Invalid Password"}`))
		return
	}
	// generate jwt
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["_id"] = user.ID
	claims["iat"] = time.Now().Unix()
	tokenString, tokenErr := token.SignedString(tokenSecret)
	if tokenErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		io.WriteString(response, `{"error":"token_generation_failed"}`)
		return
	}
	// format response data
	type ResData struct {
		ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Token string             `json:"token,omitempty" bson:"token,omitempty"`
	}

	resData := ResData{
		ID:    user.ID,
		Token: tokenString,
	}

	// return data
	json.NewEncoder(response).Encode(resData)
}
