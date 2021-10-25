## mini-kube-scheduler-plugins
The current function is to manually adjust the weight of nodes by playing weight labels for nodes and giving different scores.

### 1.Base Knowledge  
k8s version: v1.18.9
* https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/
* https://github.com/kubernetes-sigs/scheduler-plugins

### 2.write your ownself kube-scheduler-plugins  
current repo contains a score plugins, you can refer to it. 

### 3. kube-scheduler config
```yaml
## config
kube-scheduler --config

```yaml
apiVersion: kubescheduler.config.k8s.io/v1alpha2
kind: KubeSchedulerConfiguration
profiles:
- schedulerName: default-scheduler
  plugins:
    score:
      enabled:
       # score plugin name
       - name: Cks 
  pluginConfig:
  # Cks config
  - name: Cks 
    args:
      defaultWeight: 0 
```

### 4.plugin args handle

```shell
# https://github.com/yxxhero/mini-kube-scheduler-plugins/blob/9097ff78304590e7d4d5462e99719e955b36c995/pkg/scheduler/custom_scheduler.go#L36-L39
if err := framework.DecodeInto(config, &args); err != nil {
	klog.Errorf("Load Cks config error: %s", err)
	return nil, err
}
```

### 5.error handle
when your score plugin failed. you should downgrade score to 0.


