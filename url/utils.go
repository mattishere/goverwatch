package url

import "strconv"

// A helper function to generate the URL for a player.
func GenerateURL(name string, discriminator int) (url string) {
	return BaseURL + name + "-" + strconv.Itoa(discriminator)
}
