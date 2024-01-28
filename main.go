package main

import (
	"fmt"
	"github.com/andhikasamudra/fiber-starter-pack/adapter/postgres"
	"github.com/andhikasamudra/fiber-starter-pack/adapter/redis"
	"github.com/andhikasamudra/fiber-starter-pack/internal/env"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/auth"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book"
	fiberstarterpack "github.com/andhikasamudra/fiber-starter-pack/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	grpcServer()
	run()
}

func grpcServer() {
	go func() {
		pgDB := postgres.NewAdapter()
		pgDB.Connect()
		pgDB.Db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
		defer pgDB.Close()

		grpcServer := grpc.NewServer()
		fiberstarterpack.RegisterAuthServiceServer(grpcServer, auth.InitRPCServerRoute(pgDB))

		listener, err := net.Listen("tcp", fmt.Sprintf(":%s", env.GRPCServerPort())) // Use the desired port for gRPC
		if err != nil {
			log.Fatalf("Failed to listen for gRPC: %v", err)
		}
		log.Printf("gRPC server started at %s", env.GRPCServerPort())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()
}

func run() {
	app := fiber.New()
	app.Use(cors.New())

	rc := redis.NewAdapter()
	pgDB := postgres.NewAdapter()
	pgDB.Connect()
	pgDB.Db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	defer pgDB.Close()

	grpcConn, err := grpc.Dial(fmt.Sprintf("%s:%s", "localhost", env.GRPCServerPort()), grpc.WithTransportCredentials(insecure.NewCredentials())) // Replace with your server address
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer grpcConn.Close()

	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("services is up"))
	})
	app.Get("/ready", func(c *fiber.Ctx) error {
		return c.SendString("ready")
	})
	api := app.Group("")
	auth.InitRoute(api, pgDB, rc.Client)
	book.InitRoute(api, pgDB, rc.Client)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var serverShutdown sync.WaitGroup

	go func() {
		_ = <-c //nolint:all
		fmt.Println("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = app.ShutdownWithTimeout(60 * time.Second)
	}()

	// ...

	if err := app.Listen(os.Getenv("APP_PORT")); err != nil {
		log.Panic(err)
	}

	serverShutdown.Wait()

	fmt.Println("Running cleanup tasks...")
}
