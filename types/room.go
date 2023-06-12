package types

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

type RoomType int

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSiteRoomType
	DeluxeRoomType
)

type Room struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Type      RoomType           `bson:"type" json:"type"`
	BasePrice float64            `bson:"base_price" json:"base_price"`
	Price     float64            `bson:"price" json:"price"`
	HotelID   primitive.ObjectID `bson:"hotelID" json:"hotelID"`
}

type CreateRoomDto struct {
	Type      RoomType           `bson:"type" json:"type"`
	BasePrice float64            `bson:"base_price" json:"base_price"`
	HotelID   primitive.ObjectID `bson:"hotelID" json:"hotelID"`
}

type UpdateRoomDto struct {
	ID        primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	Type      *RoomType          `bson:"type" json:"type"`
	BasePrice *float64           `bson:"base_price" json:"base_price"`
}

func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func (dto *CreateRoomDto) UnmarshalJSON(b []byte) error {
	type Alias CreateRoomDto
	aux := &struct {
		BasePrice string `json:"base_price"`
		*Alias
	}{
		Alias: (*Alias)(dto),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	var err error
	dto.BasePrice, err = StringToFloat64(aux.BasePrice)
	if err != nil {
		return fmt.Errorf("cannot unmarshal string into Go struct field CreateRoomDto.base_price of type float64: %w", err)
	}

	return nil
}

func (dto *UpdateRoomDto) UnmarshalJSON(b []byte) error {
	type Alias UpdateRoomDto
	aux := &struct {
		BasePrice *string `json:"base_price"`
		*Alias
	}{
		Alias: (*Alias)(dto),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if aux.BasePrice != nil {
		dto.BasePrice = new(float64)
		*dto.BasePrice, _ = StringToFloat64(*aux.BasePrice)

	}

	return nil
}
func (rt *RoomType) UnmarshalJSON(b []byte) error {
	var roomTypeStr string
	err := json.Unmarshal(b, &roomTypeStr)
	if err != nil {
		return err
	}

	switch roomTypeStr {
	case "SingleRoomType":
		*rt = SingleRoomType
	case "DoubleRoomType":
		*rt = DoubleRoomType
	case "SeaSiteRoomType":
		*rt = SeaSiteRoomType
	case "DeluxeRoomType":
		*rt = DeluxeRoomType
	default:
		return fmt.Errorf("invalid room type: %s", roomTypeStr)
	}

	return nil
}
