package suites

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func GetAllMessagesFromRoom(roomid string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("Error pjng")
	}
	fmt.Printf("Pung succeeded")
	cur, err := getcollection(client, "rocketchat", "rocketchat_message")
	if err != nil {
		log.Error(err)
	}
	for cur.Next(ctx) {
		var mes Message
		err := cur.Decode(&mes)
		if mes.Rid != roomid {
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....

		fmt.Printf("Room: %s \n", mes.Rid)
		res, err := interpretateId(client, "rocketchat", "rocketchat_room", mes.Rid)
		if err != nil {
			log.Warn("Cannot interpretate room id")
		}
		fmt.Printf("Room Name: %s", res.Name)
		if res.Usernames != nil {
			fmt.Printf("[ %s ]", res.Usernames)
		}
		fmt.Printf("\n")
		fmt.Printf("User: %s (%s) \n", mes.User.Name, mes.User.Username)
		fmt.Printf("Message: %s \n", mes.Msg)
		fmt.Print("\n\n_____________")
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
