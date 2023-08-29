package models

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/Shashwat5522/golang_gloabalsearch_mongo/initializers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

type Object struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BoardID     Board              `bson:"boardID,omitempty"`
	Visible     string             `bson:"visible,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
	Description string             `bson:"description,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Type        string             `bson:"type,omitempty"`
}

func (obj *Object) CreateObjects() error {

	objSlice := []Object{}

	for i := 0; i < 500; i++ {
		object := GenerateObj()
		objSlice = append(objSlice, object)
	}
	slice := convertToInterfaceSlice(objSlice)
	collection := initializers.Mgr.Connection.Database("global_search").Collection("search_objects")
	_, err := collection.InsertMany(context.TODO(), slice)
	return err

}

func (obj *Object) SearchObject(searchword string) ([]Object, error) {
	collection := initializers.Mgr.Connection.Database("global_search").Collection("search_objects")
	query := bson.M{"$or": []interface{}{
		bson.M{"visible": bson.M{"$regex": searchword,"$options":"i"}},
		bson.M{"tags": bson.M{"$regex": searchword,"$options":"i"}},
		bson.M{"description": bson.M{"$regex": searchword,"$options":"i"}},
		bson.M{"title": bson.M{"$regex": searchword,"$options":"i"}},
		bson.M{"type": bson.M{"$regex": searchword,"$options":"i"}},
	}}
	cursor, err := collection.Find(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	var objects []Object

	if scanErr := cursor.All(context.TODO(), &objects); scanErr != nil {
		log.Fatal(scanErr)
	}
	return objects, nil
}

func GenerateObj() Object {

	objID := primitive.NewObjectID()
	boardID := primitive.NewObjectID()
	date := GenerateRandomDate()
	obj := Object{
		ID: objID,
		BoardID: Board{
			ID: boardID,
		},
		Visible:     AssignRandomVisibility(),
		Tags:        AssignRandomTags(),
		Description: AssignRandomDescription(boardID.Hex(), date),
		Title:       AssignRandomTitle(date),
		Type:        AssignRandomType(),
	}
	return obj
}

func AssignRandomType() string {
	types := []string{"BOARD", "FILE", "NOTE"}

	rand.Seed(time.Now().Unix())
	return types[rand.Intn(len(types))]
}

func AssignRandomVisibility() string {
	Visibilities := []string{"PUBLIC", "PRIVATE", "CONTACTS"}

	rand.Seed(time.Now().Unix())
	return Visibilities[rand.Intn(len(Visibilities))]
}

func AssignRandomDescription(boardID, date string) string {
	// startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	// endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	// duration := endDate.Sub(startDate)

	// randomDuration := time.Duration(rand.Int63n(int64(duration)))
	// randomDate := startDate.Add(randomDuration)

	// dateStr := randomDate.Format("Jan 02,2006")
	// dateStr = strings.ReplaceAll(dateStr, "2023", "")

	ans := date + ":" + boardID
	return ans

}

func GenerateRandomDate() string {
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	duration := endDate.Sub(startDate)

	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	randomDate := startDate.Add(randomDuration)

	dateStr := randomDate.Format("Jan 02,2006")
	dateStr = strings.ReplaceAll(dateStr, "2023", "")
	return dateStr
}

func AssignRandomTitle(date string) string {
	return "Board:" + date
}

func AssignRandomTags() []string {
	ans := []string{}
	tags := []string{"Ocean", "Water", "IsLand", "Vacation", "Enjoyment", "Fun", "Chill", "Fire", "Cool", "Nice"}

	for i := 0; i < rand.Intn(10); i++ {
		ans = append(ans, tags[rand.Intn(10)])
	}
	return ans
}

func convertToInterfaceSlice(objSlice []Object) []interface{} {

	ans := []interface{}{}

	for _, val := range objSlice {
		ans = append(ans, val)
	}
	return ans
}
