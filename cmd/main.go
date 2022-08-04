package main

import (
	"log"
	"sort"
	"strings"

	"hackathon/internal/database"
	"hackathon/internal/datamodel"
	"hackathon/internal/pkg/poller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pkg/errors"
)

func main() {
	redisDB, err := database.NewRedisDB()
	if err != nil {
		panic(errors.Wrap(err, "Main NewRedisDB failed"))
	}

	postgresDB, err := database.NewPostgresDB()
	if err != nil {
		panic(errors.Wrap(err, "Main NewPostgresDB failed"))
	}

	app := fiber.New()

	// app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length, ngrok-skip-browser-warning",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

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
				"data":    errors.Wrap(err, "Get topics failed").Error(),
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
				"data":    errors.Wrap(err, "Get topic failed").Error(),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"data":    topic.ToMassage(),
		})
	})

	// * create new organization
	api.Post("/organization/new", func(c *fiber.Ctx) error {
		var organization datamodel.Organization

		if err := c.BodyParser(&organization); err != nil {
			return c.Status(404).JSON(fiber.Map{
				"success": false,
				"data":    errors.Wrap(err, "Create organization BodyParser failed").Error(),
			})
		}

		err = postgresDB.AddOrganization(organization)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"success": false,
				"data":    errors.Wrap(err, "Create organization Add failed").Error(),
			})
		}

		return c.Status(201).JSON(&fiber.Map{
			"success": true,
			"data":    organization.ToMassage(),
		})
	})

	// * get all organizations
	api.Get("/organizations", func(c *fiber.Ctx) error {
		organizations, err := postgresDB.GetOrganizations()
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"success": false,
				"data":    errors.Wrap(err, "Get organizations failed").Error(),
			})
		}

		organizationsMassage := []datamodel.OrganizationMassage{}
		for _, o := range organizations {
			organizationsMassage = append(organizationsMassage, o.ToMassage())
		}

		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"data":    organizationsMassage,
		})
	})

	// * get 1 organization
	api.Get("/organization/:organization_address", func(c *fiber.Ctx) error {
		key := c.Params("organization_address")
		organization, err := postgresDB.GetOrganization(key)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"success": false,
				"data":    errors.Wrap(err, "Get organization failed").Error(),
			})
		}

		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"data":    organization.ToMassage(),
		})
	})

	go poller.RunLoop()
	log.Fatal(app.Listen(":3000"))
}
