package gocsv

import "time"

type Sample struct {
	Foo  string  `csv:"foo"`
	Bar  int     `csv:"BAR"`
	Baz  string  `csv:"Baz"`
	Frop float64 `csv:"Quux"`
	Blah *int    `csv:"Blah"`
	SPtr *string `csv:"SPtr"`
	Omit *string `csv:"Omit,omitempty"`
}

type EmbedSample struct {
	Qux string `csv:"first"`
	Sample
	Ignore string  `csv:"-"`
	Grault float64 `csv:"garply"`
	Quux   string  `csv:"last"`
}

type SkipFieldSample struct {
	EmbedSample
	MoreIgnore string `csv:"-"`
	Corge      string `csv:"abc"`
}

// Testtype for unmarshal/marshal functions on renamed basic types
type RenamedFloat64Unmarshaler float64
type RenamedFloat64Default float64

type RenamedSample struct {
	RenamedFloatUnmarshaler RenamedFloat64Unmarshaler `csv:"foo"`
	RenamedFloatDefault     RenamedFloat64Default     `csv:"bar"`
}

type MultiTagSample struct {
	Foo string `csv:"Baz,foo"`
	Bar int    `csv:"BAR"`
}

type TagSeparatorSample struct {
	Foo string `csv:"Baz|foo"`
	Bar int    `csv:"BAR"`
}

type DateTime struct {
	Foo time.Time `csv:"Foo"`
}

type Level0Struct struct {
	Level0Field level1Struct `csv:"-"`
}

type level1Struct struct {
	Level1Field level2Struct `csv:"-"`
}

type level2Struct struct {
	InnerStruct
}

type InnerStruct struct {
	BoolIgnoreField0 bool   `csv:"-"`
	BoolField1       bool   `csv:"boolField1"`
	StringField2     string `csv:"stringField2"`
}
