### docker images
Now, we have downloaded the image we need **nginx**.

Check what images we have for now 

`docker images`{{execute}}

|REPOSITORY| TAG   |IMAGE ID     |CREATED      | SIZE  |
|----------|-------|-------------|-------------|-------|
|nginx     | latest|f949e7d76d63 | 2 weeks ago | 126MB |

### docker history
Show the history of an image

`docker history nginx`{{execute}}

### docker inspect
Return low-level information on Docker objects

`docker inspect nginx`{{execute}}