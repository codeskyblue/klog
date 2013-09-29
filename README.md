# klog
[![Build Status](https://drone.io/github.com/shxsun/klog/status.png)](https://drone.io/github.com/shxsun/klog/latest)

**library for golang**

First thanks to the 3 people who stars of the project **beelog** 
which encouraged me to rewrite [beelog](https://github.com/shxsun/beelog), and **klog** comes out.

klog learn from a lot other log system for golang, like [beego](http://github.com/astaxie/beego), [seelog](https://github.com/cihub/seelog), [glog](https://github.com/golang/glog), [qiniu-log](https://github.com/qiniu/log).

# Introdution
From my experience of some project. I think there are only 5 level needed. `Debug Info Warning Error Fatal`

Default level is Info. Use `SetLevel(level)` to change.

Default output style is 
```
prefix 2006/01/02 03:04:05 [INFO] hello world
```
Use `SetFlags(klog.Fdevflag)`  will change output to 
```
prefix 2006/01/02 03:04:05 [INFO] hello.go:7 hello world
```

Color output is default enabled in console, and auto closed when redirected to a file.
```
DEBUG  "Cyan",
INFO   "Green",
WARN   "Magenta",
ERROR  "Yellow",
FATAL  "Red",
```

# How to use
**below this is a simple example**

```
package main
import "github.com/shxsun/klog"
func main(){
	k := klog.NewLogger(nil, "") // Write to stdout and without prefix

	k.Infof("Hi %s.", "Susan")
	k.Warning("Oh my god, you are alive!")
	k.Error("Yes, but I will go to Mars tomorrow. So only one day with you")
	k.Fatal("Oh no, donot leave me again... faint") // Fatal will call os.Exit(1)
}
```

More usage please reference <http://gowalker.org/github.com/shxsun/klog>
