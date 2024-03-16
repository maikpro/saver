package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Image struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Fullpath string             `bson:"path"`
}
