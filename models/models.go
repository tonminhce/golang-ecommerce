package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	firstName    *string            `json:"firstName" 	validate:"required, min=2, max=30"`
	lastName     *string            `json:"lastName" 	validate:"required, min=2, max=30"`
	password     *string            `json:"password" 	validate:"required, min=6, max=30"`
	email        *string            `json:"email" 		validate:"required, email"`
	phone        *string            `json:"phone"		validate:"required"`
	token        *string            `json:"token"`
	refreshToken *string            `json:"refreshToken"`
	createdAt    time.Time          `json:"createdAt"`
	updatedAt    time.Time          `json:"updatedAt"`
	userID       string             `json:"userID"`
	userCart     []UserCart         `json:"userCart" bson:"userCart"`
	address      []Address          `json:"address" bson:"address"`
	orderStatus  []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	productID   primitive.ObjectID `bson:"_id"`
	productName *string            `json:"productName"`
	price       *uint64            `json:"price"`
	ratings     *uint8             `json:"ratings"`
	image       *string            `json:"image"`
	description *string            `json:"description"`
}

type UserCart struct {
	productID   primitive.ObjectID `bson:"_id"`
	productName *string            `json:"productName" bson:"productName"`
	price       *uint64            `json:"price" bson:"price"`
	ratings     *uint8             `json:"ratings" bson:"ratings"`
	image       *string            `json:"image" bson:"image"`
}

type Address struct {
	addressID primitive.ObjectID `bson:"_id"`
	street    *string            `json:"street" bson:"street"`
	city      *string            `json:"city" bson:"city"`
}

type Order struct {
	orderID       primitive.ObjectID `bson:"_id"`
	orderCart     []UserCart         `json:"orderCart" bson:"orderCart"`
	orderedAt     time.Time          `json:"orderedAt" bson:"orderedAt"`
	price         int                `json:"price" bson:"price"`
	discount      *int               `json:"discount" bson:"discount"`
	paymentMethod Payment            `json:"paymentMethod" bson:"paymentMethod"`
}

type Payment struct {
	Bank bool
	Cash bool
}
