package activemq

import (
	mq "github.com/gmallard/stompngo"
	"net"
	"sync"
)

type MQ interface {
	GetQ() (<-chan mq.MessageData, error)
	Ack(mq.Headers) error
}

type q struct  {
	target string
	topic string

	once *sync.Once
	c *mq.Connection
	ch <-chan mq.MessageData
	err error
}

func NewMQ(target, topic string) MQ {
	return &q{
		target: target,
		topic: topic,
		once: new(sync.Once),
	}
}


func (p *q) GetQ() (<-chan mq.MessageData, error) {
	p.once.Do(func(){
		conn, e := net.Dial("tcp", p.target)
		if e != nil {
			p.err = e
			return
		}

		// STOMP 1.0 的标准头
		//h := stompngo.Headers{}
		// STOMP 1.1 的标准头
		h := mq.Headers{"accept-version", "1.1"}
		// @todo 强化网络断开之后重试
		c, e := mq.Connect(conn, h)
		if e != nil {
			p.err = e
			return
		}

		c.SetSubChanCap(1)
		p.c = c

		// 必须客户端响应才可以删除MQ队列数据
		f := mq.Headers{"destination", p.topic, "ack", "client"}
		// 自动删除MQ队列的数据
		//f := stompngo.Headers{"destination", "/queue/bbg_ordercache"}
		if ch, e := c.Subscribe(f); e != nil {
			p.err = e
		} else {
			p.ch = ch
		}

		return
	})

	if p.err != nil {
		return nil, p.err
	}

	return p.ch, nil
}

func (p *q) Ack(h mq.Headers) error {
	return p.c.Ack(h)
}

