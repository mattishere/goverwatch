package url

var (
	// Role icons
	TankIcon      = "https://static.playoverwatch.com/img/pages/career/icons/role/tank-f64702b684.svg#icon"
	DPSIcon       = "https://static.playoverwatch.com/img/pages/career/icons/role/offense-ab1756f419.svg#icon"
	SupportIcon   = "https://static.playoverwatch.com/img/pages/career/icons/role/support-0258e13d85.svg#icon"
	OpenQueueIcon = "https://static.playoverwatch.com/img/pages/career/icons/role/open-163b3b8ddc.svg#icon"

	// A map of the role icons to their respective role.
	RoleURLS = map[string]string{
		TankIcon:      "tank",
		DPSIcon:       "dps",
		SupportIcon:   "support",
		OpenQueueIcon: "open",
	}

	// The base URL for the official Overwatch 2 stats website.
	BaseURL = "https://overwatch.blizzard.com/en-us/career/"
)
