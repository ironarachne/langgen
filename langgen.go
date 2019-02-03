package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/utility"
	"github.com/kataras/iris"
)

// LanguageDescription is a description of a language
type LanguageDescription struct {
	Name        string
	Description string
}

//GenerateDescription generates a random language description
func GenerateDescription() LanguageDescription {
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

	languageName := strings.Title(kingdomNamer.Name + languageSuffix)

	description := languageName + ". A " + firstAdjective + ", " + secondAdjective + " language. Spoken mostly in the " + kingdomAdjective + " " + kingdomType + " of " + strings.Title(kingdomNamer.Name) + ", and sometimes by " + speakers + "."

	languageDescription := LanguageDescription{languageName, description}

	return languageDescription
}

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		rand.Seed(time.Now().UnixNano())
		description := GenerateDescription()
		ctx.JSON(description)
	})

	app.Run(iris.Addr(":7479"))
}
