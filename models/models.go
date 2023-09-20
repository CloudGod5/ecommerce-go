package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	id              primitive.ObjectID `json:"_id" bson:"_id"`
	first_name      *string            `json:"first_name" 	validate:"required,min=2,max=30"`
	last_name       *string            `json:"last_name" 	validate:"required,min=2,max=30"`
	email           *string            `json:"email" 		validate:"required,email"`
	password        *string            `json:"password" 	validate:"required,min=6`
	phone           *string            `json:"phone"		validate:"required"`
	token           *string            `json:"token"`
	refresh_token   *string            `json:"refresh_token"`
	created_at      time.Time          `json:"created_at"`
	updated_at      time.Time          `json:"updated_at"`
	user_id         string             `json:"user_id"`
	usercart        []ProductUser      `json:"usercart" bson:"usercart"`
	address_details []Address          `json:"address" bson:"address"`
	order_status    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	product_id   primitive.ObjectID `json:"_id" bson:"_id"`
	product_name *string            `json:"product_name"`
	price        *uint64            `json:"price"`
	rating       *uint8             `json:"rating"`
	image        *string            `json:"image"`
}

type ProductUser struct {
	product_id   primitive.ObjectID `json:"_id" bson:"_id"`
	product_name *string            `json:"product_name" bson:"product_name"`
	price        int                `json:"price"`
	rating       *uint              `json:"rating"`
	image        *string            `json:"image"`
}

type Address struct {
	address_id primitive.ObjectID `json:"_id" bson:"_id"`
	house      *string            `json:"house"`
	street     *string            `json:"street"`
	city       *string            `json:"city"`
	pincode    *string            `json:"pincode"`
}

type Order struct {
	order_id       primitive.ObjectID `json:"_id" bson:"_id"`
	order_cart     []ProductUser      `json:"order_cart"`
	ordered_at     time.Time          `json:"ordered_at"`
	price          int                `json:"price"`
	discount       int                `json:"discount"`
	payment_method Payment            `json:"payment_method"`
}

type Payment struct {
	digital bool `json:"digital"`
	cod     bool `json:"cod"`
}
