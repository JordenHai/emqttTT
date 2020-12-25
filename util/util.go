package util

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//q := make(chan *Message,10)

type Message struct {
	topic string
	body string
}

type Client struct{
	mqtt.Client
	Url string
}

func CreateMqttClient(url string) *Client{
	client := &Client{
		Url: url,
	}
	client.initMqtt()
	return client
}

//初始化的mqtt client
func (client *Client) initMqtt()  {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(client.Url)
	client.Client = mqtt.NewClient(opts)
}

func (client *Client) ConnectMqtt() {
	token := client.Client.Connect()
	if token.Wait() && token.Error() != nil{
		fmt.Println("Error",token.Error())
	}
}

func (client *Client) DisconnectMqtt()  {
	client.Client.Disconnect(250)
}

var f mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	topic := message.Topic()
	//info := message.Payload()

	//var m =  Message{topic: topic,body: string(info)}
	//fmt.Println(m)
	fmt.Print("[",topic,"]:")
	fmt.Printf("%s\n",message.Payload())

}

func (client *Client) SubscribeMqtt(topic string,qos int){

	token := client.Client.Subscribe(topic,byte(qos),f)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}

func (client *Client) PublishMqtt(topic string,qos int,msg string){
	//msg := time.Now().String()
	token := client.Client.Publish(topic,byte(qos),false,msg)
	if token.Wait() && token.Error()!=nil{
		fmt.Println(token.Error())
	}
}

