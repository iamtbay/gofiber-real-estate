package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/iamtbay/real-estate-api/helpers"
	"github.com/iamtbay/real-estate-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Estates struct {
	Message          string           `json:"message"`
	TotalEstateCount int64            `json:"totalEstateCount"`
	CurrentPage      int8             `json:"currentPage"`
	TotalPage        int8             `json:"totalPage"`
	Data             []*models.Estate `json:"data"`
}

func InitEstates() *Estates {
	return &Estates{}
}

var limit int64 = 5

// Get All Estates On DB
func (s *Estates) GetAllEstates(page int) (*Estates, error) {
	//db collection and ctx operations
	collection := mongoDB.Client.Database("Go-Real-Estate").Collection("estates")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//query on db section
	filter := bson.M{}

	//Pagination section
	totalCount, err := collection.CountDocuments(ctx, filter)
	totalPage, page := helpers.PageConverter(totalCount, limit, page)

	if err != nil {
		fmt.Println(err)
	}
	if totalCount < 1 {

		return nil, errors.New("no data to show")
	}
	options := options.Find().SetSkip(int64(page)*limit - limit).SetLimit(limit)
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	//get data section
	var results []*models.Estate
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	//return
	return &Estates{
		TotalEstateCount: totalCount,
		TotalPage:        int8(totalPage),
		CurrentPage:      int8(page),
		Data:             results,
	}, nil
}

// Search Estates By Query OnDB
func (s *Estates) GetEstatesByQuery(page int, searchInput *models.EstateSearch) (*Estates, error) {
	//
	collection := mongoDB.Client.Database("Go-Real-Estate").Collection("estates")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	//filters
	filter := bson.M{}
	featureFilter := bson.M{}
	if len(searchInput.Features) > 0 {
		arr := bson.A{}
		for i := 0; i <= len(searchInput.Features)-1; i++ {
			arr = append(arr, searchInput.Features[i])
			fmt.Println(arr, i)
		}
		featureFilter["$in"] = arr
	}
	filter["features"] = featureFilter
	setFilter := func(filterField string, inputValue any) {
		if inputValue != "" {
			filter[filterField] = inputValue
		}
	}
	rangeFilter := func(filterField string, minValue, maxValue int) {
		if minValue > 0 || maxValue > 0 {
			subFilter := bson.M{}
			if minValue > 0 {
				subFilter["$gte"] = minValue
			}
			if maxValue > 0 {
				subFilter["$lte"] = maxValue
			}
			filter[filterField] = subFilter
		}
	}
	setFilter("estate_type", searchInput.EstateType)
	setFilter("estate_status", searchInput.EstateStatus)
	setFilter("location.city", searchInput.Location.City)
	setFilter("location.state", searchInput.Location.State)
	setFilter("location.country", searchInput.Location.Country)

	rangeFilter("rooms.bathroom", int(searchInput.BathroomsMin), int(searchInput.BathroomsMax))
	rangeFilter("rooms.bedroom", int(searchInput.BedroomsMin), int(searchInput.BedroomsMax))
	rangeFilter("rooms.total_rooms", int(searchInput.TotalRoomMin), int(searchInput.TotalRoomMax))
	rangeFilter("floor", int(searchInput.FloorMin), int(searchInput.FloorMax))
	rangeFilter("year_built", int(searchInput.YearBuiltMin), int(searchInput.YearBuiltMax))
	rangeFilter("square_mt", int(searchInput.SquareMtMin), int(searchInput.SquareMtMax))
	rangeFilter("price", int(searchInput.PriceMin), int(searchInput.PriceMax))

	//pagination
	totalCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	totalPage, page := helpers.PageConverter(totalCount, limit, page)
	if page < 1 {
		return nil, errors.New("sorry we couldn't find any estate based on your search terms")
	}
	//
	options := options.Find().SetSkip(int64(page)*limit - limit).SetLimit(limit)
	cursor, err := collection.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	var results []*models.Estate
	if err := cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	return &Estates{
		TotalEstateCount: totalCount,
		TotalPage:        int8(totalPage),
		CurrentPage:      int8(page),
		Data:             results,
	}, nil
}

// Get Single Estate On DB
func (s *Estates) GetSingleEstate(idString string) (*models.Estate, error) {
	//db open, ctx operations
	collection := mongoDB.Client.Database("Go-Real-Estate").Collection("estates")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	//c
	estateID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return nil, err
	}
	//db query operations
	filter := bson.M{"_id": estateID}
	var estate *models.Estate
	err = collection.FindOne(ctx, filter).Decode(&estate)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return nil, errors.New("couldn't find. check the id")
		}
	}
	return estate, nil

}

// Add New Estate To DB
func (s *Estates) AddNewEstate(estateInfo *models.Estate) error {
	collection := mongoDB.Client.Database("Go-Real-Estate").Collection("estates")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	//
	filter := bson.M{
		"owner_id":      estateInfo.OwnerID,
		"title":         estateInfo.Title,
		"estate_type":   estateInfo.EstateType,
		"estate_status": estateInfo.EstateStatus,
		"price":         estateInfo.Price,
		"description":   estateInfo.Description,
		"year_built:":   estateInfo.YearBuilt,
		"location":      estateInfo.Location,
		"floor":         estateInfo.Floor,
		"rooms":         estateInfo.Rooms,
		"features":      estateInfo.Features,
		"images":        estateInfo.Images,
		"square_mt":     estateInfo.SquareMt,
		"created_at":    estateInfo.CreatedAt,
	}
	_, err := collection.InsertOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// Update an Estate On DB
func (s *Estates) UpdateEstate(estateInfo *models.Estate, idString, ownerID string) ([]string, error) {
	//db,ctx operations
	collection := mongoDB.Client.Database("Go-Real-Estate").Collection("estates")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	//find estate id
	estateID, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return nil, errors.New("something went wrong! check id")
	}
	//filter
	filter := bson.D{
		{Key: "_id", Value: estateID},
		{Key: "owner_id", Value: ownerID},
	}
	var images *Images
	_ = collection.FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"images": 1, "_id": 0})).Decode(&images)
	var imgShouldDel []string
	for _, v := range images.Images {
		found := false
		for _, vU := range estateInfo.Images {
			if vU == v {
				found = true
				break
			}
		}
		if !found {
			imgShouldDel = append(imgShouldDel, v)
		}
	}
	//update
	update := bson.D{
		{
			Key: "$set",
			Value: bson.M{
				"title":         estateInfo.Title,
				"estate_type":   estateInfo.EstateType,
				"estate_status": estateInfo.EstateStatus,
				"price":         estateInfo.Price,
				"description":   estateInfo.Description,
				"year_built:":   estateInfo.YearBuilt,
				"location":      estateInfo.Location,
				"floor":         estateInfo.Floor,
				"rooms":         estateInfo.Rooms,
				"features":      estateInfo.Features,
				"images":        estateInfo.Images,
				"square_mt":     estateInfo.SquareMt,
				"created_at":    estateInfo.CreatedAt,
			},
		},
	}
	var updatedEstate *models.Estate
	_ = collection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedEstate)

	return imgShouldDel, nil
}

// Delete an Estate On DB
func (s *Estates) DeleteEstate(idString, ownerID string) error {
	//db, context ops
	collection := mongoDB.Client.Database("Go-Real-Estate").Collection("estates")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// db op
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return errors.New("something went wrong! check the id")
	}
	//get images path
	filterForImg := bson.M{"_id": id, "owner_id": ownerID}
	options := options.FindOne().SetProjection(bson.M{"images": 1, "_id": 0})
	var images *Images
	_ = collection.FindOne(ctx, filterForImg, options).Decode(&images)

	filter := bson.M{"_id": id, "owner_id": ownerID}
	singleRes := collection.FindOneAndDelete(ctx, filter)
	if singleRes.Err() != nil {
		return errors.New("something went wrong! check the id")
	}
	//delete image paths

	for i := 0; i <= len(images.Images)-1; i++ {
		imageSplit := strings.Split(images.Images[i], "/")
		if err := os.Remove(fmt.Sprintf("./public/uploads/%v", imageSplit[len(imageSplit)-1])); err != nil {
			return err
		}
	}

	return nil
}

type Images struct {
	Images []string
}
