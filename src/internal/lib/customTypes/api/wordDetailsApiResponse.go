package customtypes

import (
	"strings"

	"github.com/fallncrlss/dictionary-app-backend/src/internal/lib/enums"
	"github.com/fallncrlss/dictionary-app-backend/src/pkg/model"
)

type WordDetailsAPIResponse struct {
	ID       string `json:"id"`
	Word     string `json:"word"`
	Metadata struct {
		Operation string `json:"operation,omitempty"`
		Provider  string `json:"provider,omitempty"`
		Schema    string `json:"schema,omitempty"`
	} `json:"metadata"`
	Results []struct {
		ID             string `json:"id,omitempty"`
		Language       string `json:"language,omitempty"`
		Type           string `json:"type,omitempty"`
		Word           string `json:"word,omitempty"`
		LexicalEntries []struct {
			Language        string                              `json:"language,omitempty"`
			Text            string                              `json:"text,omitempty"`
			LexicalCategory responseIdentifiableTextStruct      `json:"lexicalCategory"`
			Derivatives     sliceResponseIdentifiableTextStruct `json:"derivatives"`
			Entries         []entries                           `json:"entries,omitempty"`
			PhrasalVerbs    sliceResponseIdentifiableTextStruct `json:"phrasalVerbs"`
			Phrases         sliceResponseIdentifiableTextStruct `json:"phrases"`
		} `json:"lexicalEntries,omitempty"`
	} `json:"results"`
}

type sense struct {
	ID               string                              `json:"id,omitempty"`
	Definitions      []string                            `json:"definitions,omitempty"`
	ShortDefinitions []string                            `json:"shortDefinitions,omitempty"`
	Constructions    sliceResponseTextStruct             `json:"constructions,omitempty"`
	Examples         sliceResponseTextStruct             `json:"examples,omitempty"`
	DomainClasses    sliceResponseIdentifiableTextStruct `json:"domainClasses,omitempty"`
	Domains          sliceResponseIdentifiableTextStruct `json:"domains,omitempty"`
	SemanticClasses  sliceResponseIdentifiableTextStruct `json:"semanticClasses,omitempty"`
	Synonyms         sliceSynonyms                       `json:"synonyms,omitempty"`
	ThesaurusLinks   []thesaurusLinks                    `json:"thesaurusLinks,omitempty"`
	Notes            sliceNotes                          `json:"notes,omitempty"`
}

type synonyms struct {
	Language string `json:"language,omitempty"`
	Text     string `json:"text,omitempty"`
}

type thesaurusLinks struct {
	EntryID string `json:"entry_id,omitempty"`
	SenseID string `json:"sense_id,omitempty"`
}

type entries struct {
	Etymologies    []string            `json:"etymologies,omitempty"`
	Notes          sliceNotes          `json:"notes,omitempty"`
	Pronunciations slicePronunciations `json:"pronunciations,omitempty"`
	Inflections    []inflections       `json:"inflections"`
	Senses         sliceSenses         `json:"senses,omitempty"`
}

type notes struct {
	Text string `json:"text,omitempty"`
	Type string `json:"type,omitempty"`
}

type pronunciation struct {
	AudioFile        string   `json:"audioFile"`
	Dialects         []string `json:"dialects,omitempty"`
	PhoneticNotation string   `json:"phoneticNotation,omitempty"`
	PhoneticSpelling string   `json:"phoneticSpelling,omitempty"`
}

type inflections struct {
	InflectedForm       string                              `json:"inflectedForm,omitempty"`
	Regions             sliceResponseIdentifiableTextStruct `json:"regions,omitempty"`
	Registers           sliceResponseIdentifiableTextStruct `json:"registers,omitempty"`
	GrammaticalFeatures []struct {
		ID   string `json:"id,omitempty"`
		Text string `json:"text,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"grammaticalFeatures,omitempty"`
}

type responseIdentifiableTextStruct struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

type responseTextStruct struct {
	Text string `json:"text,omitempty"`
}

type (
	sliceResponseIdentifiableTextStruct []responseIdentifiableTextStruct
	sliceResponseTextStruct             []responseTextStruct
	sliceSynonyms                       []synonyms
	sliceNotes                          []notes
	slicePronunciations                 []pronunciation
	sliceSenses                         []struct {
		sense
		SubSenses []sense
	}
)

// GetTextSlice TODO: there is the way to remove copy-paste of the GetTextSlice in different structures?
func (s sliceResponseIdentifiableTextStruct) GetTextSlice() []string {
	result := make([]string, len(s))

	for i, item := range s {
		result[i] = item.Text
	}

	return result
}

func (s sliceResponseTextStruct) GetTextSlice() []string {
	result := make([]string, len(s))

	for i, item := range s {
		result[i] = strings.TrimSpace(item.Text)
	}

	return result
}

func (s sliceNotes) GetTextSlice() []string {
	result := make([]string, len(s))

	for i, item := range s {
		result[i] = item.Text
	}

	return result
}

func (s sliceSynonyms) GetTextSlice() []string {
	result := make([]string, len(s))

	for i, item := range s {
		result[i] = item.Text
	}

	return result
}

func (s sliceSenses) ToModel() []model.WordSense {
	result := make([]model.WordSense, len(s))

	for i, sense := range s {
		result[i] = model.WordSense{
			Definitions:      sense.Definitions,
			ShortDefinitions: sense.ShortDefinitions,
			Examples:         sense.Examples.GetTextSlice(),
			Constructions:    sense.Constructions.GetTextSlice(),
			Domains:          sense.Domains.GetTextSlice(),
			DomainClasses:    sense.DomainClasses.GetTextSlice(),
			SemanticClasses:  sense.SemanticClasses.GetTextSlice(),
			Synonyms:         sense.Synonyms.GetTextSlice(),
			Notes:            sense.Notes.GetTextSlice(),
		}
	}

	return result
}

func (s slicePronunciations) Get() pronunciation {
	for _, item := range s {
		if item.PhoneticNotation == enums.IPA {
			return item
		}
	}

	if len(s) == 1 {
		return s[0]
	}

	return pronunciation{}
}
