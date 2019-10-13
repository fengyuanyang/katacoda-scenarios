Before we start, we need to know which image we are going to run.

There are two ways to search for images results.

1. Go to website [https://hub.docker.com/](https://hub.docker.com/)
and search for images then

2. Run `docker search nginx`{{execute}}


Result should be like it

| NAME  | DESCRIPTION              | STARS | OFFICIAL  | AUTOMATED |
|-------|--------------------------|-------|-----------|-----------|
| nginx | Official build of Nginx. | 12049 | [OK]      |           |

#### Note
If the name of the images not starts with XXX/, then it's a offical one.
Other images uploaded by users ALWAYS start with user name.

Such as `fengyuanyang/node-chrome-debug-video-start`
