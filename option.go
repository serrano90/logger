package log

// Option
type Option struct {
	Level       LogLevel
	Encoding    Encoding
	Fields      map[string]interface{}
	Development bool
}
