> [!NOTE]
> Due to Season 9 changing a lot, I'm still in the process of updating a lot of stuff! This is a temporary version that "works", but the code is less than stellar. I'll begin improving the website soon!

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
- PC & Console ranks (rank, divison & icon)
- An easy to understand and intuitive data structure layout
- Useful constants (role icons)

## Limitations
- The [source website](https://overwatch.blizzard.com/en-us/career/) is fairly slow and rather confusing in its design (it seems to only show competitive statistics for the current season).

## Installation & Usage
You can install Goverwatch by adding it to the `go.mod`, then running `go mod tidy` (or however you prefer installing libraries):
```
require (
    github.com/mattishere/goverwatch latest
)
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

And that's it!

Furthermore, you can access some of the submodules the package has to offer:
- **`github.com/mattishere/goverwatch/url`** exposes constants such as a role map (icon: role), role icons and the base URL. You can also find the `GenerateURL()` helper function that lets you easily generate a new URL for a player's name and tag.
- **`github.com/mattishere/goverwatch/data`** exposes all of the data structures used in the library.