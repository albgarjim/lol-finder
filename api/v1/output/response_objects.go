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
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"top"`

			Jungle struct {
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"jungle"`

			Mid struct {
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"mid"`

			Adc struct {
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"adc"`

			Support struct {
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"support"`

		} `bson:"rolesPlayed"`

		Server      string       `bson:"server"`
		SoloQRank   string       `bson:"soloQRank"`
		OpggData    OpggData     `bson:"opggData"`

	} `bson:"playerData"`

	LookingFor struct {
		GameMode   string  `bson:"gameMode"`
		BuddyRank struct {
			MinRank  string  `bson:"minRank"`
			MaxRank  string  `bson:"maxRank"`
		} `bson:"buddyRank"`

		BuddyRoles struct {
			Top struct {
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"top"`

			Jungle struct {
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"jungle"`

			Mid struct {
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"mid"`

			Adc struct {
				LaneInfo  string  `bson:"laneInfo"`
				Selected  bool    `bson:"selected"`
			} `bson:"adc"`

			Support struct {
				LaneInfo  string  `bson:"laneInfo"`
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
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"top"`

			Jungle struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"jungle"`

			Mid struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"mid"`

			Adc struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"adc"`

			Support struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"support"`

		} `json:"rolesPlayed"`

		Server      ServerObj    `json:"server"`
		SoloQRank   RankObj      `json:"soloQRank"`
		OpggData    OpggData     `json:"opggData"`

	} `json:"playerData"`

	LookingFor struct {
		GameMode   LadderObj  `json:"gameMode"`
		BuddyRank struct {
			MinRank  RankObj  `json:"minRank"`
			MaxRank  RankObj  `json:"maxRank"`
		} `json:"buddyRank"`

		BuddyRoles struct {
			Top struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"top"`

			Jungle struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"jungle"`

			Mid struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"mid"`

			Adc struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"adc"`

			Support struct {
				LaneInfo  RoleObj  `json:"laneInfo"`
				Selected  bool    `json:"selected"`
			} `json:"support"`

		} `json:"buddyRoles"`

		ChatUsage ChatUsage `json:"chatUsage"`

	}  `json:"lookingFor"`
}
