package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Product string             `json:"product,omitempty" validate:"required"`
	Stock   int                `json:"stock,omitempty" validate:"required"`
}
