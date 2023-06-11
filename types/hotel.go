package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Location string             `bson:"location" json:"location"`
	Rooms    []Room             `bson:"rooms" json:"rooms"`
}
type AggregatedHotel struct {
	Hotel Hotel  `bson:"_id"`
	Rooms []Room `bson:"rooms"`
}
