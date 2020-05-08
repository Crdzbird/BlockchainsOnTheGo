package main

import (
	"context"
	"google.golang.org/grpc"
	"grpcBlockchain/chainer"
	"grpcBlockchain/proto"
	"log"
	"net"
	"time"
)

type Server struct {
	chain *chainer.BlockChain
}

func (s Server) AddBlock(ctx context.Context, request *proto.BlockRequest) (*proto.BlockResponce, error) {
	block := s.chain.AppendBlock(request.Data)
	return &proto.BlockResponce{
		Hash: block.Hash,
	}, nil
}

func (s Server) StreamGetBlocks(request *proto.ChainRequest, server proto.BlockChain_StreamGetBlocksServer) error {
	for _, block := range s.chain.Blocks{
		result := &proto.ChainStreamResponse{
			Block: &proto.Block{
				PrvHash: block.PrvHash,
				Data: block.Data,
				Hash: block.Data,
			},
		}
		server.Send(result)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (s Server) GetChain(ctx context.Context, request *proto.ChainRequest) (*proto.ChainResponce, error) {
	panic("implement me")
}

func main(){
	lis, err := net.Listen("tcp",":8003")
	if err != nil {
		log.Fatalf("unable to listen on port: %v", err)
	}
	srv := grpc.NewServer()
	server := &Server{chainer.MakeBlockChain()}
	proto.RegisterBlockChainServer(srv,server)
	srv.Serve(lis)
}