package main

import (
	"context"
	"github.com/reation/micro_service_goods_service/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	GoodsListAddress   = "192.168.1.104:8010"
	GoodsDetailAddress = "192.168.1.104:8011"
)

func main() {
	goodsList()
}

func goodsDetail() {
	conn, err := grpc.Dial(GoodsDetailAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewGoodsDetailClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GoodsDetail(ctx, &protoc.GoodsDetailRequest{GoodsId: 10})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Println(r.GetGoodDetail())

}

func goodsList() {

	conn, err := grpc.Dial(GoodsListAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewGoodsListClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GoodsList(ctx, &protoc.GoodsListRequest{Type: 0, Id: 0, Limit: 20})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	for _, v := range r.GetGoodList() {
		log.Println(v)
	}

}
