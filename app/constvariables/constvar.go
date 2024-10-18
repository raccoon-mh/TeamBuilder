package constvariables

var (
	startBannerVersion = "v0.0.1"

	logo = "    __          __  ______                     __  ___      __            \n" +
		"   / /   ____  / / /_  __/__  ____ _____ ___  /  |/  /___ _/ /_____  _____\n" +
		"  / /   / __ \\/ /   / / / _ \\/ __ `/ __ `__ \\/ /|_/ / __ `/ //_/ _ \\/ ___/\n" +
		" / /___/ /_/ / /___/ / /  __/ /_/ / / / / / / /  / / /_/ / ,< /  __/ /    \n" +
		"/_____/\\____/_____/_/  \\___/\\__,_/_/ /_/ /_/_/  /_/\\__,_/_/|_|\\___/_/     "
	startBannerSubTitle = "Github = https://github.com/raccoon-mh/LoLTeamMaker , Author = raccoon-mh"

	TeminalRedStart = "\033[31m"
	TeminalEnd      = "\033[0m"

	StartBanner = logo +
		TeminalRedStart +
		startBannerVersion +
		TeminalEnd + "\n" +
		startBannerSubTitle + "\n\n"
)
