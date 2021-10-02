package enums

const (
	AmericanEnglishCode = "en-us"
	BritishEnglishCode  = "en-gb"
	AmericanEnglishName = "American English"
	BritishEnglishName  = "British English"
)

var LanguageCodes = map[string]string{
	AmericanEnglishCode: AmericanEnglishName,
	BritishEnglishCode:  BritishEnglishName,
}
