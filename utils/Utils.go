package utils

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (id primitive.ObjectID) String() string {
	return fmt.Sprintf("ObjectID(%q)", id.Hex())
}