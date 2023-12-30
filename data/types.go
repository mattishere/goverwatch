package data

// Stats contains the profile and ranks of the player
type Stats struct {
	Profile Profile
	Ranks   Ranks
}

// Profile contains the profile data of the player
type Profile struct {
	Name            string
	Tag             int
	Private         bool
	Exists          bool
	ProfilePicture  string
	Title           string
	EndorsementIcon string
	URL             string
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
