It's time to run our first app-nginx
execute *docker run* 

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--name        |           |Assign a name to the container       |
|--interactive | -i		   |Keep STDIN open even if not attached |
|--tty         | -t		   |Allocate a pseudo-TTY                |
|--rm          |   		   |Automatically remove the container when it exits|


With parameter `name`, we can delete it much easier
`docker run --name nginx nginx`{{execute}}

Check what containers are running for now
*docker ps*

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--all         |-a         |Show all containers (default shows just running)|

`docker ps`{{execute}}



## Run container and enter
`docker run -p 8080:8080 -it bre bash`{{execute}}

## Run tomcat on port 8080
`bash /opt/tomcat/bin/startup.sh`{{execute}}

## Reach tomcat from browser
https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com


