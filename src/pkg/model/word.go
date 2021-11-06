package model

type WordSense struct {
	Definitions      []string `json:"definitions" validate:"required"`
	ShortDefinitions []string `json:"shortDefinitions" validate:"required"`
	Examples         []string `json:"examples" validate:"required"`
	Constructions    []string `json:"constructions" validate:"required"`
	Domains          []string `json:"domains" validate:"required"`
	DomainClasses    []string `json:"domainClasses" validate:"required"`
	SemanticClasses  []string `json:"semanticClasses" validate:"required"`
	Synonyms         []string `json:"synonyms" validate:"required"`
	Notes            []string `json:"notes" validate:"required"`
}

type Word struct {
	Name          string      `json:"name" validate:"required"`
	Language      string      `json:"language" validate:"required"`
	Phonetic      string      `json:"phonetic,omitempty"`
	PhoneticAudio string      `json:"phoneticAudio,omitempty"`
	Senses        []WordSense `json:"senses,omitempty"`
	PhrasalVerbs  []string    `json:"phrasalVerbs,omitempty"`
	Phrases       []string    `json:"phrases,omitempty"`
	Notes         []string    `json:"notes,omitempty"`
}
