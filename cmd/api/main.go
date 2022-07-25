package main

import (
	"log"
	"sort"
	"strings"

	"hackathon/lib/datamodel"
	"hackathon/lib/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func main() {
	redisDB, err := redis.NewRedisDB()
	if err != nil {
		panic(errors.Wrap(err, "Main NewRedisDB failed"))
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	api := app.Group("/api")

	// * get all topics
	api.Get("/topics", func(c *fiber.Ctx) error {
		var topics []datamodel.Topic

		switch key := strings.ToLower(c.Query("filter")); key {
		case "poll":
			topics, err = redisDB.GetTopics("vote")
		case "question":
			topics, err = redisDB.GetTopics("ask")
		default:
			topics, err = redisDB.GetTopics("*")
		}
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"success": false,
				"data":    errors.Wrap(err, "GetAllTopics failed").Error(),
			})
		}

		topicsMassage := []datamodel.TopicMassage{}
		for _, t := range topics {
			topicsMassage = append(topicsMassage, t.ToMassage())
		}

		switch sortBy := strings.ToLower(c.Query("sort")); sortBy {
		case "count":
			sort.Sort(datamodel.TopicMsgByResponseCount(topicsMassage))
		default:
			sort.Sort(datamodel.TopicMsgByTitle(topicsMassage))
		}

		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"data":    topicsMassage,
		})
	})

	// * get 1 topic
	api.Get("/topic/:topic_address", func(c *fiber.Ctx) error {
		topic, err := redisDB.GetTopic(c.Params("topic_address"))
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"success": false,
				"data":    errors.Wrap(err, "GetTopic failed").Error(),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"data":    topic.ToMassage(),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
