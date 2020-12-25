package main

import (
	"emqttTT/core"
	"emqttTT/util"
	"log"
	"math/rand"
	"os"
	"time"
)

const timeout = 20 * time.Second


func createClient() func(string){
	return func(url string) {
		log.Println("task url",url)
		q := util.CreateMqttClient(url)
		q.ConnectMqtt()
		q.PublishMqtt("date",rand.Intn(2),time.Now().String())
		time.Sleep(time.Duration(1)*time.Second)
		//q.ConnectMqtt()
	}
}

func TestDemo() {
	r := core.New(timeout)
	urls := []string{
		"tcp://127.0.0.1:1883",
	}
	for {
		r.Add(createClient(), createClient())
		if err := r.Start(urls); err != nil {
			switch err {
			case core.ErrorInterrupt:
				log.Println("Interrupt Error")
				os.Exit(1)
			case core.ErrorTimeOut:
				log.Println("Timeout Error")
				os.Exit(2)
			}
		}
	}
}



func main()  {
	TestDemo()
}
