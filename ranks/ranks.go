package ranks

import (
	"path"
	"strconv"
	"strings"

	"github.com/mattishere/goverwatch/data"
)

// A helper function to get the rank from a rank image name.
func GetRoleRank(source string) (data.Rank, error) {
	clean := strings.TrimSuffix(source, path.Ext(source))
	clean = strings.TrimPrefix(clean, "rank/")

	elements := strings.Split(clean, "-")
	rank := strings.ReplaceAll(elements[0], "Tier", "")
	division, err := strconv.Atoi(elements[1])
	if err != nil {
		return data.Rank{}, err
	}

	return data.Rank{
		Rank:     rank,
		Division: division,
		Icon:     source,
	}, nil
}
