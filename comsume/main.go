package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
)

type TestHandler struct{}

var msg_chan_name = flag.String("msg_chan_name", "default", "消息通道名称")

func (h *TestHandler) HandleMessage(msg *nsq.Message) error {
	if len(msg.Body) == 0 {
		return fmt.Errorf("空消息")
	}
	// 模拟业务处理（耗时 1 秒）
	time.Sleep(1 * time.Second)
	log.Printf("收到消息: %s (消息ID: %s)", msg.Body, msg.ID)
	return nil // 返回 nil 表示处理成功，NSQ 会标记消息为已完成
}

func main() {
	flag.Parse()
	// 配置消费者
	cfg := nsq.NewConfig()
	// 每 15 秒从 lookupd 刷新节点信息（可选）
	cfg.LookupdPollInterval = 15 * time.Second

	// 创建消费者：订阅 "test_topic" 主题的 "test_channel" 通道
	consumer, err := nsq.NewConsumer("test_topic", *msg_chan_name, cfg)
	if err != nil {
		log.Fatalf("创建消费者失败: %v", err)
	}

	// 设置消息处理器
	consumer.AddHandler(&TestHandler{})

	// 从 lookupd 发现 nsqd 节点（使用 lookupd 的 HTTP 端口 4161）
	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Fatalf("连接 lookupd 失败: %v", err)
	}

	// 阻塞等待（保持消费者运行）
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
