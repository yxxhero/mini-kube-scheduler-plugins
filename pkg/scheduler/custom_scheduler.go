package scheduler

import (
	"context"
	"strconv"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	klog "k8s.io/klog/v2"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

// plugin Name
const (
	Name        = "Cks"
	WeightLabel = "weight"
)

type Args struct {
	// prometheus server address
	DefaultWeight int `json:"defaultWeight"`
}

type Cks struct {
	frameworkHandle framework.FrameworkHandle
	args            Args
}

// verify Cks implements the ScorePlugin interface
var _ framework.ScorePlugin = &Cks{}

//New : create a new Cks
func New(config *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	klog.Infof("Create new Cks scheduler")
	args := Args{}
	if err := framework.DecodeInto(config, &args); err != nil {
		klog.Errorf("Load Cks config error: %s", err)
		return nil, err
	}
	klog.Infof("default weight for node", args.DefaultWeight)

	cs := &Cks{
		frameworkHandle: handle,
		args:            args,
	}

	return cs, nil
}

// Score : calculate score
func (k *Cks) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {

	klog.Infof("Calculate score for pod %s on node %s", pod.Name, nodeName)

	// get node info
	nodeInfo, err := k.frameworkHandle.SnapshotSharedLister().NodeInfos().Get(nodeName)
	if err != nil {
		// downgrades Score to zero
		return 0, framework.NewStatus(framework.Success, "downgrades score to zero")
	}

	// get node weight label
	s, ok := nodeInfo.Node().Labels[WeightLabel]

	if !ok {
		klog.Infof("Node %s has no weight label", nodeName)
		return int64(k.args.DefaultWeight), nil
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		klog.Errorf("Parse weight label error: %s", err)
		return int64(k.args.DefaultWeight), nil
	}

	klog.Infof("Calculate score for pod %s in namespace %s on node %s successfully! score: %d", pod.Name, pod.Namespace, nodeName, s)
	return i, framework.NewStatus(framework.Success, "")

}

// Name : get plugin name
func (k *Cks) Name() string {
	return Name
}

// ScoreExtensions : get score extensions
func (k *Cks) ScoreExtensions() framework.ScoreExtensions {
	return k
}

// NormalizeScore : normalize scores
func (k *Cks) NormalizeScore(context.Context, *framework.CycleState, *v1.Pod, framework.NodeScoreList) *framework.Status {
	return nil
}
