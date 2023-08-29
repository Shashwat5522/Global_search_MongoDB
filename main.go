package main

import (
	"fmt"

	"github.com/Shashwat5522/golang_gloabalsearch_mongo/initializers"
	"github.com/Shashwat5522/golang_gloabalsearch_mongo/router"
)

func init() {
	initializers.GetInstance()
}
func main() {
	fmt.Println("hello world")
	// InsertMany()
	router.Run()
}

// func InsertMany() {
// 	objID:=primitive.NewObjectID()
// 	boardID:=primitive.NewObjectID()
// 	date:=GenerateRandomDate()
// 	obj := models.Object{
// 		ID: objID,
// 		BoardID: models.Board{
// 			ID: boardID,
// 		},
// 		Visible: AssignRandomVisibility(),
// 		Tags: AssignRandomTags(),
// 		Description:AssignRandomDescription(boardID.Hex(),date),
// 		Title:       AssignRandomTitle(date),
// 		Type:        AssignRandomType(),
// 	}

// 	collection := initializers.Mgr.Connection.Database("global_search").Collection("search_objects")
// 	result, err := collection.InsertOne(context.TODO(), obj)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(result)
// 	log.Println(obj)

// }

// func AssignRandomType() string {
// 	types := []string{"BOARD", "FILE", "NOTE"}

// 	rand.Seed(time.Now().Unix())
// 	return types[rand.Intn(len(types))]
// }

// func AssignRandomVisibility() string {
// 	Visibilities := []string{"PUBLIC", "PRIVATE", "CONTACTS"}

// 	rand.Seed(time.Now().Unix())
// 	return Visibilities[rand.Intn(len(Visibilities))]
// }

// func AssignRandomDescription(boardID,date string) string {
// 	// startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
// 	// endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

// 	// duration := endDate.Sub(startDate)

// 	// randomDuration := time.Duration(rand.Int63n(int64(duration)))
// 	// randomDate := startDate.Add(randomDuration)

// 	// dateStr := randomDate.Format("Jan 02,2006")
// 	// dateStr = strings.ReplaceAll(dateStr, "2023", "")

// 	ans := date + ":" + boardID
// 	return ans

// }

// func GenerateRandomDate()string{
// 	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
// 	endDate := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

// 	duration := endDate.Sub(startDate)

// 	randomDuration := time.Duration(rand.Int63n(int64(duration)))
// 	randomDate := startDate.Add(randomDuration)

// 	dateStr := randomDate.Format("Jan 02,2006")
// 	dateStr = strings.ReplaceAll(dateStr, "2023", "")
// 	return dateStr
// }

// func AssignRandomTitle(date string)string{
// 	return "Board:"+date
// }

// func AssignRandomTags()[]string{
// 	ans:=[]string{}
// 	tags:=[]string{"Ocean","Water","IsLand","Vacation","Enjoyment","Fun","Chill","Fire","Cool","Nice"}

// 	for i:=0;i<rand.Intn(10);i++{
// 		ans=append(ans,tags[rand.Intn(10)])
// 	}
// 	return ans
// }
