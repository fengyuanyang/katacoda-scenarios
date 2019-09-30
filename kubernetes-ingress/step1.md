## Prerequisites
In order to deploy Ingress
**Ingress controller** is necessary

We might need to deploy an Ingress controller such as **ingress-nginx**
Note. Here we are using helm to deploy ingress-nginx

For more information about **ingress-nginx**

refer to [ingress-nginx](https://github.com/kubernetes/ingress-nginx/blob/master/docs/deploy/index.md#using-helm)

---
## Deploy ingress-nginx
Before starting using helm, we need to install **tiller** first.
More info about [helm init](https://helm.sh/docs/install/)
`helm init`{{execute}}

Create deployment ingress-nginx with helm via **command** below
`helm install stable/nginx-ingress --name ingress-nginx`{{execute}}

Check ingress-nginx status
`helm ls`{{execute}}

One Chart named ingress-nginx should be listed as below
`
NAME            REVISION        UPDATED                         STATUS          CHART                 APP VERSION      NAMESPACE
ingress-nginx   1               Mon Sep 30 13:57:59 2019        DEPLOYED        nginx-ingress-1.22.1  0.25.1           default
`