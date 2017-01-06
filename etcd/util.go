package etcd

import (
	etcderror "github.com/coreos/etcd/error"
	"github.com/coreos/go-etcd/etcd"
)

func RangeNodeFunc(node *etcd.Node, fn func(*etcd.Node)) {
	if !isDir(node) {
		fn(node)
		return
	}

	for _, child := range node.Nodes {
		if isDir(child) {
			go RangeNodeFunc(child, fn)
		} else {
			fn(child)
		}
	}
}

func isDir(n *etcd.Node) bool {
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
