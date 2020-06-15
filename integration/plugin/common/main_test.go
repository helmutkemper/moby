package common // import "github.com/helmutkemper/moby/integration/plugin/common"

import (
	"fmt"
	"os"
	"testing"

	"github.com/helmutkemper/moby/pkg/reexec"
	"github.com/helmutkemper/moby/testutil/environment"
)

var testEnv *environment.Execution

func TestMain(m *testing.M) {
	if reexec.Init() {
		return
	}
	var err error
	testEnv, err = environment.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	testEnv.Print()
	os.Exit(m.Run())
}

func setupTest(t *testing.T) func() {
	environment.ProtectAll(t, testEnv)
	return func() { testEnv.Clean(t) }
}
