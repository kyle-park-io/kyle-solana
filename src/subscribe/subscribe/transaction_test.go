package subscribe

import (
	"testing"

	"subscribe/configs"

	"github.com/spf13/viper"
)

// go test -v --run TestSignatureSubscribe
func TestSignatureSubscribe(t *testing.T) {
	configs.SetDevEnv()

	url := viper.GetString("rpc.dev.ws")
	txHash := "3wPggj9kwXMnHYs4rStuXWAaSj5Z4MAcnbVBWNacYVEpGGJ1AYExsVhhSjRB3JeqQEGWvBkeoduoJABb4vtBGA1s"

	if err := SignatureSubscribe(url, txHash); err != nil {
		t.Error(err)
	}
}
