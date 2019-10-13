It's time to run our first app-nginx
### docker run

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--name        |           |Assign a name to the container       |
|--interactive | -i		   |Keep STDIN open even if not attached |
|--tty         | -t		   |Allocate a pseudo-TTY                |
|--rm          |   		   |Automatically remove the container when it exits|
|--publish     | -p        |Publish a container’s port(s) to the host|
|--detach      | -d        |Run container in background and print container ID|


execute *docker run* with parameter 

`name` - we can delete it much easier later

`p` - we can access nginx via port 80 

`it` - keep STDIN open and allocate a TTY

`docker run -d -p 80:80 --name nginx nginx`{{execute}}

### docker ps
Check what containers are running for now

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--all         |-a         |Show all containers (default shows just running)|

`docker ps`{{execute}}

### docker exec
Run a command in a running container

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--name        |           |Assign a name to the container       |
|--interactive | -i		   |Keep STDIN open even if not attached |
|--tty         | -t		   |Allocate a pseudo-TTY                |

execute *docker exec* with parameter 

`it` - in a detach mode

`docker exec -it nginx bash`{{execute}}

Now, we can execute any command inside the container.
Such as, take a look on nginx properties
`cat /etc/nginx/nginx.conf`{{execute}}


## Reach negix from browser
https://[[HOST_SUBDOMAIN]]-80-[[KATACODA_HOST]].environments.katacoda.com


