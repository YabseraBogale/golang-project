package lsp

type InitilizeRequest struct {
	Request
	Params InitilizeRequestParams `json:"params"`
}

type InitilizeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientinfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
