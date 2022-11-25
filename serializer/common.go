package serializer

type Response struct {
	Status int    `json:"status"`
	Data   string `json:"data"`
	Msg    string `json:"msg"`
	Err    string `json:"err"`
}
