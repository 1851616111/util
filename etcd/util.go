package etcd

import (
	"fmt"
	etcderror "github.com/coreos/etcd/error"
	"github.com/coreos/go-etcd/etcd"
)

func RangeNodeFunc(node *etcd.Node, fn func(*etcd.Node)) {
	if !IsDir(node) {
		fn(node)
		return
	}

	for _, child := range node.Nodes {
		if IsDir(child) {
			go RangeNodeFunc(child, fn)
		} else {
			fn(child)
		}
	}
}

func ListDir(rsp *etcd.Response) []string {
	dirNodeValues := []string{}

	for _, node := range rsp.Node.Nodes {
		dirNodeValues = append(dirNodeValues, node.Value)
	}
	return dirNodeValues
}

func IsDir(n *etcd.Node) bool {
	return n.Dir
}

func AlreadyExistErr(err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*etcd.EtcdError); ok && e.ErrorCode == etcderror.EcodeNodeExist {
		return true
	}

	return false
}

func NotDirErr(err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*etcd.EtcdError); ok && e.ErrorCode == etcderror.EcodeNotDir {
		return true
	}

	return false
}

func KeyNotFound(err error) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*etcd.EtcdError); ok && e.ErrorCode == etcderror.EcodeKeyNotFound {
		return true
	}

	return false
}

func receiver(c chan *etcd.Response, stop chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	stop <- true
}
