#### In order to access deployment nginx, we need to create a service

Execute command below to create service
`kubectl create -f service.yml`{{execute}}

#### Check Service
`kubectl get svc`{{execute}}

A service named nginx should be listed as below
`
NAME       TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                  AGE
...
nginx      NodePort       10.102.194.141   <none>        80:32462/TCP                 19s
`

## Access deployment nginx
For the sake to access deployment nginx via port 80, create a ingress with specific path **test**
then, it's able to access nginx via url pattern such as 
http://xxxxx/test/

#### Create Ingress via command
`kubectl create -f ingress.yml`{{execute}}


#### Check Ingress
`kubectl get ingress`{{execute}}

`
NAME           HOSTS   ADDRESS   PORTS   AGE
test-ingress   *                 80      53s
`

#### Access Application after deploying Ingress
[Access Application](https://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com/test/)
