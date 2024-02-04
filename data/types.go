package data

// Stats contains the profile and ranks of the player
type Stats struct {
	Profile Profile
	PC      Platform
	Console Platform
}

// Profile contains the profile data of the player
type Profile struct {
	Name            string
	Tag             int
	IsPrivate       bool
	Exists          bool
	ProfilePicture  string
	Title           string
	EndorsementIcon string
	URL             string
}

type Platform struct {
	HasRanks bool
	Ranks    Ranks
}

// Ranks contains the ranks of the player
type Ranks struct {
	Tank    Rank
	DPS     Rank
	Support Rank
}

// Rank contains the rank data of the player
type Rank struct {
	Rank     string
	Division int
	Icon     string
}
