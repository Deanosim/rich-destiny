package main

// Empty ones need to be kept!
var largeImageActivityModes = map[string][]string{
	"anniversary": {"Dares of Eternity"},
	"beyondlight": {"Empire Hunt"},
	"control":     {"Control: Competitive", "Control: Quickplay"},
	"crucible": {"Rift", "Quickplay PvP", "Salvage", "Supremacy", "Mayhem", "Team Scorched", "Scorched", "Survival", "Competitive Co-Op",
		"Breakthrough", "Competitive PvP", "Momentum", "Zone Control", "Clash: Competitive", "Countdown", "Showdown", "Elimination",
		"Rumble", "Clash", "Lockdown", "Momentum Control", "The Crucible", "Classic Mix"},
	"destinylogo":        {},
	"doubles":            {"All Doubles"},
	"dungeon":            {},
	"explore":            {},
	"forge":              {"Forge Ignition"},
	"gambit":             {},
	"hauntedforest":      {},
	"ironbanner":         {},
	"lostsector":         {},
	"menagerie":          {"The Menagerie"},
	"nightmarehunt":      {},
	"privatecrucible":    {"Private Matches"},
	"raid":               {},
	"reckoning":          {"The Reckoning"},
	"seasonchosen":       {},
	"seasonhaunted":      {},
	"seasonlost":         {},
	"seasonrisen":        {},
	"seasonsplicer":      {},
	"shadowkeep":         {},
	"socialall":          {"Social", "All"},
	"storypvecoopheroic": {"Heroic Adventure", "Offensive", "Story", "PvE"},
	"strikes":            {"Scored Prestige Nightfall", "Scored Nightfall Strikes", "Nightfall Strikes", "Prestige Nightfall", "Strike", "Vanguard Op", "Nightfall"},
	"thenine":            {"Trials of the Nine", "Trials of the Nine Countdown", "Trials of the Nine Survival"},
	"thewitchqueen":      {},
	"trialsofosiris":     {"The Sundial", "Lighthouse Simulation"},
	"vexoffensive":       {},
	"wellspring":         {},
}

// Maps classID to class name
var classImages = map[int64]string{
	0: "titan",
	1: "hunter",
	2: "warlock",
}

var scoredLostSectors = []string{"K1", "Concealed", "E15", "Perdition", "2A", "Veles", "Quarry", "Scavenger's", "XII", "Drowned", "Starlight", "Aphelion", "Metamorphosis", "Sepulcher", "Extraction"}
var storyMissions = map[string][]string{
	"shadowkeep":    {"A Mysterious Disturbance", "In Search of Answers", "Ghosts of Our Past", "In the Deep", "Beyond"},
	"beyondlight":   {"Darkness's Doorstep", "The New Kell", "Rising Resistance", "The Warrior", "The Technocrat", "The Kell of Darkness", "Sabotaging Salvation", "The Aftermath"},
	"thewitchqueen": {"The Arrival", "The Investigation", "The Ghosts", "The Communion", "The Mirror", "The Cunning", "The Last Chance", "The Ritual", "Memories of"},
}

var raidProgressionMap = map[string][]string{
	"gos":  {"Embrace", "Undergrowth", "The Consecrated Mind", "The Sanctified Mind"},                                                   // 4
	"votd": {"Acquisition", "Collection", "Exhibition", "Dominion"},                                                                     // 4
	"dsc":  {"Desolation / Crypt Security", "Atraks-1, Fallen Exo", "The Descent", "Taniks, The Abomination"},                           // 4
	"lw":   {"Kalli, The Corrupted", "Shuro Chi, The Corrupted", "Morgeth, The Spirekeeper", "The Vault", "Riven of a Thousand Voices"}, // 5
	"vog":  {"Raise the Spire / Confluxes", "The Oracles", "The Templar", "The Gatekeeper", "Atheon, Time's Conflux"},                   // 5
}
