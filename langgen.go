package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/utility"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	kingdomNamer := naminglanguage.GeneratePlace()

	firstAdjective := utility.RandomItem(languageAdjectives)
	secondAdjective := utility.RandomItem(languageAdjectives)

	if firstAdjective == secondAdjective {
		secondAdjective = utility.RandomItem(languageAdjectives)
	}

	speakers := utility.RandomItem(professionSpeakers)
	kingdomAdjective := utility.RandomItem(kingdomAdjectives)
	kingdomType := utility.RandomItem(kingdomTypes)

	languageSuffix := utility.RandomItem(languageSuffixes)

	languageName := kingdomNamer.Name + languageSuffix

	description := strings.Title(languageName) + ". A " + firstAdjective + ", " + secondAdjective + " language. Spoken mostly in the " + kingdomAdjective + " " + kingdomType + " of " + strings.Title(kingdomNamer.Name) + ", and sometimes by " + speakers + "."

	fmt.Println(description)
}
