package common

import (
	"fmt"
	"github.com/ouncenet/ounced/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunKaspadForTesting runs ounced for testing purposes
func RunKaspadForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	ouncedRunCommand, err := StartCmd("KASPAD",
		"ounced",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Kaspad started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := ouncedRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Kaspad closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := ouncedRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Kaspad stopped")
	}
}
