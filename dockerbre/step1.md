## Build Image from Dockerfile

Image will be pulled from docker registry.

For more information about the parent image

refer [jdk6-mvn3](https://hub.docker.com/r/fengyuanyang/jdk6-mvn3/dockerfile)

This image includes
1. JDK 1.6
2. MAVEN 3.2.5
3. TOMCAT 7

Build docker image via **command** below

`docker build -t bre .`{{execute}}

After building successfully, check if image **bre** has been created

`docker images`{{execute}}

Result should be like it

| REPOSITORY | IMAGE ID     | CREATED        | SIZE  |
|------------|--------------|----------------|-------|
| bre        | XXXXXXXXXXXX | 21 seconds ago | 823MB |
