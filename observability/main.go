package main

import (
	"context"
	"github.com/KTo1/ozon/observability/server"
	"github.com/KTo1/ozon/observability/storage"
	"github.com/KTo1/ozon/observability/tracer"
	"github.com/gofiber/contrib/otelfiber"
	fiber "github.com/gofiber/fiber/v2"
	redisotel "github.com/redis/go-redis/extra/redisotel/v9"
	redis "github.com/redis/go-redis/v9"
	"log"
)

func main() {
	app := fiber.New()

	tr, err := tracer.InitTracer("http://localhost:14268/api/traces", "Note service")
	if err != nil {
		log.Fatal("init jaeger failed: ", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		log.Fatal("redis client create failed", err)
	}

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(client); err != nil {
		panic(err)
	}

	app.Use(otelfiber.Middleware(otelfiber.WithServerName("my-server")))

	handler := server.NewFiberHandler(storage.NewNoteStorage(client), tr)

	app.Post("/v1/create", handler.CreateNote)
	app.Get("/v1/get", handler.GetNote)

	log.Fatal(app.Listen(":3001"))
}
