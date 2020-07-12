package suites

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"strings"
	"time"
)
func getConnection(url string) (*mongo.Client, context.Context, error, context.CancelFunc){
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil{
		return nil, nil, err, cancel
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, nil, err, cancel
	}
	fmt.Printf("Pung succeeded")
	return client, ctx, nil, cancel
}
func GetAllRooms() {
	client, ctx, err, cancel := getConnection("mongodb://localhost:27017")
	defer cancel()
	cur, err := getCollection(client, "rocketchat", "rocketchat_room",bson.D{})
	if err != nil {
		log.Error(err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Usernames"})
	for cur.Next(ctx) {
		var room Room
		err := cur.Decode(&room)
		//	if mes.Rid != roomid {
		//		continue
		//	}
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("%s  ->  %s  ->  %s",room.Id, room.Name,room.Usernames)

		table.Append([]string{room.Id, room.Name,strings.Join(room.Usernames,",")})


	}
	if log.GetLevel() == log.DebugLevel{
	table.Render()
	}

}
func GetAllMessagesByFilter(filter bson.D, baseurl string) {
    client, ctx, err, cancel := getConnection("mongodb://localhost:27017")
    defer cancel()
	cur, err := getCollection(client, "rocketchat", "rocketchat_message",filter)
	if err != nil {
		log.Error(err)
	}
	for cur.Next(ctx) {
		var mes Message
		err := cur.Decode(&mes)
		//	if mes.Rid != roomid {
		//		continue
		//	}
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....

		log.Debugf("Room: %s \n", mes.Rid)
		log.Debugf("Message id: %s \n", mes.Id)
		res, err := interpretateId(client, "rocketchat", "rocketchat_room", mes.Rid)
		if err != nil {
			log.Warn("Cannot interpretate room id")
		}
		log.Infof("Room Name: %s", res.Name)
		if res.Usernames != nil {
			log.Infof("[ %s ]", res.Usernames)
		}
		if mes.Type == "jitsi_call_started" {
			log.Debugf("This is a Jitsi call. Skipping\n_________\n")
			continue
		}
		log.Infof("\n")
		log.Infof("User: %s (%s) \n", mes.User.Name, mes.User.Username)
		log.Infof("Message: %s \n", mes.Msg)
		log.Infof(mes.Time.Format("2 Jan 15:04:05"))
		if mes.URLS != nil {
			for i := range mes.Attachments{
			log.Infof("URL: %s\n", mes.URLS[i].Url)
			}
		}
		if mes.Attachments != nil{
			for i := range mes.Attachments{
				if mes.Attachments[i].Title !=""{
					log.Infof("Attatchment title: %s \n", mes.Attachments[i].Title)
				}
				if mes.Attachments[i].ImageUrl !=""{
					log.Infof("Attatchment image URL: %s%s \n", baseurl,mes.Attachments[i].ImageUrl)
				}
				if mes.Attachments[i].Description !=""{
					log.Infof("Attatchment Description: %s \n", mes.Attachments[i].Description)
				}
			}
		}
		log.Debugf("\n\n_____________\n")
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

}
