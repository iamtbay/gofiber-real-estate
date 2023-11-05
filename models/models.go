package models

import "time"

type User struct {
	ID       string `json:"_id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name"`
	Surname  string `json:"surname" bson:"surname"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Admin    bool   `json:"admin" bson:"admin"`
}

type Location struct {
	City        string    `json:"city" bson:"city"`
	State       string    `json:"state" bson:"state"`
	Country     string    `json:"country" bson:"country"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}
type Rooms struct {
	Bedroom    int8 `json:"bedroom" bson:"bedroom"`
	Bathroom   int8 `json:"bathroom" bson:"bathroom"`
	TotalRooms int8 `json:"total_rooms" bson:"total_rooms"`
}

type Estate struct {
	ID           string    `json:"_id" bson:"_id,omitempty"`
	OwnerID      string    `json:"owner_id" bson:"owner_id,omitempty"`
	Title        string    `json:"title" bson:"title,omitempty"`
	EstateType   string    `json:"estate_type" bson:"estate_type,omitempty"`
	EstateStatus string    `json:"estate_status" bson:"estate_status,omitempty"`
	Price        int32     `json:"price" bson:"price,omitempty"`
	Description  string    `json:"description" bson:"description,omitempty"`
	YearBuilt    int16     `json:"year_built" bson:"year_built,omitempty"`
	Location     Location  `json:"location" bson:"location,omitempty"`
	Floor        uint8     `json:"floor" bson:"floor,omitempty"`
	Rooms        Rooms     `json:"rooms" bson:"rooms,omitempty"`
	Features     []string  `json:"features" bson:"features,omitempty"`
	Images       []string  `json:"images" bson:"images,omitempty"`
	SquareMt     int16     `json:"square_mt" bson:"square_mt,omitempty"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at,omitempty"`
}

type EstateSearch struct {
	EstateType   string   `json:"estate_type" bson:"estate_type,omitempty"`
	EstateStatus string   `json:"estate_status" bson:"estate_status,omitempty"`
	Location     Location `json:"location" bson:"location,omitempty"`
	Features     []string `json:"features" bson:"features,omitempty"`
	//floor
	FloorMin uint8 `json:"floor_min" bson:"floor_min"`
	FloorMax uint8 `json:"floor_max" bson:"floor_max"`
	//rooms
	BedroomsMin  uint8 `json:"bedroom_min" bson:"bedroom_min"`
	BedroomsMax  uint8 `json:"bedroom_max" bson:"bedroom_max"`
	BathroomsMin uint8 `json:"bathroom_min" bson:"bathroom_min"`
	BathroomsMax uint8 `json:"bathroom_max" bson:"bathroom_max"`
	TotalRoomMin uint8 `json:"total_room_min" bson:"total_room_min"`
	TotalRoomMax uint8 `json:"total_room_max" bson:"total_room_max"`
	//year
	YearBuiltMin uint16 `json:"year_built_min" bson:"year_built_min"`
	YearBuiltMax uint16 `json:"year_built_max" bson:"year_built_max"`
	//price
	PriceMax int32 `json:"price_max" bson:"price_max,omitempty"`
	PriceMin int32 `json:"price_min" bson:"price_min,omitempty"`
	//sq mt
	SquareMtMin uint32 `json:"square_mt_min" bson:"square_mt_min,omitempty"`
	SquareMtMax uint32 `json:"square_mt_max" bson:"square_mt_max,omitempty"`
	//CreatedAt   time.Time `json:"created_at" bson:"created_at,omitempty"`
}
