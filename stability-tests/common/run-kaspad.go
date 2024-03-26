package common

import (
	"fmt"
	"github.com/Kalibriumnet/Kalibrium/domain/dagconfig"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
)

// RunKalibriumdForTesting runs Kalibrium for testing purposes
func RunKalibriumdForTesting(t *testing.T, testName string, rpcAddress string) func() {
	appDir, err := TempDir(testName)
	if err != nil {
		t.Fatalf("TempDir: %s", err)
	}

	kalibriumdRunCommand, err := StartCmd("Kalibrium",
		"Kalibrium",
		NetworkCliArgumentFromNetParams(&dagconfig.DevnetParams),
		"--appdir", appDir,
		"--rpclisten", rpcAddress,
		"--loglevel", "debug",
	)
	if err != nil {
		t.Fatalf("StartCmd: %s", err)
	}
	t.Logf("Kalibrium started with --appdir=%s", appDir)

	isShutdown := uint64(0)
	go func() {
		err := kalibriumdRunCommand.Wait()
		if err != nil {
			if atomic.LoadUint64(&isShutdown) == 0 {
				panic(fmt.Sprintf("Kalibrium closed unexpectedly: %s. See logs at: %s", err, appDir))
			}
		}
	}()

	return func() {
		err := kalibriumdRunCommand.Process.Signal(syscall.SIGTERM)
		if err != nil {
			t.Fatalf("Signal: %s", err)
		}
		err = os.RemoveAll(appDir)
		if err != nil {
			t.Fatalf("RemoveAll: %s", err)
		}
		atomic.StoreUint64(&isShutdown, 1)
		t.Logf("Kalibrium stopped")
	}
}
