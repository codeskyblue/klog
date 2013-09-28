# klog
[![Build Status](https://drone.io/github.com/shxsun/klog/status.png)](https://drone.io/github.com/shxsun/klog/latest)

**write for golang**
*Developing now.*

First thanks to the 3 people who stars of the project **beelog**. 
This is what that encouraged me to rewrite [beelog](https://github.com/shxsun/beelog), and so **klog** comes out.

klog learn from a lot other log system for golang, like [beego](http://github.com/astaxie/beego), [seelog](https://github.com/cihub/seelog), [glog](https://github.com/golang/glog), [qiniu-log](https://github.com/qiniu/log).

# Introdution
From my experience of some project. I think there are only 4 level needed. `Info Warning Error Fatal`

Default output style is 
```
2006-01-02 03:04:05 [INFO:prefix] hello.go:7  hello world
```
Color output is auto enabled in console, and closed when redirected to a file.

Info colored with green, and Warning is blue, Error is yellow, fatal is red.

# How to use
## ths simple example
```
package main
import "github.com/shxsun/klog"
func main(){
	k := klog.NewLogger(nil, "") // Write to stdout and without prefix
	defer k.Flush() // This will make sure, every thing will output, but I often forgot this.
	k.Infof("Hi %s.", "Susan")
	k.Warning("Oh my god, you are alive!")
	k.Error("Yes, but I will go to Mars tomorrow. So only one day with you")
	k.Fatal("Oh no, donot leave me again... faint")
}
```

More usage please reference <http://gowalker.org/github.com/shxsun/klog>