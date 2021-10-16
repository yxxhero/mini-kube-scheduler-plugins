package main

import (
	"math/rand"
	"os"
	"time"

	"cks/pkg/scheduler"

	"github.com/spf13/pflag"

	cliflag "k8s.io/component-base/cli/flag"

	"k8s.io/component-base/logs"
	_ "k8s.io/component-base/metrics/prometheus/clientgo"
	klog "k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Register custom plugins to the scheduler framework.
	// Later they can consist of scheduler profile(s) and hence
	// used by various kinds of workloads.
	command := app.NewSchedulerCommand(app.WithPlugin(scheduler.Name, scheduler.New))

	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// utilflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.
	// utilflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		klog.Errorf("kube-scheduler exitting with error: %s", err)
		os.Exit(1)
	}
}
