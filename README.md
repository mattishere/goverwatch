# Goverwatch
An Overwatch 2 statistics & profile library.

Goverwatch relies on [Colly](https://github.com/gocolly/colly) to scrape information about an Overwatch 2 player from the [Overwatch 2 stats website](https://overwatch.blizzard.com/en-us/career/).

It provides an elegant data structure that's easy to understand and use in your projects.

## TOC
- [Features](#features)
- [Limitations](#limitations)
- [Installation & Usage](#installation--usage)

## Features
- Player's profile (title, endorsement level icon, profile picture)
- PC Ranks (rank, divison & icon)
- An easy to understand and intuitive data structure
- Useful constants

## Limitations
> [!NOTE]
> I'm already looking into ways to better improve the library - I'm currently in contact with Blizzard support, and am looking into alternative scraping libraries to extend the featureset of Goverwatch.
- Due to Colly not supporting **dynamic content**, the current version of Goverwatch is limited to only PC stats and no hero statistics.
- The [source website](https://overwatch.blizzard.com/en-us/career/) is fairly slow and rather confusing in its design (it seems to only show competitive statistics for the current season).

## Installation & Usage
You can install Goverwatch by adding it to the `go.mod`, then running `go mod tidy` (or however you prefer installing libraries):
```
...

require (
    github.com/mattishere/goverwatch latest
)

...
```

Then, you can import it into a Go file:
```Go
package main

import (
	"fmt"

	"github.com/mattishere/goverwatch"
)

func main() {
	stats, err := goverwatch.GetStats("MattHere", 2211)
	if err != nil {
		panic(err)
	}

	fmt.Println(stats)
}
```

And that's it! There are a couple more things you should know, though. The `Stats` struct looks like this:
```Go
type Stats struct {
	Profile (
		Name            string
		Tag             int
		ProfilePicture  string
		Title           string
		EndorsementIcon string
		URL             string // The URL to the stats website.
	)
	Ranks (
		Tank (
			Rank     string
			Division int
			Icon     string
		)
		DPS (...)  
		Support (...)
	)
}
```

Furthermore, you can access some of the submodules the package has to offer:
- **`github.com/mattishere/goverwatch/url`** exposes constants such as a role map (icon: role), role icons and the base URL. You can also find the `GenerateURL()` helper function that lets you easily generate a new URL for a player's name and tag.
- **`github.com/mattishere/goverwatch/data`** exposes all of the data structures used in the library.