package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wakar473/Ecommerce-Website/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc {

	return func(c *gin.Context){
		user_id := c.Query("id")
		if user_id == ""{
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error":"Invalid code"})
			c.Abort()
			return
		}
		address,err := ObjectIDFromHex(user_id)
		if err != nil {
			c.IntentedJSON(500, "Internal Server Error")
		}

		var addresses models.Address

		addresses.Address_id = primitive.NewObjectID()

		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		match_filter := bson.D{{key:"$match", Value: bson.D{primitive.E{key:"_id", Value: address}}}}
		unwind := bson.D{{key: "$unwind", Value:bson.D{primitive.E{KEy:"path", Value:"$address"}}}}
		group := bson.D{{Key: "$group", Value:bson.D{primitive.E{KEy:"_id", Value:"$address_id"},{Key:"count",Value: bson.D{primitive.E{Key:"$sum", Value: 1}}}}}}
		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{match_filter,unwind,group})
		if err != nil {
			c.IndentedJSON(500, "Internal server error")
		}

		var addressinfo []bson.M
		if err = pointcursor.All(ctx, &addressinfo); err !=nil {
			panic(err)
		}

		var size int32
		for _, address_no := range addressinfo{
			count := address_no["count"]
			size = count.(int32)
		}
		if size < 2 {
			filter := bson.D{primitive.E{Key:"_id", Value: address}}
			update := bson.D{{key:"$push", Value: bson.D{primitive.E{Key:"address", Value: addresses}}}}
			_, err := UserCollection.UpdateOne(crx, filter, update)
			if err != nil {
				fmt.Println(err)
			}
		}else {
			c.IndentedJSON(400, "Not Allowed")
		}
		defer cancel()
		ctx.Done()
	}

}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context){
		user_id := c.Query("id")
		if user_id == ""{
			c.Header("Content_Type", "application/json")
			c.Json(http.StatusNotFound, gin.H{"Error":"Invalid"})
			c.Abort()
			return
		}
		usert_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IntentedJSON(500, "Internal Server Error")
		}

        var editaddress models.Address
		if err := c.BindJSON(&editaddress); err != nil{
		   c.IndentedJSON(http.StatusBadRequest, err.Error())
		}
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key:"_id", Value:usert_id}}
		update := bson.D{{Key:"$set",Value:bson.D{primitive.E{Key:"address.0.house_name",Value: editaddress.House},{Key:"address.0.street_name", Value: editaddress.Street}, {Key:"address.0.city_name", Value: editaddress.City},{Key:"address.0.pin_code", Value:editaddress.Pincode}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "Something went wrong")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "successfully updated the home address")

	}

}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
       user_id := c.Query("id")
	   if user_id == ""{
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusNotFound, gin.H{"Error":"Invalid"})
		c.Abort()
		return
	   }
	   usert_id, err := primitive.ObjectIDFromHex(user_id)
	   if err != nil {
		c.IndentedJSON(500, "Internal Server Error")
	   }
	   var editaddress models.Address
	   if err := c.BindJSON(&editaddress); err !=nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	   }
	   var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	   defer cancel()
	   filter := bson.D{primitive.E{key:"_id", Value: user_id}}
	   update := bson.D{{key:"$set", Value:bson.D{primitive.E{Key:"address.1.house_name", Value: editaddress.House},{Key:"address.1.street_name", Value: editaddress.Street}, {Key:"address.1.city_name", Value: editaddress.City}, {Key:"address.1.pin_code", Value:editaddress.Pincode}}}}}
	   _, err == UserCollection.UpdateOne(ctx, filter,update)
	   if err != nil {
		c.IndentedJSON(500, "something went wrong")
	   }
	}



func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Erroe": "Invalid Search Index"})
			c.Abort()
			return

		}
		addresses := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IntentedJSON(500, "Internal Server Error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{Primitive.E{Key: "_id", Value: usert_id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", value: addresses}}}}
		if err != nil {
			c.IndentedJSON(404, "Wrong command")
			return
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(200, "Successfully Deleted")
	}

}
