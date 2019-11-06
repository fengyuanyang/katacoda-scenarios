#### Deploy Dashboard for kubernetes
`kubectl apply -f kubernetes-dashboard.yaml`{{execute}}


#### Deploy Application(bre) to kubernetes
`kubectl apply -f deployment.yml`{{execute}}

#### To access Dashboard, you must create a secure channel to your Kubernetes cluster. Run the following command
`kubectl proxy --address='0.0.0.0' --port=8001 --accept-hosts='^*$' &`{{execute}}

#### Access Dashboard after deploying Dashboard 
[Access Dashboard](https://[[HOST_SUBDOMAIN]]-8001-[[KATACODA_HOST]].environments.katacoda.com/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/overview?namespace=default)

#### Check deployments status
`kubectl get deployments`{{execute}}
`
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
bre       1         1         1            1           5m
`

This command will create a pod, one container built from image 'bre' will be used and run on that pod

#### Create a service on 8888 port to access Application , which is the port tomcat start at (container)

selector `app: bre` is used in Service, which select pods with labels match `app: bre`

	# deployment.yml
	...
	template:
	  metadata:
	  	labels:
	  	  app: bre
	  	  ...

#### Execute command below to create Service
`kubectl apply -f service.yml`{{execute}}



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
[Access Application](https://[[HOST_SUBDOMAIN]]-30000-[[KATACODA_HOST]].environments.katacoda.com/)



