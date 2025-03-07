package database

import (
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wakar473/Ecommerce-Website/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("cant find the product")
    ErrUserIdIsNotValid = errors.New("this user is not valid")
    ErrCantUpdateUser = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCanGetItem = errors.New("was unable to get the item from the cart")
	ErrCantBuyCarItem = errors.New("cannot update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID,userID string) error { 
	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id":productID})
	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}
	var productCart []models.ProductUser
	err = searchfromdb.All(ctx, &productcart)
	if err != nil {
		log.Print(err)
		return ErrCantDecodeProducts
	}

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{key:"_id", Value: id}}
	update := bson.D{{Key:"$push", Value: bson.D{primitive.E{key:"usercart", Value: bson.D{{key:"$each", Value:productCart}}}}}}

	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
	return nil

}

func RemoveCartItem() gin.HandlerFunc{

}

// func GetItemFromCart() gin.HandlerFunc{

// }

func BuyItemFromCart() gin.HandlerFunc{

}
func InstantBuyer() gin.HandlerFunc{
	
}

