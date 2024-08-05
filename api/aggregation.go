package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Useing Redis for storing intermediate state and performing quick lookups

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func recordClickInRedis(adID string) {
	err := rdb.Incr(ctx, adID).Err()
	if err != nil {
		panic(err)
	}
}

func getTopAds() {
	vals, err := rdb.ZRevRangeWithScores(ctx, "ads_rank", 0, 99).Result()
	if err != nil {
		panic(err)
	}

	// process top ads
	for _, v := range vals {
		fmt.Println(v.Member, v.Score)
	}
}

func aggregateClicks() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		getTopAds()
		// reset for next interval
		rdb.FlushAll(ctx)
	}
}
