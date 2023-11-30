package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       int                `bson:"price"`
	Stock       int                `bson:"stock"`
	Sizes       []Size             `bson:"sizes"`
}

type Size struct {
	Name  string `bson:"name"`
	Stock int    `bson:"stock"`
}
