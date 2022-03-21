package names

type Culture string

const CultureAny Culture = ""

type Gender string

const (
	GenderAny    Gender = ""
	GenderFemale Gender = "F"
	GenderMale   Gender = "M"
)

type Options struct {
	Culture Culture
	Gender  Gender
}
