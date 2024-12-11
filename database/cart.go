package database

import (
	"errors"

	"github.com/gin-gonic/gin"
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

func AddToCart()gin.HandlerFunc {

}

func RemoeItem() gin.HandlerFunc{

}

func GetItemFromCart() gin.HandlerFunc{

}

func BuyFromCart() gin.HandlerFunc{

}
func InstantBuy() gin.HandlerFunc{
	
}

