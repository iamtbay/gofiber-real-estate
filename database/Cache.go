package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

func redisSet(ctx context.Context, objID string, dataJSON any) error {
	fmt.Println("redis set start")
	key := createRedisKey(objID)
	data, err := json.Marshal(dataJSON)
	if err != nil {
		return errors.New("json couldn't marshal")
	}
	err = rdb.Set(ctx, key, data, 12*time.Hour).Err()
	if err != nil {
		fmt.Println(err)
		return errors.New("redis insert error")
	}
	return nil
}
