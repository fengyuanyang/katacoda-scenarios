## Run Test and Check Video

Run your Test case by using RemoteWebDriver with URL below
`http://[[HOST_SUBDOMAIN]]-4444-[[KATACODA_HOST]].environments.katacoda.com/wd/hub`

A Sample Test code

	DesiredCapabilities capabilities = DesiredCapabilities.chrome();
	capabilities.setBrowserName("chrome");
	capabilities.setPlatform(Platform.LINUX);
	RemoteWebDriver driver = new RemoteWebDriver(new URL("http://[[HOST_SUBDOMAIN]]-4444-[[KATACODA_HOST]].environments.katacoda.com/wd/hub"), capabilities);

	driver.get("http://google.com");
    driver.findElement(By.name("q")).sendKeys("Selenium Grid Tutorial");
    driver.findElement(By.name("q")).sendKeys(Keys.ENTER);

    System.out.println(driver.getSessionId());
    driver.quit();

### get test video
Remember to copy your SessionId, your video can be fetched with that sessionID via URL below
	
`https://[[HOST_SUBDOMAIN]]-4444-[[KATACODA_HOST]].environments.katacoda.com/grid/admin/HubVideoDownloadServlet/?sessionId=[sessionId]`

Note:
video will be saved in Hub/Node container **tmp** folder by default
