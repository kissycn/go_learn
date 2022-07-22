package client

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestWatch(t *testing.T) {
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

	kvCli := clientv3.NewKV(cli)

	// 模拟KV的变化
	go func() {
		for {
			kvCli.Put(context.TODO(), "school/students", "jerry")
			kvCli.Delete(context.TODO(), "school/students")
			time.Sleep(1 * time.Second)
		}
	}()

	// 先GET到当前的值，并监听后续变化
	if getResp, err := kvCli.Get(context.TODO(), "school/students"); err != nil {
		fmt.Println(err)
		return
	} else {
		if len(getResp.Kvs) > 0 {
			fmt.Println("当前值:", string(getResp.Kvs[0].Value))
		}
		// 获得当前revision
		startVer := getResp.Header.Revision + 1
		watcher := clientv3.NewWatcher(cli)
		fmt.Println("从该版本向后监听:", startVer)
		// 创建上下文
		ctx, cancelFunc := context.WithCancel(context.TODO())
		time.AfterFunc(5*time.Second, func() {
			cancelFunc()
		})
		watchRespChan := watcher.Watch(ctx, "school/students", clientv3.WithRev(startVer))
		for watchResp := range watchRespChan {
			for _, event := range watchResp.Events {
				switch event.Type {
				case mvccpb.PUT:
					fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
				case mvccpb.DELETE:
					fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
				}
			}
		}
	}
}
