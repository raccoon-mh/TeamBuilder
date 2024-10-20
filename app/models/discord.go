package models

type DiscordUser struct {
	ID               string  `json:"id"`
	Username         string  `json:"username"`
	Avatar           string  `json:"avatar"`
	Discriminator    string  `json:"discriminator"`
	PublicFlags      int     `json:"public_flags"`
	Flags            int     `json:"flags"`
	Banner           string  `json:"banner"`
	AccentColor      int     `json:"accent_color"`
	GlobalName       string  `json:"global_name"`
	AvatarDecoration *string `json:"avatar_decoration_data"`
	BannerColor      string  `json:"banner_color"`
	Clan             *string `json:"clan"`
	MfaEnabled       bool    `json:"mfa_enabled"`
	Locale           string  `json:"locale"`
	PremiumType      int     `json:"premium_type"`
	Email            string  `json:"email"`
	Verified         bool    `json:"verified"`
}
