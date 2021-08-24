package models

import "github.com/robinsond76/bookings/internal/forms"

//Template data holds data sent from handler to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string //cross site request forgery token
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
