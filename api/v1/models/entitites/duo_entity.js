
const duo = {
    playerData: {
        playerName: "monkaccino",
        rolesPlayed: {
            top: {
                selected: true,
                roleInfo: "id-role-top"
            },
            jungle: {
                selected: false,
                roleInfo: "id-role-jungle"
            },
            mid: {
                selected: false,
                roleInfo: "id-role-mid"
            },
            adc: {
                selected: false,
                roleInfo: "id-role-adc"
            },
            support: {
                selected: false,
                roleInfo: "id-role-support"
            }
        },
        server: "id-server-euw",
        soloQRank: "id-rank-gold-2",
        opggData: {
            opggURL: "https://discord-url.com",
            visible: false
        }
    },
    lookingFor: {
        ladderMode: "id-ladder-soloq",
        buddyRank: {
            minRank: "id-rank-gold-4",
            maxRank: "id-rank-gold-1",
        },
        buddyRoles: {
            top: {
                selected: false,
                roleInfo: "id-role-top"
            },
            jungle: {
                selected: false,
                roleInfo: "id-role-jungle"
            },
            mid: {
                selected: true,
                roleInfo: "id-role-mid"
            },
            adc: {
                selected: true,
                roleInfo: "id-role-adc"
            },
            support: {
                selected: true,
                roleInfo: "id-role-support"
            }
        },
        chatUsage: {
            discordURL: "https://discord-url.com",
            chatUsage: false
        }
    }
}

const ranksdb = [
    {
        _id: "id-rank-iron-4",
        label: "Iron 4",
        img: "http://img-rank-iron-4-url.com"
    },
    {
        _id: "id-rank-iron-3",
        label: "Iron 3",
        img: "http://img-rank-iron-3-url.com"
    },
    {
        _id: "id-rank-iron-2",
        label: "Iron 2",
        img: "http://img-rank-iron-2-url.com"
    },
    {
        _id: "id-rank-iron-1",
        label: "Iron 1",
        img: "http://img-rank-iron-1-url.com"
    },
    {
        _id: "id-rank-bronze-4",
        label: "Bronze 4",
        img: "http://img-rank-bronze-4-url.com"
    },
    {
        _id: "id-rank-bronze-3",
        label: "Bronze 3",
        img: "http://img-rank-bronze-3-url.com"
    },
    {
        _id: "id-rank-bronze-2",
        label: "Bronze 2",
        img: "http://img-rank-bronze-2-url.com"
    },
    {
        _id: "id-rank-bronze-1",
        label: "Bronze 1",
        img: "http://img-rank-bronze-1-url.com"
    },
    {
        _id: "id-rank-silver-4",
        label: "Silver 4",
        img: "http://img-rank-silver-4-url.com"
    },
    {
        _id: "id-rank-silver-3",
        label: "Silver 3",
        img: "http://img-rank-silver-3-url.com"
    },
    {
        _id: "id-rank-silver-2",
        label: "Silver 2",
        img: "http://img-rank-silver-2-url.com"
    },
    {
        _id: "id-rank-silver-1",
        label: "Silver 1",
        img: "http://img-rank-silver-1-url.com"
    },
    {
        _id: "id-rank-gold-4",
        label: "Gold 4",
        img: "http://img-rank-gold-4-url.com"
    },
    {
        _id: "id-rank-gold-3",
        label: "Gold 3",
        img: "http://img-rank-gold-3-url.com"
    },
    {
        _id: "id-rank-gold-2",
        label: "Gold 2",
        img: "http://img-rank-gold-2-url.com"
    },
    {
        _id: "id-rank-gold-1",
        label: "Gold 1",
        img: "http://img-rank-gold-1-url.com"
    },
    {
        _id: "id-rank-platinum-4",
        label: "Platinum 4",
        img: "http://img-rank-platinum-4-url.com"
    },
    {
        _id: "id-rank-platinum-3",
        label: "Platinum 3",
        img: "http://img-rank-platinum-3-url.com"
    },
    {
        _id: "id-rank-platinum-2",
        label: "Platinum 2",
        img: "http://img-rank-platinum-2-url.com"
    },
    {
        _id: "id-rank-platinum-1",
        label: "Platinum 1",
        img: "http://img-rank-platinum-1-url.com"
    },
    {
        _id: "id-rank-diamond-4",
        label: "Diamond 4",
        img: "http://img-rank-diamond-4-url.com"
    },
    {
        _id: "id-rank-diamond-3",
        label: "Diamond 3",
        img: "http://img-rank-diamond-3-url.com"
    },
    {
        _id: "id-rank-diamond-2",
        label: "Diamond 2",
        img: "http://img-rank-diamond-2-url.com"
    },
    {
        _id: "id-rank-diamond-1",
        label: "Diamond 1",
        img: "http://img-rank-diamond-1-url.com"
    },
    {
        _id: "id-rank-master",
        label: "Master",
        img: "http://img-rank-master-url.com"
    },
    {
        _id: "id-rank-grandmaster",
        label: "Grandmaster",
        img: "http://img-rank-grandmaster-url.com"
    },
    {
        _id: "id-rank-challenger",
        label: "Challenger",
        img: "http://img-rank-challenger-url.com"
    }
]

const serversDB = [
    {
        _id: "id-server-br",
        label: "BR",
        name: "Brazil",
        img: "http://server-br-url.com"
    },
    {
        _id: "id-server-eune",
        label: "EUNE",
        name: "Europe Nordic & East",
        img: "http://server-eune-url.com"
    },
    {
        _id: "id-server-euw",
        label: "EUW",
        name: "Europe West",
        img: "http://server-euw-url.com"
    },
    {
        _id: "id-server-jp",
        label: "JP",
        name: "Japan",
        img: "http://server-jp-url.com"
    },
    {
        _id: "id-server-lan",
        label: "LAN",
        name: "Latin America North",
        img: "http://server-lan-url.com"
    },
    {
        _id: "id-server-las",
        label: "LAS",
        name: "Latin America South",
        img: "http://server-las-url.com"
    },
    {
        _id: "id-server-na",
        label: "NA",
        name: "North America",
        img: "http://server-na-url.com"
    },
    {
        _id: "id-server-oce",
        label: "OCE",
        name: "Oceania",
        img: "http://server-oce-url.com"
    },
    {
        _id: "id-server-ru",
        label: "RU",
        name: "Russia",
        img: "http://server-ru-url.com"
    },
    {
        _id: "id-server-tr",
        label: "TR",
        name: "Turkey",
        img: "http://server-tr-url.com"
    },
    {
        _id: "id-server-kr",
        label: "KR",
        name: "Republic of Korea",
        img: "http://server-kr-url.com"
    },
]

const laddersDB = [
    {
        _id: "id-ladder-soloq",
        label: "Ranked Solo / Duo",
        img: "http://img-ladder-solo-url.com"
    },
    {
        _id: "id-ladder-flex",
        label: "Flex Queue",
        img: "http://img-ladder-flex-url.com"
    },
    {
        _id: "id-ladder-normals",
        label: "Normals",
        img: "http://img-ladder-normals-url.com"
    },
    {
        _id: "id-ladder-aram",
        label: "ARAM",
        img: "http://img-ladder-aram-url.com"
    },
    {
        _id: "id-ladder-custom",
        label: "Custom",
        img: "http://img-ladder-custom-url.com"
    },
]

const rolesDB = [
    {
        _id: "id-role-top",
        label: "Top",
        img: "http://img-top-role-url.com"
    },
    {
        _id: "id-role-jungle",
        label: "Jungle",
        img: "http://img-jungle-role-url.com"
    },
    {
        _id: "id-role-mid",
        label: "Mid",
        img: "http://img-mid-role-url.com"
    },
    {
        _id: "id-role-adc",
        label: "Adc",
        img: "http://img-adc-role-url.com"
    },
    {
        _id: "id-role-support",
        label: "Support",
        img: "http://img-support-role-url.com"
    },
]