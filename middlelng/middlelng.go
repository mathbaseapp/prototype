package middlelng

// MiddleLanguage 中間言語を表す
type MiddleLanguage interface {
	Map(callback func(MiddleLanguage) interface{}) []interface{}
}
