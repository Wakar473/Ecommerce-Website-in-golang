package database

import (
	"context"
	"errors"
	"log"
	"time"

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

func RemoveCartItem(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	id , err := primitve.ObjectIDFromHex(userID)
	if err!= nil{
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	filter := bson.D(primitive.E{Key:"_id", value:id})
	update := bson.M{"$pull":bson.M{"usercart": bson.M{"_id":productID}}}
	_, err =UpdateMany(ctx, filter, update)
	if err != nil {
		return ErrCantRemoveItemCart
	}
	return nil

}

func GetItemFromCart(ctx context.Context, userCollection *mongo.Collection, userID string) error {
	/*fetch the cart of the user 
	find the cart total 
	create an order with the items
	empty up the cart*/

	id, err := primitive.ObjectIDFromHex(userID)
	if err != NIL {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	var getcartitems models.User
	var ordercart models.Order

	ordercart.Order_ID = primitive.NewObjectID()
	ordercart.Ordered_At = time.Now()

}

func BuyItemFromCart(ctx context.Context, userCollection *mongo.Collection, userID string) error{
	/*fetch the cart of the user
	find the cart total
	create an order with the items
	empty up the cart*/

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	var getcartitems models.User
	var ordercart models.Order

	ordercart.Order_ID = primitive.NewObjectID()
	ordercart.Ordered_At = time.Now()
	ordercart.Order_cart =make([]models.ProductUser, 0)
	ordercart.Payment_Method.COD =true

	unwind := bson.D{{Key:"$unwind", Value:bson.D{primitive.E{Key:"path", Value"$usercart"}}}}
	//without the unwind you won't get access to every single user's price user card supplies 
	grouping := bson.D{{Key:"$group", Value:bson.D{primitive.E{Key:"_id", Value:"$_id"}, {Key:"total", Value: bson.D{primitive.E{Key:"$sum", Value:"$usercart.price"}}}}}}
	currentresults, err := userCollection.Aggregate(ctx, mongo.Pipeline{unwind, grouping})
	ctx.Done()
	if err != nil {
		panic(err)
	}

}
func InstantBuyer() gin.HandlerFunc{
	
}

