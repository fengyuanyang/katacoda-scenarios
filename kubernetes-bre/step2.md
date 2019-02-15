#### Deploy Application(bre) to kubernetes
`kubectl create -f deployment.yml`{{execute}}

This command will create a pod, one container built from image 'bre' will be used and run on that pod

#### To access Application, must create a service with port on 8080, which is the port tomcat start at (container)
#### selector `app: bre` is used in Service, which select pods with labels match `app: bre`

	# deployment.yml
	...
	template:
	  metadata:
	  	labels:
	  	  app: bre
	  	  ...

#### Execute command below to create Service
`kubectl create -f service.yml`{{execute}}

#### Check deployments status
`kubectl get deployments`{{execute}}
`
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
bre       1         1         1            1           5m
`

#### Check pods status
`kubectl get pods`{{execute}}

`
NAME                   READY     STATUS    RESTARTS   AGE
bre-687989f5bc-5btrt   1/1       Running   0          2m
`
Repeat if the STATUS not shows `Running`

#### Check services status
`kubectl get svc`{{execute}}
`
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
bre          NodePort    10.98.14.144   <none>        8090:30001/TCP   6m
kubernetes   ClusterIP   10.96.0.1      <none>        443/TCP          7m
`

#### Access Application after deploying Service
https://[[HOST_SUBDOMAIN]]-30001-[[KATACODA_HOST]].environments.katacoda.com/


