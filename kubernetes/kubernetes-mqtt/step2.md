
Deploy Mqtt on Port 1833 with helm    
[Install mqtt](https://hub.helm.sh/charts/halkeye/mosquitto)

* Add halkeye repository    
 `helm repo add halkeye https://halkeye.github.io/helm-charts/`{{execute}}
* Install Mqtt chart    
install chart with a specific name.     
command - `helm install [NAME] [CHART] [flags]`       
`helm install mosquitto halkeye/mosquitto --version 0.2.0`{{execute}}


Check deployments status

`kubectl get deployments`{{execute}}
`
NAME   READY   UP-TO-DATE   AVAILABLE   AGE
mosquitto   1/1     1            1           42s
`
