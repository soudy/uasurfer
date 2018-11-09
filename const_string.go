package uasurfer

var deviceTypeStrings = map[DeviceType]string{
	DeviceUnknown:  "unknown",
	DeviceComputer: "desktop",
	DeviceTablet:   "tablet",
	DevicePhone:    "mobile",
	DeviceConsole:  "console",
	DeviceWearable: "wearable",
	DeviceTV:       "tv",
}

func (dt DeviceType) String() string {
	if deviceType, ok := deviceTypeStrings[dt]; ok {
		return deviceType
	}

	return "unknown"
}

var browserTypeStrings = map[BrowserName]string{
	BrowserUnknown:       "Unknown",
	BrowserChrome:        "Chrome",
	BrowserIE:            "IE",
	BrowserSafari:        "Safari",
	BrowserFirefox:       "Firefox",
	BrowserOpera:         "Opera",
	BrowserBlackberry:    "Blackberry",
	BrowserUCBrowser:     "UCBrowser",
	BrowserSilk:          "Silk",
	BrowserNokia:         "Nokia",
	BrowserNetFront:      "NetFront",
	BrowserQQ:            "QQ",
	BrowserMaxthon:       "Maxthon",
	BrowserSogouExplorer: "SogouExplorer",
	BrowserSpotify:       "Spotify",
	BrowserWebView:       "WebView",
	BrowserSamsung:       "SamsungBrowser",
	BrowserFacebook:      "Facebook",
	BrowserMiui:          "MiuiBrowser",
	BrowserAppleBot:      "AppleBot",
	BrowserBaiduBot:      "BaiduBot",
	BrowserBingBot:       "BingBot",
	BrowserDuckDuckGoBot: "DuckDuckGoBot",
	BrowserFacebookBot:   "FacebookBot",
	BrowserGoogleBot:     "GoogleBot",
	BrowserLinkedInBot:   "LinkedInBot",
	BrowserMsnBot:        "MsnBot",
	BrowserPingdomBot:    "PingdomBot",
	BrowserTwitterBot:    "TwitterBot",
	BrowserYandexBot:     "YandexBot",
}

func (bn BrowserName) String() string {
	if browserType, ok := browserTypeStrings[bn]; ok {
		return browserType
	}

	return "Unknown"
}

var osTypeStrings = map[OSName]string{
	OSUnknown:      "Unknown",
	OSWindowsPhone: "WindowsPhoneOS",
	OSWindows:      "Windows",
	OSMacOSX:       "OS X",
	OSiOS:          "iOS",
	OSAndroid:      "Android",
	OSBlackberry:   "Blackberry",
	OSChromeOS:     "ChromeOS",
	OSKindle:       "Kindle",
	OSWebOS:        "webOS",
	OSLinux:        "Linux",
	OSPlaystation:  "Playstation",
	OSXbox:         "Xbox",
	OSNintendo:     "Nintendo",
	OSBot:          "Bot",
}

func (os OSName) String() string {
	if osType, ok := osTypeStrings[os]; ok {
		return osType
	}

	return "Unknown"
}

var platformStrings = map[Platform]string{
	PlatformUnknown:      "Unknown",
	PlatformWindows:      "Windows",
	PlatformMac:          "Mac",
	PlatformLinux:        "Linux",
	PlatformiPad:         "iPad",
	PlatformiPhone:       "iPhone",
	PlatformiPod:         "iPod",
	PlatformBlackberry:   "Blackberry",
	PlatformWindowsPhone: "WindowsPhone",
	PlatformPlaystation:  "Playstation",
	PlatformXbox:         "Xbox",
	PlatformNintendo:     "Nintendo",
	PlatformBot:          "Bot",
}

func (p Platform) String() string {
	if platform, ok := platformStrings[p]; ok {
		return platform
	}

	return "Unknown"
}
