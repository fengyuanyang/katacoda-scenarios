We have accessed to our nginx app, time to delete it
### docker rm
execute *docker rm* 

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--force       | -f        |Force the removal of a running container (uses SIGKILL)|

Since our container is still running, we need to stop it before deleting it

First, stop container
`docker stop nginx`{{execute}}
Second, delete container
`docker rm nginx`{{execute}}

or delete container with -f (force).
`docker rm -f nginx`

Now, check container status again
`docker ps -a`{{execute}}

Container with name `nginx` has been removed.

### docker rm
Remove our image `nginx`
If we don't need the image any more, we could delete it
`docker rmi nginx`{{execute}}