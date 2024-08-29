package main

import (
	"fmt"
	matchpb "gabaithon-09-back/pkg/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type MatchServer struct {
	matchpb.UnimplementedMatchServiceServer
	matchingQueue chan *matchpb.MatchRequest
	lock          sync.Mutex
}

func NewMatchServer() *MatchServer {
	return &MatchServer{
		matchingQueue: make(chan *matchpb.MatchRequest, 2),
	}
}

func main() {
	port := 8083
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	matchpb.RegisterMatchServiceServer(s, NewMatchServer())
	reflection.Register(s)

	go func() {
		log.Printf("start server on port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping server...")
	s.GracefulStop()
}

func (s *MatchServer) FindMatch(req *matchpb.MatchRequest, stream matchpb.MatchService_FindMatchServer) error {
	response := &matchpb.MatchResponse{Message: "Received request from Player ID: " + req.PlayerId}
	if err := stream.Send(response); err != nil {
		return err
	}

	s.lock.Lock()
	s.matchingQueue <- req
	s.lock.Unlock()

	timeout := time.After(30 * time.Second)

	for {
		select {
		case peer := <-s.matchingQueue:
			if peer.PlayerId != req.PlayerId {
				response := &matchpb.MatchResponse{Message: "Match found with " + peer.PlayerId}
				if err := stream.Send(response); err != nil {
					return err
				}

				peerResponse := &matchpb.MatchResponse{Message: "Match found with " + req.PlayerId}
				if err := stream.Send(peerResponse); err != nil {
					return err
				}
				return nil
			} else {
				s.lock.Lock()
				s.matchingQueue <- peer
				s.lock.Unlock()
			}

		case <-timeout:
			return fmt.Errorf("timeout: no match found for Player ID: %s", req.PlayerId)
		}
	}
}
