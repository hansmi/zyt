package xtextlanguage

import "golang.org/x/text/language"

func MustCompose(part ...any) language.Tag {
	tag, err := language.Compose(part...)
	if err != nil {
		panic(err)
	}

	return tag
}
