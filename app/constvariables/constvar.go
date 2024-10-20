package constvariables

var (
	// Terminal
	TeminalRedStart = "\033[31m"
	TeminalEnd      = "\033[0m"

	// version
	startBannerVersion = "v0.0.1"

	// logo
	logo = "    __          __  ______                     __  ___      __            \n" +
		"   / /   ____  / / /_  __/__  ____ _____ ___  /  |/  /___ _/ /_____  _____\n" +
		"  / /   / __ \\/ /   / / / _ \\/ __ `/ __ `__ \\/ /|_/ / __ `/ //_/ _ \\/ ___/\n" +
		" / /___/ /_/ / /___/ / /  __/ /_/ / / / / / / /  / / /_/ / ,< /  __/ /    \n" +
		"/_____/\\____/_____/_/  \\___/\\__,_/_/ /_/ /_/_/  /_/\\__,_/_/|_|\\___/_/     "
	startBannerSubTitle = "Github = https://github.com/raccoon-mh/LoLTeamMaker , Author = raccoon-mh"

	// banner
	StartBanner = logo +
		TeminalRedStart +
		startBannerVersion +
		TeminalEnd + "\n" +
		startBannerSubTitle + "\n\n"

	// Discord
	DiscordAuthorizeEndpoint    = "https://discord.com/oauth2/authorize"
	DiscordTokenEndpoint        = "https://discord.com/api/oauth2/token"
	DiscordTokenReovokeEndpoint = "https://discord.com/api/oauth2/token/revoke"
	DiscordIdentifyEndpoint     = "https://discord.com/api/users/@me"
)
