package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/reflection"
	"libraryService/internal/service/author"
	"libraryService/internal/service/book"
	"libraryService/internal/service/member"
	"net"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	desc "libraryService/proto"
)

type Server struct {
	http     *http.Server
	grpc     *grpc.Server
	listener net.Listener
}

// Configuration is an alias for a function that will take in a pointer to a Repository and modify it
type Configuration func(r *Server) error

// New takes a variable amount of Configuration functions and returns a new Server
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (r *Server, err error) {
	// Create the Server
	r = &Server{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(r); err != nil {
			return
		}
	}
	return
}

func (s *Server) Run(logger *zap.Logger) (err error) {
	if s.http != nil {
		go func() {
			if err = s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("ERR_INIT_REST", zap.Error(err))
				return
			}
		}()
		logger.Info("http server started on http://localhost" + s.http.Addr)
	}

	if s.grpc != nil {
		go func() {
			if err = s.grpc.Serve(s.listener); err != nil {
				return
			}
		}()
		logger.Info("grpc server started on http://localhost" + s.listener.Addr().String())
	}

	return
}

func (s *Server) Stop(ctx context.Context) (err error) {
	if s.http != nil {
		if err = s.http.Shutdown(ctx); err != nil {
			return
		}
	}

	return
}

func WithGRPCServer(port string, author *author.Service, book *book.Service, member *member.Service) Configuration {
	return func(s *Server) (err error) {
		s.listener, err = net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
		if err != nil {
			return
		}
		s.grpc = grpc.NewServer()
		reflection.Register(s.grpc)
		desc.RegisterAuthorServer(s.grpc, author)
		desc.RegisterBookServer(s.grpc, book)
		desc.RegisterMemberServer(s.grpc, member)
		return
	}
}

func WithHTTPServer(handler http.Handler, port string) Configuration {
	return func(s *Server) (err error) {
		s.http = &http.Server{
			Handler: handler,
			Addr:    ":" + port,
		}
		return
	}
}
