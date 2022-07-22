package client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestKV(t *testing.T) {
	// 获取上下文
	rootContext := context.Background()
	// 客户端初始化
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	defer cli.Close()
	// etcd clientv3 >= v3.2.10, grpc/grpc-go >= v1.7.3
	if cli == nil || err == context.DeadlineExceeded {
		// handle errors
		fmt.Println(err)
		panic("invalid connection!")
	}
	kvc := clientv3.NewKV(cli)
	//设置值
	uuid := uuid.New().String()
	fmt.Printf("new value is :%s\r\n", uuid)
	// 设置上下文，超时为：2s
	ctx1, cancelFunc1 := context.WithTimeout(rootContext, time.Duration(2)*time.Second)
	kvc.Put(ctx1, "cc", uuid)
	cancelFunc1()
	// 获取值
	ctx2, cancelFunc2 := context.WithTimeout(rootContext, time.Duration(2)*time.Second)
	resp, _ := kvc.Get(ctx2, "cc")
	for _, kv := range resp.Kvs {
		fmt.Println(string(kv.Value))
	}
	// 删除
	if delResp, err := kvc.Delete(ctx2, "cc"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("delete %s for %t\n", "cc", delResp.Deleted > 0)
	}

	cancelFunc2()
}
