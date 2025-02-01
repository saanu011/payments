package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Server struct {
	srv  *http.Server
	conf Config
}

func New(conf Config, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Handler:      handler,
			ReadTimeout:  time.Duration(conf.ReadTimeoutMs) * time.Millisecond,
			WriteTimeout: time.Duration(conf.WriteTimeoutMs) * time.Millisecond,
			Addr:         fmt.Sprintf("0.0.0.0:%d", conf.Port),
		},
		conf: conf,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.srv.Addr)
	if err != nil {
		return err
	}

	go func() {
		log.Printf("Starting server on port %d\n", s.conf.Port)

		err := s.srv.Serve(lis)
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start http server")
		}
	}()

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.conf.GracefulShutdownTimeoutMs)*time.Millisecond)
	defer cancel()

	log.Println("Shutting down http server")

	return s.srv.Shutdown(ctx)
}
