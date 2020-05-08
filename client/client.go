package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpcBlockchain/proto"
	"io"
	"log"
	"time"
)

var client proto.BlockChainClient

func main() {
	start := flag.Bool("start", false, "Start Blockchain")
	stream := flag.Bool("stream", false, "Start Streaming")
	flag.Parse()
	conn, err := grpc.Dial(":8003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}
	client = proto.NewBlockChainClient(conn)
	if *start {
		fmt.Println("start")
		startBlockchain()
	}
	if *stream {
		startStream()
	}
}

func startStream(){
	log.Println("blocks:")
	req := &proto.ChainRequest{}
	resStream, err := client.StreamGetBlocks(context.Background(),req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		blockStream, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		block := blockStream.Block
		log.Printf("prev hash: %s, data %s, , hash: %s \n", block.PrvHash, block.Hash, block.Data)
	}
}

func startBlockchain(){
	fmt.Println("started")
	for {
		block, addErr := client.AddBlock(context.Background(), &proto.BlockRequest{
			Data: time.Now().String(),
		})
		if addErr != nil {
			log.Fatalf("unable to add block : %v", addErr)
		}
		log.Printf("new block hash -> %s \n", block.Hash)
		time.Sleep(1 * time.Second)
	}
}