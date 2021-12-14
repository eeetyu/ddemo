package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"gotest/kitex_gen/api"
	"gotest/kitex_gen/api/gotest"
	"log"
	"time"
)

func main() {
	client, err := gotest.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.AddRequest{ID: 1, UserID: 100, RoomID: 200, Price: 66, Time: time.Now().Format(time.ANSIC),
		ProductID: 1111}
	resp, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	time.Sleep(time.Second)
	//
	//req = &api.AddRequest{ID:2,UserID: 110,RoomID: 200,Price: 166,Time: time.Now().Format(time.ANSIC),
	//	ProductID: 11121}
	//resp, err = client.Add(context.Background(), req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)
	//time.Sleep(time.Second)
	//
	//req = &api.AddRequest{ID:3,UserID: 110,RoomID: 500,Price: 166,Time: time.Now().Format(time.ANSIC),
	//	ProductID: 11121}
	//resp, err = client.Add(context.Background(), req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)
	//time.Sleep(time.Second)

	s_req := &api.SelectRequest{RoomID: 200}
	s_resp, err := client.Select(context.Background(), s_req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s_resp)
	time.Sleep(time.Second)

	s_req = &api.SelectRequest{RoomID: 500}
	s_resp, err = client.Select(context.Background(), s_req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s_resp)
	time.Sleep(time.Second)

	d_req := &api.DeleteRequest{ID: 1}
	d_resp, err := client.Delete(context.Background(), d_req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(d_resp)
	time.Sleep(time.Second)

	s_req = &api.SelectRequest{RoomID: 200}
	s_resp, err = client.Select(context.Background(), s_req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(s_resp)
	time.Sleep(time.Second)
}
