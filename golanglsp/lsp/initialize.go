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
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitilizeResponse(id int) InitilizeResponse {
	return InitilizeResponse{
		Response: Response{
			ID:  &id,
			RPC: "2.0",
		},
		Result: InitilizeResult{
			Capabilities: ServerCapabilities{},
			ServerInfo: ServerInfo{
				Name:    "golanglsp",
				Version: "0.0.0.0-beta-final",
			},
		},
	}
}
