package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	app := fiber.New()

	runLogFile, _ := os.OpenFile(
		"log.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	var count int32

	go func() {
		for range time.Tick(time.Second * 10) {
			log.Info().Msg(fmt.Sprintf("requests total =%d", count))
		}
	}()

	app.Get("/", func(c *fiber.Ctx) {
		//atomic.AddInt32(&count, 1)
		count++
		time.Sleep(time.Millisecond * 100)

		c.Send("Hello, World!")
		//atomic.AddInt32(&count, -1)
		count--
	})

	app.Get("/status", func(c *fiber.Ctx) {
		c.Send("I'm fine!")
	})

	app.Listen(8080)
}
