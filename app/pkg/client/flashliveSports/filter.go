package flashliveSports

const (
	DefaultLocale   = "ru_RU"
	DefaultTimeZone = "+3"
)

type Filter struct {
	Locale   string
	TimeZone string
	SportID  string
}
