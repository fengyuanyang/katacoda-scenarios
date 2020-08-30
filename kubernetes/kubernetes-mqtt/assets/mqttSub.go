package main

import (
	"flag"
	"fmt"
	"os"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

/*
Options:
 [-help]                      Display help
 [-a pub|sub]                 Action pub (publish) or sub (subscribe)
 [-m <message>]               Payload to send
 [-n <number>]                Number of messages to send or receive
 [-q 0|1|2]                   Quality of Service
 [-clean]                     CleanSession (true if -clean is present)
 [-id <clientid>]             CliendID
 [-user <user>]               User
 [-password <password>]       Password
 [-broker <uri>]              Broker URI
 [-topic <topic>]             Topic
 [-store <path>]              Store Directory
*/

func main() {
	topic := "sample"
	broker := "tcp://mosquitto.default.svc.cluster.local:1883"
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	id := flag.String("id", "testgoidSub", "The ClientID (optional)")
	cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	qos := flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")
	//num := flag.Int("num", 1, "The number of messages to publish or subscribe (default 1)")
	payload := flag.String("message", "", "The message text to publish (default empty)")
	//action := flag.String("action", "", "Action publish or subscribe (required)")
	store := flag.String("store", ":memory:", "The Store Directory (default use memory store)")
	flag.Parse()

	fmt.Printf("Sample Info:\n")
	//fmt.Printf("\taction:    %s\n", action)
	//fmt.Printf("\tbroker:    %s\n", *broker)
	fmt.Printf("\tclientid:  %s\n", *id)
	fmt.Printf("\tuser:      %s\n", *user)
	fmt.Printf("\tpassword:  %s\n", *password)
	fmt.Printf("\ttopic:     %s\n", topic)
	fmt.Printf("\tmessage:   %s\n", *payload)
	fmt.Printf("\tqos:       %d\n", *qos)
	fmt.Printf("\tcleansess: %v\n", *cleansess)
	//fmt.Printf("\tnum:       %d\n", *num)
	fmt.Printf("\tstore:     %s\n", *store)

	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(*id)
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetCleanSession(*cleansess)

	if *store != ":memory:" {
		opts.SetStore(MQTT.NewFileStore(*store))
	}


	receiveCount := 0
	choke := make(chan [2]string)

	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		choke <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, byte(*qos), nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		incoming := <-choke
		fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
		receiveCount++
	}

	client.Disconnect(250)
	fmt.Println("Sample Subscriber Disconnected")
}