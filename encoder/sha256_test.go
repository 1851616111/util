package encoder

import (
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	fmt.Println(Base64(HmacSha256("GETcvm.api.qcloud.com/v2/index.php?Action=DescribeInstances&Nonce=11886&Region=gz&SecretId=AKIDz8krbsJ5yKBZQpn74WFkmLPx3gnPhESA&Timestamp=1465185768&instanceIds.0=ins-09dx96dg&limit=20&offset=0", "Gu5t9xGARNpq86cd98joQYCN3Cozk1qA")))
}
