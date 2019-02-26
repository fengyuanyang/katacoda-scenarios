Run redis container via **command** below
`docker run --name some-redis -p 6379:6379 -d redis`


Run spring redis via **command** below
`./gradlew :spring-session-sample-javaconfig-redis:tomcatRun`{{execute}}


Visit website via web below
https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com/