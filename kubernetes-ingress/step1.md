## Prerequisites
In order to deploy Ingress
**Ingress controller** is necessary

We might need to deploy an Ingress controller such as **ingress-nginx**

For more information about **ingress-nginx**

refer to [helm:ingress-nginx](https://github.com/kubernetes/ingress-nginx/blob/master/docs/deploy/index.md#using-helm)

---
## Deploy ingress-nginx
Create deployment ingress-nginx with helm via **command** below
`helm install stable/nginx-ingress --name ingress-nginx`{{execute}}

Check ingress-nginx status
`helm ls`{{execute}}