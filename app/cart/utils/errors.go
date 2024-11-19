package utils

import "github.com/cloudwego/kitex/pkg/klog"

func MustHandlerError(err error) {
	if err != nil {
		klog.Fatal(err)
	}
}
