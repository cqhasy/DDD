package infrastruce

type ResponseA struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
	Number int    `json:"number"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}
type ResponseB struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
