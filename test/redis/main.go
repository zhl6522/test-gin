package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

//redis

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{		// 不能使用冒号，目标赋值给全局
		Addr:"127.0.0.1:6379",
		Password:"",
		DB:0,
	})
	_,err = redisdb.Ping().Result()
	if err != nil {
	 return err
	}
	return nil
}

func main() {
	err := initRedis()
	if err != nil {
	 fmt.Printf("connect redis failed,err:%v\n",err)
	 return
	}
	fmt.Println("连接redis成功！")
	// zset
	key := "rank"
	items := []redis.Z{
		redis.Z{Score:90, Member:"JavaScript"},
		redis.Z{Score:93, Member:"PHP"},
		redis.Z{Score:95, Member:"Golang"},
		redis.Z{Score:99, Member:"Java"},
	}
	// 把元素追加到key
	num, err :=redisdb.ZAdd(key, items...).Result()
	if err != nil {
	 fmt.Printf("zadd failed, err:%v\n",err)
	 return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 给Golang的分数加10
	newScore, err := redisdb.ZIncrBy(key, "Golang", 10).Result()
	if err != nil {
	 fmt.Printf("zincrby failed, err:%v\n",err)
	 return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)
	// 取分数最高的3个
	ret, err := redisdb.ZRevRange(key, 0, 2).Result()
	if err != nil {
	 fmt.Printf("zrevrange failed, err:%v\n",err)
	 return
	}
	for _,z := range ret{
		fmt.Println(z.Member, z.Score)
	}
}
