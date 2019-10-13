Now, we have downloaded the image we need.
Check what images we have for now 

In order to run images, need to pull it first via command 
`docker images`{{execute}}

`
Using default tag: latest
latest: Pulling from library/nginx
b8f262c62ec6: Pull complete
e9218e8f93b1: Pull complete
7acba7289aa3: Pull complete
`

## Run container and enter
`docker run -p 8080:8080 -it bre bash`{{execute}}

## Run tomcat on port 8080
`bash /opt/tomcat/bin/startup.sh`{{execute}}

## Reach tomcat from browser
https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com


