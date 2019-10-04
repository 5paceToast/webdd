package util

var (
	// S is the singleton of the Settings
	S settings
)

type settings struct {
	Asciidoc bool
	Dir      string
	Log      bool
	LogLevel string
	Markdown bool
	Safe     bool
}
