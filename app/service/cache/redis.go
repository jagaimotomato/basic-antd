package cache

import (
	tem "basic-antd/app/model/v1/template"
	"basic-antd/init/global/redis"
	"encoding/json"
	"strconv"
	"time"
)

func GenKey(str string, id int) string {
	return str + "_" + strconv.Itoa(id)
}

func GenListKey(str ...string) (strs string) {
	for _, v := range str {
		strs = strs + v
	}
	return
}

func SetTemplateList(key string, list []tem.Template) (err error) {
	for _, v := range list {
		jsonVal, err := json.Marshal(v)
		if err != nil {
			continue
		}
		redis.Rdb.RPush(redis.Ctx, key, jsonVal)
	}
	redis.Rdb.Expire(redis.Ctx, key, time.Hour*24)
	return
}

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
