package subscribe

import (
	"testing"

	"subscribe/configs"

	"github.com/spf13/viper"
)

// go test -v --run TestAccountSubscribe
func TestAccountSubscribe(t *testing.T) {
	configs.SetDevEnv()

	url := viper.GetString("rpc.dev.ws")
	pubKey := "4PkiqJkUvxr9P8C1UsMqGN8NJsUcep9GahDRLfmeu8UK"

	if err := AccountSubscribe(url, pubKey); err != nil {
		t.Error(err)
	}
}
