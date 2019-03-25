package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/ironarachne/naminglanguage"
	"github.com/ironarachne/random"
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

	firstAdjective := random.Item(languageAdjectives)
	secondAdjective := random.Item(languageAdjectives)

	if firstAdjective == secondAdjective {
		secondAdjective = random.Item(languageAdjectives)
	}

	speakers := random.Item(professionSpeakers)
	kingdomAdjective := random.Item(kingdomAdjectives)
	kingdomType := random.Item(kingdomTypes)

	writingSystemType := random.Item(writingSystemTypes)
	writingSystemStyle := random.Item(writingSystemStyles)

	languageSuffix := random.Item(languageSuffixes)

	languageName := strings.Title(kingdomNamer.Name + languageSuffix)

	description := languageName + ". A " + firstAdjective + ", " + secondAdjective + " language. Spoken mostly in the " + kingdomAdjective + " " + kingdomType + " of " + strings.Title(kingdomNamer.Name) + ", and sometimes by " + speakers + ". It uses a writing system that's " + writingSystemType + " and composed primarily of " + writingSystemStyle + " characters."

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
