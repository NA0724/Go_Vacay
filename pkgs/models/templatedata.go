package models

// holds data set from handlets to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	DataMap   map[string]interface{}
	CSRFToken string
	FlashMsg  string
	Warning   string
	Error     string
}
