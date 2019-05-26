## Start a Selenium Grid Hub
Hub can be run and listen at port 4444 by command below
`docker run -p 4444:4444 --name hub -d fengyuanyang/node-chrome-debug-video java -cp selenium-video-node-2.8.jar:selenium-server-standalone-3.9.1.jar:selenium-remote-driver-3.9.1.jar:httpcore-4.4.11.jar:httpclient-4.5.8.jar:commons-io-2.6.jar org.openqa.grid.selenium.GridLauncherV3 -servlets com.aimmac23.hub.servlet.HubVideoDownloadServlet -role hub`{{execute HOST1}}

Grid console can be accessed by url below
`http://「ip」:4444/grid/console`

[Grid status](https://[[HOST_SUBDOMAIN]]-4444-[[KATACODA_HOST]].environments.katacoda.com/grid/console)