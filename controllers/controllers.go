package controllers

import (
	"context"
	"go/token"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb/cmd/influx_tools/generate"
	"github.com/wakar473/Ecommerce-Website/database"
	"github.com/wakar473/Ecommerce-Website/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Product")
var Validate = validator.New()


func HashPassword (password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	/*bytes,err :=
	bytes: Yeh variable hashed password ko store karta hai.
    err: Agar hashing me koi problem hoti hai to yeh error handle karega.*/
	if err != nil {
		log.Panic(err)

	}
	return string(bytes)
		
}


func VerifyPassword (userPassword string, givenPassword string) (bool, string){
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	valid := true
	msg := ""

	if err != nil {
		msg = "login or Password is incorrect"
		valid = false
	}
	return valid, msg

}

func Signup () gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel ()
		var user models.User
		if err := c.BindJSON(&user); err!= nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		if err != nil{
			log.Panic(err)
			c.json(http.StatusBadRequest, gin.H{"error":"user already exists"})
		}
		count,err = UserCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})

		defer cancel()
		if err!= nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error": err})
			return
		}

		if count>0{
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone no is already in use"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password =&password
		
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"the user did not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Successfully signed in!")

	}

}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user model.User
		if err := c.BindJSON(&user); err !=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err !=nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}

		PasswordIsValid, msg :=VerifyPassword(*user.Password, *founder.Password)

		defer cancel()

		if !PasswordIsValid{
			c.JSON{http.StatusInternalServerError, gin.H{"error": msg}}
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ :=generate.TokenGenerator(*founduser.Email, *founduser.First_Name. *founderuser.Last_Name *founduser.User_ID)
		defer cancel()

		generate.UpdateAllTokens(token, refreshToken, founderuser.User_ID)

		c.JSON(http.StatusFound, founduser)
	}
}

func ProductViewerAdmin() gin.HandlerFunc{

}

func SearchProduct() gin.HandlerFunc{
	return func(c *gin.Context){

		var productlist []models.Product
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)	
		defer cancel()

		ProductCollection.Find(ctx, bson.D{{}})
		if err! = nil {
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong, please try after some time")
			return

		}

		err = cursor.All(ctx, &productlist)

		if err!= nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		defer cursor.Close()
		if err := cursor.err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}
		defer cancel()
		c.IndentedJSON(200, productlist)
	}

}

func SearchProductByQuery() gin.HandlerFunc{
	return func( c *gin.Context){
		var searchProducts []models.Product
		queryParam := c.Query("name")

		//you want to check if it's empty

		if queryParam == ""{
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(HTTP.StatusNotFound, gin.H{"Error":"Invalid search index"})
			c.Abort()
			return

		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		searchquerydb, err := ProductCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex":queryParam}})

		if err != nil {
			c.IndentedJSON(404, "something went wrong while fetching the data")
			return
		}

		err = searchquerydb.All(ctx, &searchProducts)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		defer searchquerydb.Close(ctx)

		if err := searchquerydb.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid request")
			return
		}

		defer cancel()
		c.IndentedJSON(200,searchProducts)
	}


}