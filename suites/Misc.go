package suites

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	BASEURL string

)
type Message struct {
	Id   string    `bson:"_id"`
	Rid  string    `bson:"rid"`
	Target string  `bson:"tmid"`
	Msg  string    `bson:"msg"`
	User User      `bson:"u"`
	URLS  []Url    `bson:"urls"`
	Time time.Time `bson:"ts"`
	Type string    `bson:"t"`
	Thread []Message
	Attachments []Attachment `bson:"attachments"`

}
type Url struct {
	Url string `bson:"url"`
}
type User struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
	Name     string `bson:"name"`
}

type Room struct {
	Id        string `bson:"_id"`
	Name      string `bson:"name"`
	Usernames []string
}
type Attachment struct {
	Title       string `bson:"title"`
	ImageUrl   string  `bson:"image_url"`
	Description string `bson:"description"`
}

//get whole collection or filtered. Pass empty bson.D for whole collection.
func getCollection(client *mongo.Client, db string, col string, filter bson.D) (*mongo.Cursor, error) {
	collection := client.Database(db).Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	options := options.Find()
	options.SetSort(bson.D{{"ts", 1}})
	cur, err := collection.Find(ctx, filter,options)
	if err != nil {
		log.Fatal(err)
	}
	return cur, err
}

func interpretateId(client *mongo.Client, db string, col string, id string) (*Room, error) {
	result := &Room{}
	collection := client.Database(db).Collection(col)
	filter := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Error(err)
		return &Room{}, err
	}
	//log.Debug(result)
	return result, nil
}


