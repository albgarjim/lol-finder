package output

type TestObj struct {
	ID      string  `bson:"_id"`
	Title   string  `bson:"title"`
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
	ID        string      `bson:"_id"`
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

		} `bson:"buddyRoles"`

		ChatUsage ChatUsage `bson:"chatUsage"`

	}  `bson:"lookingFor"`
}

type RankObj struct {
	ID      string  `bson:"_id" json:"_id"`
	Label   string  `json:"label"`
	Img     string  `json:"img"`
}

type ServerObj struct {
	ID      string  `bson:"_id" json:"_id"`
	Label   string  `json:"label"`
	Name    string  `json:"name"`
	Img     string  `json:"img"`
}

type LadderObj struct {
	ID      string  `bson:"_id" json:"_id"`
	Label   string  `json:"label"`
	Img     string  `json:"img"`
}

type RoleObj struct {
	ID      string  `bson:"_id" json:"_id"`
	Label   string  `json:"label"`
	Img     string  `json:"img"`
}

type DuoObjComplete struct {
	ID        string      `json:"_id"`
	PlayerData struct {
		PlayerName     string       `json:"playerName"`
		RolesPlayed struct {
			Top struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"top"`

			Jungle struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"jungle"`

			Mid struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"mid"`

			Adc struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"adc"`

			Support struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"support"`

		} `json:"rolesPlayed"`

		Server      ServerObj    `json:"server"`
		SoloQRank   RankObj      `json:"soloQRank"`
		OpggData    OpggData     `json:"opggData"`

	} `json:"playerData"`

	LookingFor struct {
		LadderMode   LadderObj  `json:"ladderMode"`
		BuddyRank struct {
			MinRank  RankObj  `json:"minRank"`
			MaxRank  RankObj  `json:"maxRank"`
		} `json:"buddyRank"`

		BuddyRoles struct {
			Top struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"top"`

			Jungle struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"jungle"`

			Mid struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"mid"`

			Adc struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"adc"`

			Support struct {
				RoleInfo  RoleObj  `json:"roleInfo"`
				Selected  bool    `json:"selected"`
			} `json:"support"`

		} `json:"buddyRoles"`

		ChatUsage ChatUsage `json:"chatUsage"`

	}  `json:"lookingFor"`
}
