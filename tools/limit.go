package tools

import (
	"go.uber.org/zap"
	"time"
)

var (
	LimitQueue map[string][]int64
	ok         bool
)

func ClearLimit() {
	go func() {
		zap.L().Info("clearLimitQueue start")
		LimitQueue = nil
		now := time.Now()
		next := now.Add(time.Hour * 24)
		// fmt.Println(next) 2021-04-27 17:56:24.3795757 +0800 CST m=+86400.526372701
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		// fmt.Println(now, next, next.Sub(now)) 2021-04-26 17:56:24.3795757 +0800 CST m=+0.526372701 2021-04-27 00:00:00 +0800 CST 6h3m35.6204243s
		<-t.C
	}()
}

// 滑动窗口限流法
func Limit(name string, count uint, windowTime int64) bool {
	curTime := time.Now().Unix()
	if LimitQueue == nil {
		LimitQueue = make(map[string][]int64)
	}
	if _, ok = LimitQueue[name]; !ok {
		LimitQueue[name] = make([]int64, 0)
	}
	if uint(len(LimitQueue[name])) < count {
		LimitQueue[name] = append(LimitQueue[name], curTime)
		return true
	}

	//队列满了,取出最早访问的时间
	earlyTime := LimitQueue[name][0]
	//说明最早期的时间还在时间窗口内,还没过期,所以不允许通过
	if curTime-earlyTime <= windowTime {
		return false
	} else {
		//说明最早期的访问应该过期了,去掉最早期的
		LimitQueue[name] = LimitQueue[name][1:]
		LimitQueue[name] = append(LimitQueue[name], curTime)
	}
	return true
}
