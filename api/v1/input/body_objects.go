package input

type TestObj struct {
	Title       string             `bson:"title"`
}

type ContactMail struct {
	Sender string `bson: "sender"`
	Subject string `bson: "subject"`
	Content string `bson: "content"`
	Name string `bson: "name"`
}

type OpggData struct {
	OpggURL  string  `bson:"opggURL"`
	Visible  bool    `bson:"visible"`
}

type ChatUsage struct {
	DiscordURL   string   `bson:"discordURL"`
	Usage     bool     `bson:"chatUsage"`
}

type DuoObj struct {
	PlayerData struct {
		PlayerName     string       `bson:"playerName"`
		RolesPlayed struct {
			Top struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"top"`

			Jungle struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"jungle"`

			Mid struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"mid"`

			Adc struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"adc"`

			Support struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"support"`

			Fill struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"fill"`

		} `bson:"rolesPlayed"`

		Server      string       `bson:"server"`
		SoloQRank   string       `bson:"soloQRank"`
		OpggData    OpggData     `bson:"opggData"`

	} `bson:"playerData"`

	LookingFor struct {
		LadderMode   string  `bson:"ladderMode"`
		BuddyRank struct {
			MinRank  string  `bson:"minRank"`
			MaxRank  string  `bson:"maxRank"`
		} `bson:"buddyRank"`

		BuddyRoles struct {
			Top struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"top"`

			Jungle struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"jungle"`

			Mid struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"mid"`

			Adc struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"adc"`

			Support struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"support"`

			Fill struct {
				RoleInfo  string  `bson:"roleInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"fill"`

		} `bson:"buddyRoles"`

		ChatUsage ChatUsage `bson:"chatUsage"`

	}  `bson:"lookingFor"`
}