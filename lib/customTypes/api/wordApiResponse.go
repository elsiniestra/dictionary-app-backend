package api

import (
	"github.com/fallncrlss/dictionary-app-backend/model"
)

type WordApiResponse struct {
	Id       string   `json:"id"`
	Word     string   `json:"word"`
	Metadata metadata `json:"metadata"`
	Results  []struct {
		Id             string `json:"id,omitempty"`
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
	Id               string                              `json:"id,omitempty"`
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

type metadata struct {
	Operation string `json:"operation,omitempty"`
	Provider  string `json:"provider,omitempty"`
	Schema    string `json:"schema,omitempty"`
}

type synonyms struct {
	Language string `json:"language,omitempty"`
	Text     string `json:"text,omitempty"`
}

type thesaurusLinks struct {
	EntryId string `json:"entry_id" json:"entry_id,omitempty"`
	SenseId string `json:"sense_id" json:"sense_id,omitempty"`
}

type entries struct {
	Etymologies    []string         `json:"etymologies,omitempty"`
	Notes          sliceNotes       `json:"notes,omitempty"`
	Pronunciations []pronunciations `json:"pronunciations,omitempty"`
	Inflections    []inflections    `json:"inflections"`
	Senses         sliceSenses      `json:"senses,omitempty"`
}

type notes struct {
	Text string `json:"text,omitempty"`
	Type string `json:"type,omitempty"`
}

type pronunciations struct {
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
		Id   string `json:"id,omitempty"`
		Text string `json:"text,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"grammaticalFeatures,omitempty"`
}

type responseIdentifiableTextStruct struct {
	Id   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

type responseTextStruct struct {
	Text string `json:"text,omitempty"`
}

type sliceResponseIdentifiableTextStruct []responseIdentifiableTextStruct
type sliceResponseTextStruct []responseTextStruct
type sliceSynonyms []synonyms
type sliceNotes []notes
type sliceSenses []struct {
	sense
	SubSenses []sense
}

// GetTextSlice TODO: there is the way to remove copy-paste of the GetTextSlice in different structures?
func (s sliceResponseIdentifiableTextStruct) GetTextSlice() []string {
	var result []string
	for _, item := range s {
		result = append(result, item.Text)
	}
	return result
}
func (s sliceResponseTextStruct) GetTextSlice() []string {
	var result []string
	for _, item := range s {
		result = append(result, item.Text)
	}
	return result
}

func (s sliceNotes) GetTextSlice() []string {
	var result []string
	for _, item := range s {
		result = append(result, item.Text)
	}
	return result
}

func (s sliceSynonyms) GetTextSlice() []string {
	var result []string
	for _, item := range s {
		result = append(result, item.Text)
	}
	return result
}

func (s sliceSenses) ToModel() []model.WordSense {
	var result []model.WordSense

	for _, sense := range s {
		result = append(
			result,
			model.WordSense{
				Definitions:      sense.Definitions,
				ShortDefinitions: sense.ShortDefinitions,
				Examples:         sense.Examples.GetTextSlice(),
				Constructions:    sense.Constructions.GetTextSlice(),
				Domains:          sense.Domains.GetTextSlice(),
				DomainClasses:    sense.DomainClasses.GetTextSlice(),
				SemanticClasses:  sense.SemanticClasses.GetTextSlice(),
				Synonyms:         sense.Synonyms.GetTextSlice(),
				Notes:            sense.Notes.GetTextSlice(),
			},
		)
	}

	return result
}
