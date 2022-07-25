package redis

import (
	"fmt"
	"os"

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

func (r RedisDB) SetTopic(topic datamodel.Topic, topicType string) error {
	value, err := topic.Marshal()
	if err != nil {
		return errors.Wrap(err, "SetTopic marshal failed")
	}

	err = r.Client.Set(topic.GetKey(topicType), value, 0).Err()
	if err != nil {
		return errors.Wrap(err, "SetTopic set failed")
	}
	return nil
}

func (r RedisDB) getTopic(key string) (*datamodel.Topic, error) {
	result, err := r.Client.Get(key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "getTopic get failed")
	}

	topic, err := new(datamodel.Topic).Unmarshal(result)
	if err != nil {
		return nil, errors.Wrap(err, "getTopic unmarshal failed")
	}
	return topic, nil
}

func (r RedisDB) GetTopic(topicAddress string) (*datamodel.Topic, error) {
	var topics []datamodel.Topic
	iter := r.Client.Scan(0, "*-"+topicAddress, 0).Iterator()

	for iter.Next() {
		topic, err := r.getTopic(iter.Val())
		if err != nil {
			return nil, errors.Wrap(err, "GetTopic loop failed")
		}
		topics = append(topics, *topic)
	}

	if err := iter.Err(); err != nil {
		return nil, errors.Wrap(err, "GetTopic read iter failed")
	}

	return &topics[0], nil
}

func (r RedisDB) GetTopics(topicType string) ([]datamodel.Topic, error) {
	var topics []datamodel.Topic
	iter := r.Client.Scan(0, topicType+"-*", 0).Iterator()

	for iter.Next() {
		topic, err := r.getTopic(iter.Val())
		if err != nil {
			return nil, errors.Wrap(err, "GetTopics loop failed")
		}
		topics = append(topics, *topic)
	}

	if err := iter.Err(); err != nil {
		return nil, errors.Wrap(err, "GetTopics read iter failed")
	}

	return topics, nil
}
