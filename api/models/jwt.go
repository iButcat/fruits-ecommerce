package models

type Jwt struct {
	Header     map[string]string
	Payload    map[string]interface{}
	Validation map[string]interface{}
}
