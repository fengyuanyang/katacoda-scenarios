Create deployment nginx via **command** below

`kubectl run nginx --image=nginx`{{execute}}

In few minutes, check if deployment **nginx** has been deployed 
#### Check deployments status

`kubectl get deployments`{{execute}}

One deployments named **nginx** should be listed as below

`
NAME   READY   UP-TO-DATE   AVAILABLE   AGE
nginx  1/1     1            1           13s
`