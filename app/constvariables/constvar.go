package constvariables

var (
	// Terminal
	TeminalRedStart = "\033[31m"
	TeminalEnd      = "\033[0m"

	// version
	startBannerVersion = "v0.0.1"

	// logo

	logo = "  ______                     ____        _ __    __         \n" +
		" /_  __/__  ____ _____ ___  / __ )__  __(_) /___/ /__  _____\n" +
		"  / / / _ \\/ __ `/ __ `__ \\/ __  / / / / / / __  / _ \\/ ___/\n" +
		" / / /  __/ /_/ / / / / / / /_/ / /_/ / / / /_/ /  __/ /    \n" +
		"/_/  \\___/\\__/_/ /_/ /_/_/_____/\\____/_/_/\\____/\\___/_/     "

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
