package etcd

import (
	"encoding/json"
	"github.com/coreos/go-etcd/etcd"
)

//addr = schema://ip:port
func  NewClient(addr, user, password string) Interface {
	cli := etcd.NewClient([]string{addr})
	cli.SetCredentials(user, password)
	return &storage{cli}
}

type Interface interface {
	SetString(key, value string) error
	SetObject(key string, obj interface{}) error
	GetString(key string) (string, error)
	//getValue(key string) (string, error)
	//getDir(key string) (*etcd.Response, error)
	//delete(key string, recursive bool) error
}

type storage struct {
	*etcd.Client
}

func (c *storage) SetString(key, value string) error {
	if _, err := c.Set(key, value, 0); err != nil {
		return err
	}
	return nil
}

func (c *storage) SetObject(key string, obj interface{}) error {
	b, _ := json.Marshal(obj)
	if _, err := c.Set(key, string(b), 0); err != nil {
		return err
	}
	return nil
}

func (c *storage) GetString(key string) (string, error) {
	rsp, err := c.Get(key, true, false)
	if err != nil {
		return "", err
	}

	return rsp.Node.Value, nil
}

//func (c *Interface) getValue(key string) (string, error) {
//	var rsp *etcd.Response
//	var err error
//
//	err = notReachErrRetry(func(c *Interface) error {
//		rsp, err = c.Get(key, true, false)
//		return err
//	})
//
//	if err != nil {
//		return "", err
//	}
//
//	return rsp.Node.Value, nil
//}
//
//func (c *Interface) getDir(key string) (*etcd.Response, error) {
//	var rsp *etcd.Response
//	var err error
//
//	err = notReachErrRetry(func(c *Interface) error {
//		rsp, err = c.Get(key, true, true)
//		return err
//	})
//
//	if err != nil {
//		return nil, err
//	}
//	return rsp, nil
//}
//
//func (c *Interface) delete(key string, recursive bool) error {
//	var rsp *etcd.Response
//	var err error
//
//	err = notReachErrRetry(func(c *Interface) error {
//		rsp, err = c.Delete(key, recursive)
//		return err
//	})
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//func notReachErrRetry(f func(c *Interface) error) (err error) {
//	err = f(db.(*Interface))
//
//	if isEtcdNotReachableErr(err) {
//		refreshDB()
//		err = f(db.(*Interface))
//
//		if isEtcdNotReachableErr(err) {
//			err = errors.New("Server Internal Error")
//			return
//		}
//	}
//
//	return
//}
//
//func isEtcdNotReachableErr(err error) bool {
//	if err == nil {
//		return false
//	}
//
//	if e, ok := err.(*etcd.EtcdError); ok && e.ErrorCode == etcd.ErrCodeEtcdNotReachable {
//		return true
//	}
//
//	return false
//}

