## Get Image from Registry

About **vidwo node Project**
Project refer to [selenium-video-node](https://github.com/aimmac23/selenium-video-node)

This Jar(selenium-video-node) provides serveral servlets and Proxy
1. VideoRecordingControlServlet - used in Node
	* Start Recording
	* Stop Recording
	* Reset Recording
	* Download Video

2. HubVideoDownloadServlet - used in Hub
	* Download video with key "webremotedriver sessionId"
	`http://「ip」:4444/grid/admin/HubVideoDownloadServlet/?sessionId=[sessionId]`

3. VideoProxy - used in Hub
	* constructor - send Reset Recording Request
	* beforeSession - send Start Recording Request
	* afterSession - send Stop Recording request and download video from Node with specific sessionId key

Parameter [video.path] can be used to set downloaded video path

usage
`-Dvideo.path=<directory>`

### Pull Image - selenium/node-chrome-debug - from Docker Registry 
`docker pull fengyuanyang/node-chrome-debug-video`{{execute HOST1}}
