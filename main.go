package main

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/scheduler"
	"crawl/LearnGo-crawl/parse/zhengai"
	"crawl/LearnGo-crawl/persist"
)

func main(){

	itemsave,err:= persist.ItemSave()

	if err!=nil{
		panic(err)
	}
	e:= engine.ConcurrentEngine{
		&scheduler.QueueScheduler{},
		100,
		itemsave,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:zhengai.ParseCity,
	})
}




