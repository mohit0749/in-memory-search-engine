package preprocessor

type Preprocessor interface {
	Process(doc string) []string
}
