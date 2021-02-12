package framework

import (
	"testing"
	"time"
)

var (
	secondShareSuffix = "-second-share"
)

type TestArgs struct {
	T                   *testing.T
	Name                string
	SecondName          string
	SearchString        string
	MessageString       string
	ShareToDelete       string
	SearchStringMissing bool
	SecondShare         bool
	SecondShareSubDir   bool
	DaemonSetUp         bool
	TestPodUp           bool
	TestDuration        time.Duration
}

// LogAndDebugTestError is not intended as a replacement for the use of t.Fatalf through this e2e suite,
// but when errors occur that could benefit from a dump of the CSI Driver pod logs, use this method instead
// of simply calling t.Fatalf
func LogAndDebugTestError(t *TestArgs) {
	t.T.Logf("*** TEST %s FAILED BEGIN OF CSI DRIVER POD DUMP at time %s", t.T.Name(), time.Now().String())
	dumpCSIPods(t)
	t.T.Logf("*** TEST %s FAILED END OF CSI DRIVER POD DUMP at time %s", t.T.Name(), time.Now().String())
	t.T.Logf("*** TEST %s FAILED BEGIN OF TEST POD DUMP at time %s", t.T.Name(), time.Now().String())
	dumpTestPod(t)
	t.T.Logf("*** TEST %s FAILED END OF TEST POD DUMP at time %s", t.T.Name(), time.Now().String())
	t.T.Logf("*** TEST %s FAILED BEGIN OF TEST EVENT DUMP at time %s", t.T.Name(), time.Now().String())
	dumpTestPodEvents(t)
	t.T.Logf("*** TEST %s FAILED END OF TEST EVENT DUMP at time %s", t.T.Name(), time.Now().String())
	t.T.Fatalf(t.MessageString)
}

//TODO presumably this can go away once we have an OLM based deploy that is also integrated with our CI
// so that repo images built from PRs are used when setting up this driver's daemonset
func LaunchDriver(t *TestArgs) {
	CreateCSIDriverPlugin(t)
}
