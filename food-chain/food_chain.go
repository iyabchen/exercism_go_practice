package foodchain

// similar to the problem of house

import (
	"fmt"
	"strings"
)

const testVersion = 3

const start = "I know an old lady who swallowed a %s.\n"
const end1 = "I don't know why she swallowed the fly. Perhaps she'll die."
const end2 = "She's dead, of course!"
const swallow_phrase = "She swallowed the %s to catch the %s%s.\n"

type animal struct {
	name    string
	comment string
	special string
	catch   *animal
}

var animals []animal
var stanzas []string
var initialized bool = false

func initialize() {
	if initialized {
		return
	}
	createFoodChain()
	createStanza()
	initialized = true
}

func createFoodChain() {
	fly := animal{name: "fly", comment: "", special: ""}
	spider := animal{
		name:    "spider",
		comment: "It wriggled and jiggled and tickled inside her.",
		special: " that wriggled and jiggled and tickled inside her",
		catch:   &fly,
	}
	bird := animal{
		name:    "bird",
		comment: "How absurd to swallow a bird!",
		special: "",
		catch:   &spider,
	}
	cat := animal{
		name:    "cat",
		comment: "Imagine that, to swallow a cat!",
		special: "",
		catch:   &bird,
	}
	dog := animal{
		name:    "dog",
		comment: "What a hog, to swallow a dog!",
		special: "",
		catch:   &cat,
	}
	goat := animal{
		name:    "goat",
		comment: "Just opened her throat and swallowed a goat!",
		special: "",
		catch:   &dog,
	}
	cow := animal{
		name:    "cow",
		comment: "I don't know how she swallowed a cow!",
		special: "",
		catch:   &goat,
	}
	horse := animal{
		name:    "horse",
		comment: "",
		special: "",
	}
	// The order matters, corresponding to the verse order
	animals = append(animals, fly)
	animals = append(animals, spider)
	animals = append(animals, bird)
	animals = append(animals, cat)
	animals = append(animals, dog)
	animals = append(animals, goat)
	animals = append(animals, cow)
	animals = append(animals, horse)
}

func createStanza() {
	for i := 0; i < len(animals); i++ {
		stanza := ""
		subject := animals[i]

		s1 := fmt.Sprintf(start, subject.name)
		if subject.comment != "" {
			s1 = s1 + subject.comment + "\n"
		}
		predator := subject
		for {
			prey := predator.catch
			if prey != nil {
				s1 = s1 + fmt.Sprintf(swallow_phrase, predator.name, prey.name, prey.special)
				predator = *prey
			} else {
				break
			}
		}
		if subject.name != "horse" {
			stanza = s1 + end1
		} else {
			stanza = s1 + end2
		}
		stanzas = append(stanzas, stanza)
	}

}
func Verse(n int) string {
	if !initialized {
		initialize()
	}
	return stanzas[n-1]
}

func Verses(min, max int) string {
	if !initialized {
		initialize()
	}
	return strings.Join(stanzas[min-1:max], "\n\n")
}

func Song() string {
	if !initialized {
		initialize()
	}
	return strings.Join(stanzas[0:], "\n\n")

}
