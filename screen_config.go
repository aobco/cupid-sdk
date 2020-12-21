package config

type ScreenConfig struct {
	Origin           string `json:"origin" yaml:"origin1"`
	PlatformId       string `json:"platform_id" yaml:"platformId"`
	PlatformType     string `json:"platform_type" yaml:"platformType"`
	Url              string `json:"url" yaml:"origin2"`
	StunServer       string `json:"stun_server" yaml:"stunServer"`
	TurnServer       string `json:"turn_server" yaml:"turnServer"`
	TurnUser         string `json:"turn_user" yaml:"turnUser"`
	TurnPwd          string `json:"turn_pwd" yaml:"turnPwd"`
	RoomId           string `json:"room_id" yaml:"roomId"`
	ScreenshotServer string `json:"screenshot_server" yaml:"screenshotServer"`
	ApiServer        string `json:"api_server" yaml:"apiServer"`
	ServerIp         string `json:"server_ip" yaml:"serverIp"`
	PrivateKey       string `json:"private_key" yaml:"privateKey"`
	BondEth       	 string `json:"bond_eth" yaml:"bond_eth"`
}
