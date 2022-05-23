package main

import (
	"errors"
	"flag"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)

	flag.Set("v", "3")
	flag.Parse()

	defer klog.Flush()

	err := errors.New("fail")
	klog.Error("This is error message")
	klog.Errorf("Log using Errorf, err: %v", err)
	klog.ErrorS(err, "Log using ErrorS")

	klog.Info("This is info message")
	klog.Infof("This is info message: %v", 12345)
	klog.InfoDepth(1, "This is info message", 12345)

	klog.Warning("This is warning message")
	klog.Warningf("This is warning message: %v", 12345)
	klog.WarningDepth(1, "This is warning message", 12345)

	klog.Error("This is error message")
	klog.Errorf("This is error message: %v", 12345)
	klog.ErrorDepth(1, "This is error message", 12345)

	klog.V(3).Info("LEVEL 3 message")
	klog.V(4).Infof("LEVEL %s message", "4")
	klog.V(5).InfoS("LEVEL 5 message", "level", "5")
	//klog.Fatal("This is fatal message")
	//klog.Fatalf("This is fatal message: %v", 12345)
	//klog.FatalDepth(1, "This is fatal message", 12345)
}
