package activemq

import (
	"testing"
	"time"
)

func TestQ_GetQ(t *testing.T) {
	q := NewMQ("0.0.0.0:32776", "/queue/test-1")
	c, err := q.GetQ()
	if err != nil {
		t.Fatalf("create mq err %v\n", err)
	}

	for {
		select {
		case md := <-c:
			if md.Error != nil {
				t.Fatalf("recv msg err %v\n", md.Error)
			}
			if err := q.Ack(md.Message.Headers); err != nil {
				t.Fatalf("ack msg err %v\n", err)
			}
		case <-time.Tick(time.Millisecond * 200):
			return
		}
	}

}
