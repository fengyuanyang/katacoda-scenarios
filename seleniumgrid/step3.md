## Start a Selenium Grid Node

VPN is necessary for a node to record Video, 

With command
`docker run -d selenium/node-chrome-debug`
[entry_point.sh] will be executed
VPN will start and Ndoe will listen at port 5555

Map port Host port 5555 to Container port 5555 by executing command below
`docker run -p 5555:5555  -n node -d fengyuanyang/node-chrome-debug-video`{{execute HOST2}}

Since node has been started by script [entry_point.sh], we need to stop it and rerun it with specific parameters
* -servlets
* -proxy

`docker exec -it node bash`{{execute HOST2}}

`java -cp selenium-video-node-2.8.jar:selenium-server-standalone-3.9.1.jar:selenium-remote-driver-3.9.1.jar:httpcore-4.4.11.jar:httpclient-4.5.8.jar:commons-io-2.6.jar org.openqa.grid.selenium.GridLauncherV3 -servlets com.aimmac23.node.servlet.VideoRecordingControlServlet -proxy com.aimmac23.hub.proxy.VideoProxy -role node -browser "browserName=chrome"`{{execute HOST2}}


`ps -ef`
