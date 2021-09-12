
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
            },
            fill: {
                selected: false,
                roleInfo: "id-role-fill"
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
            maxRank: "id-rank-gold-1"
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
            },
            fill: {
                selected: false,
                roleInfo: "id-role-fill"
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
        img: "/lol-ranks/iron-4.png"
    },
    {
        _id: "id-rank-iron-3",
        label: "Iron 3",
        img: "/lol-ranks/iron-3.png"
    },
    {
        _id: "id-rank-iron-2",
        label: "Iron 2",
        img: "/lol-ranks/iron-2.png"
    },
    {
        _id: "id-rank-iron-1",
        label: "Iron 1",
        img: "/lol-ranks/iron-1.png"
    },
    {
        _id: "id-rank-bronze-4",
        label: "Bronze 4",
        img: "/lol-ranks/bronze-4.png"
    },
    {
        _id: "id-rank-bronze-3",
        label: "Bronze 3",
        img: "/lol-ranks/bronze-3.png"
    },
    {
        _id: "id-rank-bronze-2",
        label: "Bronze 2",
        img: "/lol-ranks/bronze-2.png"
    },
    {
        _id: "id-rank-bronze-1",
        label: "Bronze 1",
        img: "/lol-ranks/bronze-1.png"
    },
    {
        _id: "id-rank-silver-4",
        label: "Silver 4",
        img: "/lol-ranks/silver-4.png"
    },
    {
        _id: "id-rank-silver-3",
        label: "Silver 3",
        img: "/lol-ranks/silver-3.png"
    },
    {
        _id: "id-rank-silver-2",
        label: "Silver 2",
        img: "/lol-ranks/silver-2.png"
    },
    {
        _id: "id-rank-silver-1",
        label: "Silver 1",
        img: "/lol-ranks/silver-1.png"
    },
    {
        _id: "id-rank-gold-4",
        label: "Gold 4",
        img: "/lol-ranks/gold-4.png"
    },
    {
        _id: "id-rank-gold-3",
        label: "Gold 3",
        img: "/lol-ranks/gold-3.png"
    },
    {
        _id: "id-rank-gold-2",
        label: "Gold 2",
        img: "/lol-ranks/gold-2.png"
    },
    {
        _id: "id-rank-gold-1",
        label: "Gold 1",
        img: "/lol-ranks/gold-1.png"
    },
    {
        _id: "id-rank-platinum-4",
        label: "Platinum 4",
        img: "/lol-ranks/platinum-4.png"
    },
    {
        _id: "id-rank-platinum-3",
        label: "Platinum 3",
        img: "/lol-ranks/platinum-3.png"
    },
    {
        _id: "id-rank-platinum-2",
        label: "Platinum 2",
        img: "/lol-ranks/platinum-2.png"
    },
    {
        _id: "id-rank-platinum-1",
        label: "Platinum 1",
        img: "/lol-ranks/platinum-1.png"
    },
    {
        _id: "id-rank-diamond-4",
        label: "Diamond 4",
        img: "/lol-ranks/diamond-4.png"
    },
    {
        _id: "id-rank-diamond-3",
        label: "Diamond 3",
        img: "/lol-ranks/diamond-3.png"
    },
    {
        _id: "id-rank-diamond-2",
        label: "Diamond 2",
        img: "/lol-ranks/diamond-2.png"
    },
    {
        _id: "id-rank-diamond-1",
        label: "Diamond 1",
        img: "/lol-ranks/diamond-1.png"
    },
    {
        _id: "id-rank-master",
        label: "Master",
        img: "/lol-ranks/master.png"
    },
    {
        _id: "id-rank-grandmaster",
        label: "Grandmaster",
        img: "/lol-ranks/grandmaster.png"
    },
    {
        _id: "id-rank-challenger",
        label: "Challenger",
        img: "/lol-ranks/challenger.png"
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
        img: "/lol-roles/role-top.png"
    },
    {
        _id: "id-role-jungle",
        label: "Jungle",
        img: "/lol-roles/role-jungle.png"
    },
    {
        _id: "id-role-mid",
        label: "Mid",
        img: "/lol-roles/role-mid.png"
    },
    {
        _id: "id-role-adc",
        label: "Adc",
        img: "/lol-roles/role-adc.png"
    },
    {
        _id: "id-role-support",
        label: "Support",
        img: "/lol-roles/role-support.png"
    },
    {
        _id: "id-role-fill",
        label: "Fill",
        img: "/lol-roles/role-fill.png"
    },
]