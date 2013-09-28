package klog

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestDebugf(t *testing.T) {
	out := bytes.NewBuffer(nil)
	k := NewLogger(out, "")
	k.SetLevel(LDebug)
	k.Debugf("hello %s", "klog")
	outStr := string(out.Bytes())
	if !strings.HasSuffix(outStr, "hello klog\n") {
		t.Errorf("expect suffix with %s, but receive %s",
			strconv.Quote("hello klog"),
			strconv.Quote(outStr))
	}
}
