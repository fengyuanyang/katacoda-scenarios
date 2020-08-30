
kubernetes is intializing, wait for couple minutes and check cluster status with command below

`kubectl apply -f deployment_server_pub.yaml`{{execute}}

Check deployments status

`kubectl get deployments`{{execute}}

Check services status

`kubectl get svc`{{execute}}

[Access Dashboard](https://[[HOST_SUBDOMAIN]]-30001-[[KATACODA_HOST]].environments.katacoda.com/mqtt?message=abc