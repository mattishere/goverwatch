package goverwatch

import "fmt"

// Role is a struct that contains the name of the role and the rank of the role.
// The name of the role can be either "tank", "dps", or "support".
type Role struct {
	Name string
	Rank Rank
}

// String returns a string representation of the role.
func (r Role) String() string {
	return fmt.Sprintf("%s: %s", r.Name, r.Rank.String())
}

// Rank is a struct that contains the rank, division, and icon of a role.
type Rank struct {
	Rank    string
	Divison int
	Icon    string
}

// String returns a string representation of the rank.
func (r Rank) String() string {
	return fmt.Sprintf("%s %d", r.Rank, r.Divison)
}
