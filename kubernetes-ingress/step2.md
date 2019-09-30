
#### Execute command below to create Service
`kubectl create -f service.yml`{{execute}}

#### Check deployments status
`kubectl get deployments`{{execute}}

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
[Access Application](https://[[HOST_SUBDOMAIN]]-30001-[[KATACODA_HOST]].environments.katacoda.com/)



