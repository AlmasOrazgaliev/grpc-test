package app

import (
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"libraryService/internal/config"
	"libraryService/internal/repository"
	"libraryService/internal/service/author"
	"libraryService/internal/service/book"
	"libraryService/internal/service/member"
	"libraryService/pkg/log"
	"libraryService/pkg/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	schema      = "library"
	version     = "1.0.0"
	description = "library-service"
)

func Run() {
	logger := log.LoggerFromContext(context.Background())

	cfg, err := config.New()
	if err != nil {
		logger.Error("ERR_INIT_CONFIG", zap.Error(err))
		return
	}

	repositories, err := repository.New(
		repository.WithMongoStore(cfg.MONGO.DB))
	if err != nil {
		logger.Error("ERR_INIT_REPOSITORY", zap.Error(err))
		return
	}
	defer repositories.Close()

	authorService, err := author.New(author.WithAuthorRepository(repositories.Author))
	if err != nil {
		logger.Error("ERR_INIT_SERVICE", zap.Error(err))
		return
	}
	bookService, err := book.New(book.WithBookRepository(repositories.Book))
	if err != nil {
		logger.Error("ERR_INIT_SERVICE", zap.Error(err))
		return
	}
	memberService, err := member.New(member.WithMemberRepository(repositories.Member))
	if err != nil {
		logger.Error("ERR_INIT_SERVICE", zap.Error(err))
		return
	}

	servers, err := server.New(
		server.WithGRPCServer(cfg.GRPC.Port, authorService, bookService, memberService))
	if err != nil {
		logger.Error("ERR_INIT_SERVER", zap.Error(err))
		return
	}

	// Run our server in a goroutine so that it doesn't block.
	if err = servers.Run(logger); err != nil {
		logger.Error("ERR_RUN_SERVER", zap.Error(err))
		return
	}

	if err != nil {
		logger.Error("ERR_INIT_HANDLER", zap.Error(err))
		return
	}

	// Graceful Shutdown
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the httpServer gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	quit := make(chan os.Signal, 1) // create channel to signify a signal being sent

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	<-quit                                             // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")

	// create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err = servers.Stop(ctx); err != nil {
		panic(err) // failure/timeout shutting down the httpServer gracefully
	}

	fmt.Println("Running cleanup tasks...")
	// Your cleanup tasks go here

	fmt.Println("Server was successful shutdown.")
}
