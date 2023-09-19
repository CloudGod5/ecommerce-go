package controllers

import ()

var (
	ErrCantFindProduct    = error.New("can't find the product")
	ErrCantDecodeProduct  = error.New("can't find the product.")
	ErrUserIdIsNotValid   = error.New("this user is not valid")
	ErrCantUpdateUser     = error.New("can't add this product to the cart")
	ErrCantRemoveItemCart = error.New("can't remove this item from the cart")
	ErrCantGetItem        = error.New("unable to get the item from the cart")
	ErrCantBuyCartItem    = error.New("can't update the purchase")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstanBuyer() {

}
