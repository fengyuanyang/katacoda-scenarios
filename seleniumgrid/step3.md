## Start a Selenium Grid Node

VPN is necessary for a node to record Video, 

With command
`docker run -d selenium/node-chrome-debug`
[entry_point.sh] will be executed
VPN will start and Ndoe will listen at port 5555

Map port Host port 5555 to Container port 5555 by executing command below
`docker run -v /root:/  -p 5555:5555 -d selenium/node-chrome-debug -n node`{{execute HOST2}}

Since node has been started by script [entry_point.sh], we need to stop it and rerun it with specific parameters
* -servlets
* -proxy

`docker exec -it node bash`{{execute HOST2}}

`ps -ef`
