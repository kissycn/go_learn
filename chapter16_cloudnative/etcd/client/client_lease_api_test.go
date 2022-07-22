package client

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestLeaseApi(t *testing.T) {
	// 客户端初始化
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Close()

	// 创建一个Lease Service
	leaseCli := clientv3.NewLease(cli)
	// 创建一个10s的租约
	leaseGrantResp, _ := leaseCli.Grant(context.TODO(), 10)
	leaseId := leaseGrantResp.ID
	// 创建一个KV Service
	kvCli := clientv3.NewKV(cli)
	if _, err := kvCli.Put(context.TODO(), "/school/students", "jerry", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}

	for {
		if getResp, err := kvCli.Get(context.TODO(), "/school/students"); err != nil {
			fmt.Println(err)
			return
		} else if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		} else {
			fmt.Println("还没过期:", getResp.Kvs)
		}
		time.Sleep(2 * time.Second)
	}
}
