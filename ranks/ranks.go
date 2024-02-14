package ranks

import (
	"path"
	"strconv"
	"strings"
)

// Returns the string value of the rank and an error
func GetRoleRank(source string) string {
	base := path.Base(source)
	clean := strings.TrimSuffix(base, path.Ext(source))
	clean = strings.TrimPrefix(clean, "Rank_")

	elements := strings.Split(clean, "-")
	rank := strings.ReplaceAll(elements[0], "Tier", "")

	return rank
}

func GetRoleDivision(source string) (int, error) {
	base := path.Base(source)
	clean := strings.TrimSuffix(base, path.Ext(source))
	clean = strings.TrimPrefix(clean, "TierDivision_")

	elements := strings.Split(clean, "-")
	division, err := strconv.Atoi(elements[0])

	return division, err
}