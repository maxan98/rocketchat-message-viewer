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
func GetAllRooms() []Room {
	client, ctx, err, cancel := getConnection("mongodb://localhost:27017")
	defer cancel()
	cur, err := getCollection(client, "rocketchat", "rocketchat_room",bson.D{})
	if err != nil {
		log.Error(err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Usernames"})
	var resultBson []Room
	for cur.Next(ctx) {
		var room Room
		err := cur.Decode(&room)
		//	if mes.Rid != roomid {
		//		continue
		//	}
		resultBson = append(resultBson,room)
		if err != nil {
			log.Fatal(err)
		}
		log.Debugf("%s  ->  %s  ->  %s",room.Id, room.Name,room.Usernames)

		table.Append([]string{room.Id, room.Name,strings.Join(room.Usernames,",")})


	}
	if log.GetLevel() == log.DebugLevel{
	table.Render()
	}
return resultBson
}

func GetAllMessagesByFilter(filter bson.D, baseurl string) []Message {
    client, ctx, err, cancel := getConnection("mongodb://localhost:27017")
    defer cancel()
	cur, err := getCollection(client, "rocketchat", "rocketchat_message",filter)
	if err != nil {
		log.Error(err)
	}

	var resultStruct []Message
	for cur.Next(ctx) {
		var mes Message
		err := cur.Decode(&mes)
		//	if mes.Rid != roomid {
		//		continue
		//	}

		if err != nil {
			log.Fatal(err)
		}

		log.Debugf("Room: %s \n", mes.Rid)
		log.Debugf("Message id: %s \n", mes.Id)
		res, err := interpretateId(client, "rocketchat", "rocketchat_room", mes.Rid)
		if err != nil {
			log.Warn("Cannot interpretate room id")
		}
		log.Debugf("Room Name: %s", res.Name)
		mes.Rid = fmt.Sprintf("%s%s",res.Name , res.Usernames)
		if res.Usernames != nil {
			log.Debugf("[ %s ]", res.Usernames)
		}
		if mes.Type == "jitsi_call_started" {
			log.Debugf("This is a Jitsi call. Skipping\n_________\n")
			continue
		}
		log.Debugf("\n")
		log.Debugf("User: %s (%s) \n", mes.User.Name, mes.User.Username)
		log.Debugf("Message: %s \n", mes.Msg)
		log.Debugf(mes.Time.Format("2 Jan 15:04:05"))
		if mes.URLS != nil {
			for i := range mes.URLS{
			log.Debugf("URL: %s\n", mes.URLS[i].Url)
			}
		}

			log.Debug("here")
			curn, err := getCollection(client, "rocketchat", "rocketchat_message",bson.D{{"tmid",mes.Id}})
			log.Debugf(mes.Id)
			if err != nil {
				log.Error(err)
			}

			for curn.Next(ctx) {
				var mesn Message
				err := curn.Decode(&mesn)
				log.Debugf(mesn.Id)
				if err != nil {
					log.Fatal(err)
				}
				mes.Thread = append(mes.Thread,mesn)
			}


		if mes.Attachments != nil{
			for i := range mes.Attachments{
				if mes.Attachments[i].Title !=""{
					log.Debugf("Attatchment title: %s \n", mes.Attachments[i].Title)
				}
				if mes.Attachments[i].ImageUrl !=""{
					log.Debugf("Attatchment image URL: %s%s \n", baseurl,mes.Attachments[i].ImageUrl)
					mes.Attachments[i].ImageUrl = fmt.Sprintf("https://%s%s",baseurl,mes.Attachments[i].ImageUrl)

				}
				if mes.Attachments[i].Description !=""{
					log.Debugf("Attatchment Description: %s \n", mes.Attachments[i].Description)
				}
			}
		}
		log.Debugf("\n\n_____________\n")
		resultStruct = append(resultStruct,mes)
	}
	//TODO: REMOVE MESSAGES WHICH ARE THREAD ONES

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
return resultStruct
}
