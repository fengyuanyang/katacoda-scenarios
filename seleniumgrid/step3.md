## Start a Selenium Grid Node

VPN is necessary for a node to record Video, 

VPN will start by executing command below - [entry_point.sh] will be executed
`docker run -p 5555:5555 --name node -d fengyuanyang/node-chrome-debug-video`{{execute HOST1}}

List Selenium hub ip address by container name - for node to connect to
`docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' hub`{{execute HOST1}}

Start selenium node by executing command below
`docker exec -d node java -cp selenium-video-node-2.8.jar:selenium-server-standalone-3.9.1.jar:selenium-remote-driver-3.9.1.jar:httpcore-4.4.11.jar:httpclient-4.5.8.jar:commons-io-2.6.jar org.openqa.grid.selenium.GridLauncherV3 -servlets com.aimmac23.node.servlet.VideoRecordingControlServlet -proxy com.aimmac23.hub.proxy.VideoProxy -role node -browser "browserName=chrome" -hub http://172.18.0.2:4444/grid/register`{{execute HOST1}}