package token

import (
	"fmt"
	"testing"
)

func TestController_Run(t *testing.T) {
	ctl := NewController("wx700563a0121d82b4", "f3ee49d9f9202d21748c1034b8f49adc")
	if err := ctl.Run(); err != nil {
		t.Fatal(err)
	}
	fmt.Println(ctl.GetToken())
}
