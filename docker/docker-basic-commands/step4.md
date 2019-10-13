It's time to run our first app-nginx
### docker run
execute *docker run* 

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--name        |           |Assign a name to the container       |
|--interactive | -i		   |Keep STDIN open even if not attached |
|--tty         | -t		   |Allocate a pseudo-TTY                |
|--rm          |   		   |Automatically remove the container when it exits|
|--publish     | -p        |Publish a container’s port(s) to the host|


With parameter 

`name` - we can delete it much easier later
`p` - we can access nginx via port 80 
`it` - keep STDIN open and allocate a TTY
`docker run -it -p 80:80 --name nginx nginx`{{execute}}

### docker ps
Check what containers are running for now
*docker ps*

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--all         |-a         |Show all containers (default shows just running)|

`docker ps`{{execute}}


## Reach tomcat from browser
https://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com


