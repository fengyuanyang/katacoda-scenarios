## Deploy Service and Ingress
### In order to access deployment nginx, we need to create a service
#### Execute command below to create service
`kubectl create -f service.yml`{{execute}}

## Access deployment nginx
For the sake to access deployment nginx via port 80, create a ingress with specific path **test**
then, it's able to access nginx via url pattern such as 
http://xxxxx/test/

#### Execute command below to create Ingress
`kubectl create -f ingress.yml`{{execute}}

#### Check pods status
`kubectl get pods`{{execute}}

`
NAME                   READY     STATUS    RESTARTS   AGE
bre-687989f5bc-5btrt   1/1       Running   0          2m
`

#### Access Application after deploying Ingress
[Access Application](https://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/test/)
