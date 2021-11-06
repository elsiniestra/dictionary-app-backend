package enums

const (
	AmericanEnglishCode = "en-us"
	BritishEnglishCode  = "en-gb"
	AmericanEnglishName = "American English"
	BritishEnglishName  = "British English"
)

func GetLanguageCodes() map[string]string {
	return map[string]string{
		AmericanEnglishCode: AmericanEnglishName,
		BritishEnglishCode:  BritishEnglishName,
	}
}
