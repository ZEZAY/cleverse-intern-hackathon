package main

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"hackathon/lib/contract"
	"hackathon/lib/redis"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

func main() {
	redisDB, err := redis.NewRedisDB()
	if err != nil {
		panic(errors.Wrap(err, "Main NewRedisDB failed"))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	var firstErrorOnce sync.Once
	var firstError error
	m := sync.RWMutex{}

	for {
		time.Sleep(time.Millisecond * 2000)
		// now := time.Now().Format("Jul 26, 2022, 12:00 PM")
		// fmt.Println(now)

		client1, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
		if err != nil {
			panic(errors.Wrap(err, "Main connect node1 failed"))
		}
		client2, err := ethclient.Dial("https://data-seed-prebsc-2-s1.binance.org:8545/")
		if err != nil {
			panic(errors.Wrap(err, "Main connect node2 failed"))
		}
		client3, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
		if err != nil {
			panic(errors.Wrap(err, "Main connect node3 failed"))
		}

		blockNumber, err := client1.BlockNumber(ctx)
		if err != nil {
			panic(errors.Wrap(err, "Main get block number failed"))
		}
		fmt.Println("blockNumber =", blockNumber)
		callOpts := &bind.CallOpts{
			BlockNumber: big.NewInt(int64(blockNumber)),
			Context:     ctx,
		}

		topicAddresses, err := contract.GetAddresses(client1, callOpts)
		if err != nil {
			panic(errors.Wrap(err, "Main GetAddresses failed"))
		}

		clientChal := make(chan *ethclient.Client, 3)
		clientChal <- client1
		clientChal <- client2
		clientChal <- client3

		for _, topicAddress := range topicAddresses {
			var client *ethclient.Client

			select {
			case <-ctx.Done():
				if firstError != nil {
					panic(errors.Wrap(firstError, "Main Update failed"))
				}
				break
			case client = <-clientChal:
				wg.Add(1)
				go func(address common.Address) {
					fmt.Println("address ", address)
					defer wg.Done()
					retryErr := retry.Do(func() error {
						topic, err := contract.GetTopic(client, callOpts, address, clientChal)
						if err != nil {
							panic(errors.Wrap(err, "Main GetTopic failed"))
						}
						m.Lock()
						err = redisDB.SetTopic(*topic)
						if err != nil {
							panic(errors.Wrap(err, "Main SetTopic failed"))
						}
						m.Unlock()
						return nil
					}, retry.Attempts(100), retry.Delay(200*time.Millisecond),
						retry.DelayType(retry.FixedDelay), retry.LastErrorOnly(true),
					)
					if retryErr != nil {
						firstErrorOnce.Do(func() {
							firstError = retryErr
							cancel()
						})
					}
				}(topicAddress)
			}
		}
		wg.Wait()
	}
}
