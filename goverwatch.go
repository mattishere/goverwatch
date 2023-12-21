package goverwatch

import (
	"path"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var (
	roleURLS = map[string]string{
		"https://static.playoverwatch.com/img/pages/career/icons/role/tank-f64702b684.svg#icon":    "tank",
		"https://static.playoverwatch.com/img/pages/career/icons/role/offense-ab1756f419.svg#icon": "dps",
		"https://static.playoverwatch.com/img/pages/career/icons/role/support-0258e13d85.svg#icon": "support",
	}

	base = "https://overwatch.blizzard.com/en-us/career/"
)

// Given a name and discriminator (e.g. MattHere, 2211), returns a map of roles and their individual ranks.
// You can get a rank for a specific role in the map by using either of these three strings: "tank", "dps", "support".
// There is a chance that the map will be empty if the user has not played competitive. Additionally, if the user has not played a specific role in competitive, that role will not be in the map.
func GetRoles(name string, discriminator int) (roles map[string]Role, err error) {
	c := colly.NewCollector()

	roles = make(map[string]Role)

	var rankedRoles []string

	c.OnHTML(".Profile-playerSummary--role img", func(e *colly.HTMLElement) {
		src := e.Attr("src")

		rankedRoles = append(rankedRoles, roleURLS[src])
	})

	var i int
	c.OnHTML(".Profile-playerSummary--rank", func(e *colly.HTMLElement) {
		currentRole := rankedRoles[i]
		imgPath := path.Base(e.Attr("src"))

		rank, roleError := getRole(currentRole, imgPath)
		if roleError != nil {
			err = roleError
			return
		}
		roles[currentRole] = rank

		i++
	})

	c.Visit(base + name + "-" + strconv.Itoa(discriminator))

	return roles, err
}

func getRole(role, source string) (Role, error) {
	clean := strings.TrimSuffix(source, path.Ext(source))
	clean = strings.TrimPrefix(clean, "rank/")

	elements := strings.Split(clean, "-")
	rank := strings.ReplaceAll(elements[0], "Tier", "")
	div, err := strconv.Atoi(elements[1])
	if err != nil {
		return Role{}, err
	}

	return Role{
		Name: role,
		Rank: Rank{
			Rank:    rank,
			Divison: div,
			Icon:    source,
		},
	}, nil
}
