package s3

// Base Paths

const (
	rbxCdnProviderBaseUrl         string = "https://setup.rbxcdn.com/"
	s3BaseUrl                     string = "https://s3.amazonaws.com/"
	s3BasePath                    string = s3BaseUrl + "setup.roblox.com/"
	rbxCdnProviderChannelBasePath string = rbxCdnProviderBaseUrl + "channel/%s/"
	s3PerChannelBasePath          string = s3BasePath + "channel/%s/"
)

// Version

const (
	versionFile                     string = "version"
	rbxCdnProviderPerChannelVersion string = rbxCdnProviderChannelBasePath + versionFile
	s3PerChannelVersion             string = s3PerChannelBasePath + versionFile
)
