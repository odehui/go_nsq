package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	// 配置生产者
	cfg := nsq.NewConfig()
	// 设置消息超时时间（可选）
	cfg.MsgTimeout = 10 * time.Second

	// 连接到本地 nsqd 节点（默认端口 4150）
	producer, err := nsq.NewProducer("127.0.0.1:4150", cfg)
	if err != nil {
		log.Fatalf("创建生产者失败: %v", err)
	}
	defer producer.Stop()

	// 循环发送 5 条测试消息
	for i := 0; i < 5; i++ {
		msg := []byte(fmt.Sprintf("这是第 %d 条测试消息 %s", i+1, time.Now().Format("15:04:05")))
		// 发送消息到 "test_topic" 主题
		err := producer.Publish("test_topic", msg)
		if err != nil {
			log.Printf("发送消息 %d 失败: %v", i+1, err)
			continue
		}
		log.Printf("发送消息 %d 成功: %s", i+1, msg)
		time.Sleep(500 * time.Millisecond) // 间隔发送
	}
}
