package main

import (
	"github.com/shxsun/klog"
)

func main() {
	k := klog.NewLogger(nil, "")
	k.SetLevel(klog.LDebug)
	k.Debugf("Hello")
}
