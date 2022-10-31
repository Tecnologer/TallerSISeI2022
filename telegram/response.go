package telegram

type response struct {
	OK     bool        `json:"ok"`
	Result interface{} `json:"result"`
}
