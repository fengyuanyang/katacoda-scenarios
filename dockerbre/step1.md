First step.

## Task

Build image from Dockerfile
Image will be pulled from docker registry.
This image includes
1. JDK 1.6
2. MAVEN 3.2.5
3. TOMCAT 7

Build docker image via **commend** below

`docker build -t bre .`{{execute}}

After building successfully, check if image **bre** has been created

`docker images`{{execute}}

Result should be like it

| REPOSITORY | IMAGE ID     | CREATED        | SIZE  |
|------------|--------------|----------------|-------|
| bre        | 79bda7d73bae | 21 seconds ago | 823MB |
