## Run Test and Check Video

Run your Test case by using RemoteWebDriver with code below

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
	https://[[HOST_SUBDOMAIN]]-4444-[[KATACODA_HOST]].environments.katacoda.com/grid/admin/HubVideoDownloadServlet/?sessionId=(sessionID)
