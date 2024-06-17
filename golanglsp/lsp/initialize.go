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

type InitilizeResponse struct {
	Response
	Result InitilizeResult `json:"result"`
}
type InitilizeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverinfo"`
}

type ServerCapabilities struct {
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
