
Deploy Application to kubernetes

`kubectl apply -f deployment.yml`{{execute}}


Check deployments status

`kubectl get deployments`{{execute}}
`
NAME   READY   UP-TO-DATE   AVAILABLE   AGE
demo   1/1     1            1           42s
`

Execute command below to create Service

`kubectl apply -f service.yml`{{execute}}


Check pods status

`kubectl get pods`{{execute}}
`
NAME                   READY     STATUS    RESTARTS   AGE
demo-7d57479745-2c7gp   1/1     Running   0          66s
`

Check services status

`kubectl get svc`{{execute}}
`
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
demo         NodePort    10.101.43.190  <none>        9999:30000/TCP   22s
`

[Access Application](https://[[HOST_SUBDOMAIN]]-30000-[[KATACODA_HOST]].environments.katacoda.com/)



