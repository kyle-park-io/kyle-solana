package subscribe

import (
	"testing"

	"subscribe/configs"

	"github.com/spf13/viper"
)

// go test -v --run TestProgramSubscribe
func TestProgramSubscribe(t *testing.T) {
	configs.SetDevEnv()

	url := viper.GetString("rpc.dev.ws")
	program := "2Ym3QkbXGEZSLDSERE6zCuar6fMCHTzvmw2He3MSL1s9"

	if err := ProgramSubscribe(url, program); err != nil {
		t.Error(err)
	}
}
