`kubectl apply -f kubernetes-dashboard.yaml`{{execute}}


#### Access Dashboard after deploying Dashboard 
[Access Dashboard](https://[[HOST_SUBDOMAIN]]-8001-[[KATACODA_HOST]].environments.katacoda.com/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/overview?namespace=default)


#### Check deployments status
`kubectl get deployments -n kube-system`{{execute}}
`
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
bre       1         1         1            1           5m
`

#### Check services status
`kubectl get svc -n kube-system`{{execute}}
`
NAME         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
bre          NodePort    10.98.14.144   <none>        8090:30001/TCP   6m
kubernetes   ClusterIP   10.96.0.1      <none>        443/TCP          7m
`
