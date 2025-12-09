package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/ameijboom/chime/generated"
	"github.com/ameijboom/chime/internal/bot"
	"github.com/ameijboom/chime/internal/repository"
	"github.com/ameijboom/chime/internal/server"
	"google.golang.org/grpc"
)

var (
	GuildID        = flag.String("guild", "", "Guild ID for registering slash commands.")
	Token          = flag.String("token", "", "Discord Access Token.")
	DatabaseURL    = flag.String("database", "chime.db", "Database URL")
	GRPCPort       = flag.String("grpc-port", ":50051", "gRPC server port")
	RemoveCommands = flag.Bool("remove-commands", false, "Removes commands upon shutdown")
)

func init() {
	flag.Parse()
}

func main() {
	// Initialize database
	repo, err := repository.NewDatabase(*DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := repo.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize Discord bot
	discordBot, err := bot.NewBot(*Token, *GuildID)
	if err != nil {
		log.Fatalf("Failed to create Discord Bot: %v", err)
	}

	err = discordBot.Open()
	if err != nil {
		log.Fatalf("Failed to start Discord Bot: %v", err)
	}
	defer discordBot.Close()

	// log.Println("Discord bot started successfully")

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	eventsServer := server.NewServer(repo, discordBot)
	generated.RegisterEventsServer(grpcServer, eventsServer)

	// Start gRPC server in goroutine
	listener, err := net.Listen("tcp", *GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", *GRPCPort, err)
	}

	go func() {
		log.Printf("gRPC server listening on %s", *GRPCPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Wait for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down...")
	grpcServer.GracefulStop()
}
