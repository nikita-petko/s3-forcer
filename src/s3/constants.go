package s3

// Base Paths

const (
	rbxCdnProviderBaseUrl         string = "https://setup.rbxcdn.com/"
	s3BaseUrl                     string = "https://s3.amazonaws.com/"
	s3BasePath                    string = s3BaseUrl + "setup.roblox.com/"
	rbxCdnProviderChannelBasePath string = rbxCdnProviderBaseUrl + "channel/%s/"
	s3PerChannelBasePath          string = s3BasePath + "channel/%s/"
)

const (
	deployHistoryFile                     string = "DeployHistory.txt"
	rbxCdnProviderPerChannelDeployHistory string = rbxCdnProviderChannelBasePath + deployHistoryFile
	s3PerChannelDeployHistory             string = s3PerChannelBasePath + deployHistoryFile
)
