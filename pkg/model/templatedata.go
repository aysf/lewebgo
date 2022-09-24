package model

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CSRFToken string
	Warning   string
	Error     string
	Flash     string
}
