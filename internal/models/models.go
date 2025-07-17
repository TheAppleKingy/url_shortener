package models

type ReqData struct {
	Url string `json:"url"`
}

type RespData struct {
	Result string `json:"result"`
}

type ToSave struct {
	Id  int
	Old string
	New string
}
