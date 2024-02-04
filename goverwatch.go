package goverwatch

import (
	"path"

	"github.com/gocolly/colly"
	"github.com/mattishere/goverwatch/data"
	"github.com/mattishere/goverwatch/ranks"
	"github.com/mattishere/goverwatch/url"
)

// Returns a Stats struct containing the profile and ranks of the player
func GetStats(name string, discriminator int) (stats data.Stats, err error) {
	c := colly.NewCollector()

	link := url.GenerateURL(name, discriminator)

	stats.Profile.Name = name
	stats.Profile.Tag = discriminator
	stats.Profile.URL = link

	// Profile data (should be moved to the profile package eventually!)
	c.OnHTML(".Profile-player--portrait", func(e *colly.HTMLElement) {
		if e.Attr("src") != "" {
			stats.Profile.ProfilePicture = e.Attr("src")
			stats.Profile.Exists = true
		}
	})

	if !stats.Profile.Exists {
		// Is the profile private
		c.OnHTML(".Profile-private---msg", func(e *colly.HTMLElement) {
			stats.Profile.IsPrivate = true
		})

		c.OnHTML(".Profile-player--title", func(e *colly.HTMLElement) {
			stats.Profile.Title = e.Text
		})
		c.OnHTML(".Profile-playerSummary--endorsement", func(e *colly.HTMLElement) {
			stats.Profile.EndorsementIcon = e.Attr("src")
		})

		// Ranks (should be moved to the ranks package eventually!)
		if !stats.Profile.IsPrivate {
			var pcRoles []string
			c.OnHTML(".mouseKeyboard-view .Profile-playerSummary--role img", func(e *colly.HTMLElement) {
				src := e.Attr("src")
				pcRoles = append(pcRoles, url.RoleURLS[src])
				stats.PC.HasRanks = true
			})

			var i int
			c.OnHTML(".mouseKeyboard-view .Profile-playerSummary--rank", func(e *colly.HTMLElement) {
				if len(pcRoles) == 0 {
					return
				}
				role := pcRoles[i]
				imgPath := path.Base(e.Attr("src"))

				rank, roleError := ranks.GetRoleRank(imgPath)
				if roleError != nil {
					err = roleError
					return
				}

				switch role {
				case "tank":
					stats.PC.Ranks.Tank = rank
				case "dps":
					stats.PC.Ranks.DPS = rank
				case "support":
					stats.PC.Ranks.Support = rank
				}

				i++
			})

			var consoleRoles []string
			c.OnHTML(".Profile-playerSummary--role", func(e *colly.HTMLElement) {
				src := e.ChildAttr("use", "href")
				consoleRoles = append(consoleRoles, url.RoleURLS[src])
				stats.Console.HasRanks = true
			})

			var j int
			c.OnHTML(".controller-view .Profile-playerSummary--rank", func(e *colly.HTMLElement) {
				if len(consoleRoles) == 0 {
					return
				}
				role := consoleRoles[j]
				imgPath := path.Base(e.Attr("src"))

				rank, roleError := ranks.GetRoleRank(imgPath)
				if roleError != nil {
					err = roleError
					return
				}

				switch role {
				case "tank":
					stats.Console.Ranks.Tank = rank
				case "dps":
					stats.Console.Ranks.DPS = rank
				case "support":
					stats.Console.Ranks.Support = rank
				}

				j++
			})

		}
	}

	c.Visit(link)
	return stats, err
}
