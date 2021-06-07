package cache

import (
	redis "basic-antd/init/global"
	"encoding/json"
	"time"
)

func GetList(key string, max, min int64) (list []interface{}, count int64, err error) {
	listString, err := redis.Rdb.LRange(redis.Ctx, key, min, max).Result()
	if err != nil {
		return
	}
	count, err = redis.Rdb.LLen(redis.Ctx, key).Result()
	if err != nil {
		return
	}
	var data interface{}
	for _, v := range listString {
		err = json.Unmarshal([]byte(v), &data)
		list = append(list, data)
	}
	return
}

func Del(key string) (err error) {
	err = redis.Rdb.Del(redis.Ctx, key).Err()
	if err != nil {
		return
	}
	return
}

func IsExist(key string) (isExist int64, err error) {
	isExist, err = redis.Rdb.Exists(redis.Ctx, key).Result()
	if err != nil {
		return
	}
	return
}

func Set(key string, value interface{}, t time.Duration) (err error) {
	redisVal, err := json.Marshal(value)
	if err != nil {
		return
	}
	err = redis.Rdb.Set(redis.Ctx, key, redisVal, t).Err()
	if err != nil {
		return
	}
	return
}

func Get(key string) (value string, err error) {
	value, err = redis.Rdb.Get(redis.Ctx, key).Result()
	if err != nil {
		return
	}
	return
}

func LPush(key string, value interface{}) (err error) {
	err = redis.Rdb.LPush(redis.Ctx, key, value).Err()
	if err != nil {
		return
	}
	return
}

func LRem(key string, count int64, value interface{}) (err error) {
	err = redis.Rdb.LRem(redis.Ctx, key, count, value).Err()
	if err != nil {
		return
	}
	return
}

func SetExpire(key string, expire time.Duration) (err error) {
	err = redis.Rdb.Expire(redis.Ctx, key, expire).Err()
	if err != nil {
		return
	}
	return
}

func GetListLength(key string) (length int64, err error) {
	length, err = redis.Rdb.LLen(redis.Ctx, key).Result()
	if err != nil {
		return
	}
	return
}
