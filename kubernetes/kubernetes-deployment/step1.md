
kubernetes is intializing, wait for couple minutes and check cluster status with command below

`kubectl cluster-info`{{execute}}

create dashboard so we can monitor status

`kubectl apply -f kubernetes-dashboard.yaml`{{execute}}

Check deployments status

`kubectl get deployments -n kube-system`{{execute}}
`
NAME                      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
kubernetes-dashboard      0/1     1            0           2s
`

Check services status

`kubectl get svc -n kube-system`{{execute}}
`
NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
kubernetes-dashboard   ClusterIP   10.97.20.6   <none>        443/TCP                  44s
`

To access Dashboard, you must create a secure channel to your Kubernetes cluster. Run the following command

`kubectl proxy --address='0.0.0.0' --port=8001 --accept-hosts='^*$' &`{{execute}}

[Access Dashboard](https://[[HOST_SUBDOMAIN]]-8001-[[KATACODA_HOST]].environments.katacoda.com/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/overview?namespace=default)