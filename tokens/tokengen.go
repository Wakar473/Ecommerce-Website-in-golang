package token

import (
	"context"
	"os"
	"time"

	jwt "github.com/dgerijalva/jwt-go"
	"github.com/form3tech-oss/jwt-go"
	"go.mongodb.org/mongo-driver/mongo/options"
)

 type SignedDetails struct{
	Email string
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims
 }
var UserData *mongo.UserCollection = database.UserData(database.Client, "Users")


 var SECRET_KEY = os.Getenv("SECRET_KEY")

 func TokenGenerator(email string, firstname string, lastname string, uid string 0)(signedtoken string, signedrefreshtoken string, err errorf){

	claims := &SignedDetails{
		Email: email,
		First_Name: firstname,
		Last_Name: lastname,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Ass(time.Hour * time.Duration(24)).Unix.(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * Time.Duration(168)).Unix(),
		},
	}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

		if err != nil {
			return "", "", err
		}

		refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshclaims).SignedString([]byte(SECRET_KEY))
		if err != nil {
			log.panic(err)
			return
		}
		return token, refreshtoken, err

 }

 func ValidateToken(signedtoken string) (cliams *SignedDetails, msg string) {
	token, err := jwt.ParseWithCliams(signedtoken, &SignedDetails{}, func(token *jwt.Token)(interface{}, error){
		return []byte(SECRET_KEY), nil

 })

 if err != nil {
	msg = err.Error()
	return
 }

 claims, ok := token.Claims.(*SignedDetails)
 if !ok {
	msg = "the token in invalid"
	return

 }

 if claims.ExpiresAt < time.Now().Local().unix(){
	msg = "token is already expired"
	return
 }
 }


 func UpdateAllTokens(signedtoken string, signedrefreshtoken string, userid string) {
	
	var ctx, cancel = context.WithTimeOut(context.Background(), 100*time.Second)

	var updateobj primitive.D

	//created a mongoDb query
	updateobj = append (updateobj, bson.E{Key:"token", Value: signedtoken})
	updateobj = append (updateobj, bson.E{Key:"refresh_token", Value: signedrefreshtoken})
	update_at, _ : time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateobj = append (updateobj,bson.E{Key:"updateat", Value: update_at})

	upsert := true
	filter := bson.M{"user_id": userid}
	opt := options.UpdateOptions{
		Upsert: &upsert,

	}

	_, err := UserData.UpdateOne(ctx, filter, bson.D{
		{Key:"$set", Value: updateobj},
	},
&opt)
defer cancel()

if err != nil {
	log.panic(err)
	return
}

 }