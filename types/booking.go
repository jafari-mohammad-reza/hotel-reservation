package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userID" json:"userID"`
	RoomID    primitive.ObjectID `bson:"roomID" json:"roomID"`
	StartDate primitive.DateTime `bson:"startDate" json:"startDate"`
	EndDate   primitive.DateTime `bson:"endDate" json:"endDate"`
}
