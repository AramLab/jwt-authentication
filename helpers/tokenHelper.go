package helpers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/AramLab/golang-jwt-project/database"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Структура для полезной нагрузки.
type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.StandardClaims
}

// Секретный ключ, для подписи токена, для проверки его подлинности.
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	// Шаг 1. Создаем полезную нагрузку claims для токена.
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	// Шаг 1. Создаем полезную нагрузку refreshClaims для токена.
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	// Шаг 2. Содаем jwt и указываем полезную нагрузку.
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Шаг 3. Подписываем jwt и получаем уже подписаный токен.
	signedToken, err = jwtToken.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	// Шаг 2. Содаем refreshJwt и указываем полезную нагрузку.
	refreshJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	// Шаг 3. Подписываем refreshJwt и получаем уже подписаный токен.
	signedRefreshToken, err = refreshJwtToken.SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return signedToken, signedRefreshToken, err
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateObj primitive.D // type D []E | D is an ordered representation of a BSON document.
	// type E struct {
	// Key   string
	// Value interface{}
	// }

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: Updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)

	if err != nil {
		log.Panic(err)
		return
	}
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg
}
