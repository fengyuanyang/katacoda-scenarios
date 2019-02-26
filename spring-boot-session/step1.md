Run redis container via **command** below
`docker run --name some-redis -p 6379:6379 -d redis`{{execute}}

Enter redis container via **command** below
`docker run -it --link some-redis:redis --rm redis redis-cli -h redis -p 6379`{{execute}}

Run spring redis via **command** below
`./gradlew :spring-session-sample-boot-redis:bootRun`{{execute}}


Visit website via web below
https://[[HOST_SUBDOMAIN]]-8080-[[KATACODA_HOST]].environments.katacoda.com/
