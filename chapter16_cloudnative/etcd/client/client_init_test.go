package client

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

// 测试客户端连接
func TestEtcdClientInit(t *testing.T) {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)
	// 客户端配置
	config = clientv3.Config{
		// 节点配置
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
	} else {
		// 输出集群信息
		fmt.Println(client.Cluster.MemberList(context.TODO()))
	}
	client.Close()
}
