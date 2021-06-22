package signaling

// URI default signaling server
const URI = "https://signal.tienlx.tools/socket.io/"

// ConnectInfo SDP by offer or answer
type ConnectInfo struct {
	Source string `json:"source"`
	SDP    string `json:"sdp"`
}
