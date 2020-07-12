package suites

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Message struct {
	Id   string `bson:"_id"`
	Rid  string `bson:"rid"`
	Msg  string `bson:"msg"`
	User User   `bson:"u"`
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

func getcollection(client *mongo.Client, db string, col string) (*mongo.Cursor, error) {
	collection := client.Database(db).Collection(col)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	return cur, err
}

func interpretateId(client *mongo.Client, db string, col string, id string) (*Room, error) {
	result := &Room{}
	collection := client.Database(db).Collection(col)
	filter := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Error(err)
		return &Room{}, err
	}
	//log.Info(result)
	return result, nil
}
