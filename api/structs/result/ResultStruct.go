package result

type ResultStruct struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Datas   interface{} `json:"datas,omitempty"`
}
