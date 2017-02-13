package house

import (
	"fmt"
	"strings"
)

const testVersion = 1

type sentence struct {
	subject string
	verb    string
}

var mapping = map[int]sentence{
	12: sentence{"the horse and the hound and the horn", "belonged to"},
	11: sentence{"the farmer sowing his corn", "kept"},
	10: sentence{"the rooster that crowed in the morn", "woke"},
	9:  sentence{"the priest all shaven and shorn", "married"},
	8:  sentence{"the man all tattered and torn", "kissed"},
	7:  sentence{"the maiden all forlorn", "milked"},
	6:  sentence{"the cow with the crumpled horn", "tossed"},
	5:  sentence{"the dog", "worried"},
	4:  sentence{"the cat", "killed"},
	3:  sentence{"the rat", "ate"},
	2:  sentence{"the malt", "lay in"},
	1:  sentence{"the house that Jack built", ""},
}

func lyric(n int) string {
	if n == 1 {
		return mapping[n].subject
	} else {
		return fmt.Sprintf("%s\nthat %s %s", mapping[n].subject, mapping[n].verb, lyric(n-1))
	}
}

func Verse(n int) string {
	return "This is " + lyric(n) + "."
}

func Song() string {
	s := ""
	for i := 1; i <= len(mapping); i++ {
		s = s + Verse(i) + "\n\n"
	}
	return strings.TrimSpace(s)
}
