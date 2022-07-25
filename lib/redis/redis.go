package redis

import (
	"fmt"
	"os"
	"strings"

	"hackathon/lib/datamodel"
	"hackathon/lib/utils"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type RedisDB struct {
	Client *redis.Client
}

func NewRedisDB() (*RedisDB, error) {
	err := utils.LoadEnv()
	if err != nil {
		return nil, errors.Wrap(err, "NewRedisDB file")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})

	_, err = client.Ping().Result()
	if err != nil {
		return nil, errors.New("RedisClient create client failed")
	}

	redisDB := RedisDB{
		Client: client,
	}
	return &redisDB, nil
}

func (r RedisDB) SetTopic(topic datamodel.Topic) error {
	value, err := topic.Marshal()
	if err != nil {
		return errors.Wrap(err, "SetTopic marshal failed")
	}

	err = r.Client.Set(strings.ToLower(topic.Address.String()), value, 0).Err()
	if err != nil {
		return errors.Wrap(err, "SetTopic set failed")
	}
	return nil
}

func (r RedisDB) GetTopic(topicAddress string) (*datamodel.Topic, error) {
	result, err := r.Client.Get(topicAddress).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic get failed")
	}

	topic, err := new(datamodel.Topic).Unmarshal(result)
	if err != nil {
		return nil, errors.Wrap(err, "GetTopic unmarshal failed")
	}
	return topic, nil
}

func (r RedisDB) GetAllTopics() ([]datamodel.Topic, error) {
	var topics []datamodel.Topic
	iter := r.Client.Scan(0, "*", 0).Iterator()

	for iter.Next() {
		topic, err := r.GetTopic(iter.Val())
		if err != nil {
			return nil, errors.Wrap(err, "GetAllTopics loop failed")
		}
		topics = append(topics, *topic)
	}

	if err := iter.Err(); err != nil {
		return nil, errors.Wrap(err, "GetAllTopics read iter failed")
	}

	return topics, nil
}
