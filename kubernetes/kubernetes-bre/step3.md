#### Deploy Dashboard for kubernetes
`kubectl create -f kubernetes-dashboard.yaml`{{execute}}

#### To access Dashboard, you must create a secure channel to your Kubernetes cluster. Run the following command
`kubectl proxy --address='0.0.0.0' --port=8001 --accept-hosts='^*$' &`{{execute}}

#### Access Dashboard after deploying Dashboard 
[Access Dashboard](https://[[HOST_SUBDOMAIN]]-8001-[[KATACODA_HOST]].environments.katacoda.com/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/overview?namespace=default)


Deploying status and logs can be whatched via Dashboard
Click 'SKIP/略過' when you see authentication phase