package main

import (
	"flag"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
)

func main() {
	topic := "sample"
	broker := "tcp://mosquitto.default.svc.cluster.local:1883"
	//broker := flag.String("broker", "tcp://127.0.0.1:1883", "The broker URI. ex: tcp://10.10.1.1:1883")
	password := flag.String("password", "", "The password (optional)")
	user := flag.String("user", "", "The User (optional)")
	id := flag.String("id", "testgoidpub", "The ClientID (optional)")
	cleansess := flag.Bool("clean", false, "Set Clean Session (default false)")
	qos := flag.Int("qos", 0, "The Quality of Service 0,1,2 (default 0)")
	store := flag.String("store", ":memory:", "The Store Directory (default use memory store)")
	flag.Parse()

	fmt.Printf("Info:\n")
	fmt.Printf("\tbroker:    %s\n", broker)
	fmt.Printf("\tclientid:  %s\n", *id)
	fmt.Printf("\tuser:      %s\n", *user)
	fmt.Printf("\tpassword:  %s\n", *password)
	fmt.Printf("\tqos:       %d\n", *qos)
	fmt.Printf("\tcleansess: %v\n", *cleansess)
	fmt.Printf("\tstore:     %s\n", *store)
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(*id)
	opts.SetUsername(*user)
	opts.SetPassword(*password)
	opts.SetCleanSession(*cleansess)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Publisher Started")

	//Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	r.GET("/mqtt", func(c *gin.Context) {
		message := c.DefaultQuery("message", "None")

		fmt.Println("---- doing publish ----")
		token := client.Publish(topic, byte(*qos), false, message)
		token.Wait()
		fmt.Println("---- doing publish done ----")

		c.JSON(200, gin.H{
			"status":  "sent",
			"message": message,
		})

	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
