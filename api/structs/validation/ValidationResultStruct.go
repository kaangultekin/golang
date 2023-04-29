package validation

type ValidationResultStruct struct {
	Success     bool
	ErrorFields []map[string]string
}
