## mini-kube-scheduler-plugins

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
```

### 3.plugin args handle


### 4.error handle
when your score plugin failed. you should downgrade score to 0.

### 5.unittest
### 6.debug


