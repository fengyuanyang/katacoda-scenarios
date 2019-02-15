### Deploy Application(bre) to kubernetes
`kubectl create -f deployment.yml`{{execute}}

This command will create a pod, one container built from image 'bre' will be used and run on that pod

### To access Application, must create a service with port on 8080, which is the port tomcat start at (container)
### selector `app: bre` is used in Service, which select pods with labels match `app: bre`

	# deployment.yml
	...
	template:
	  metadata:
	  	labels:
	  	  app: bre
	  	  ...

### Execute command below to create Service
`kubectl create -f service.yml`{{execute}}

### Access Tomcat after deploying Service
https://[[HOST_SUBDOMAIN]]-30001-[[KATACODA_HOST]].environments.katacoda.com/


