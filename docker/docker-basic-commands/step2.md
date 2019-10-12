In previous step, we have found the image name we would like to use.
`nginx`

In order to run images, need to pull it first via command 
`docker pull nginx`{{execute}}

## Run container and enter
`docker run -p 8080:8080 -it bre bash`{{execute}}

## Run tomcat on port 8080
`bash /opt/tomcat/bin/startup.sh`{{execute}}

## Reach tomcat from browser
https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com


