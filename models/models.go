package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	id              primitive.ObjectID
	first_name      *string
	last_name       *string
	email           *string
	password        *string
	phone           *string
	token           *string
	refresh_token   *string
	created_at      time.Time
	updated_at      time.Time
	user_id         string
	usercart        []ProductUser
	address_details []Address
	order_status    []Order
}

type Product struct {
	product_id   primitive.ObjectID
	product_name *string
	price        *uint64
	rating       *uint8
	image        *string
}

type ProductUser struct {
	product_id   primitive.ObjectID
	product_name *string
	price        int
	rating       *uint
	image        *string
}

type Address struct {
	address_id primitive.ObjectID
	house      *string
	street     *string
	city       *string
	zipcode    *string
}

type Order struct {
	order_id       primitive.ObjectID
	order_cart     []ProductUser
	ordered_at     time.Time
	price          int
	discount       int
	payment_method Payment
}

type Payment struct {
	digital bool
	cod     bool
}
