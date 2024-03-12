package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// mqtt.DEBUG = log.New(os.Stdout, "", 0)
	// mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("gotrivial")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("test/test", 0, func(c mqtt.Client, m mqtt.Message) {
		fmt.Printf(">>>>>>>> %s\n", m.Payload())
	}); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		fmt.Println("ping 1")
		c.Publish("test/hello", 0, false, fmt.Sprintf("123.%d", rand.Int()%10))
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
		time.Sleep(6 * time.Second)
	}

	// if token := c.Subscribe("go-mqtt/sample", 0, nil); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// 	os.Exit(1)
	// }

	// for i := 0; i < 5; i++ {
	// 	text := fmt.Sprintf("this is msg #%d!", i)
	// 	token := c.Publish("go-mqtt/sample", 0, false, text)
	// 	token.Wait()
	// }

	// time.Sleep(6 * time.Second)

	// if token := c.Unsubscribe("go-mqtt/sample"); token.Wait() && token.Error() != nil {
	// 	fmt.Println(token.Error())
	// 	os.Exit(1)
	// }

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}
