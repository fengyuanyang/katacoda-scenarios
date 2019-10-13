We have accessed our nginx app successfully, time to delete it
### docker rm
Remove one or more containers

|Name          | shorthand |Description                          |
|--------------|-----------|-------------------------------------|
|--force       | -f        |Force the removal of a running container (uses SIGKILL)|

Since our container is still running, we need to stop it before removing it

First, stop container
`docker stop nginx`{{execute}}

Second, delete container
`docker rm nginx`{{execute}}

or delete container with -f (force).
`docker rm -f nginx`

Now, check container status again
`docker ps -a`{{execute}}

Container with name `nginx` has been removed.

### docker rmi
Remove one or more images

If we don't need the image any more, we could delete it
`docker rmi nginx`{{execute}}