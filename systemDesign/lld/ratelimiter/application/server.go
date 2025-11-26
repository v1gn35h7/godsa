package application

type RLServer struct {
	statergy RateLimitStrategy
}

func NewRLServer(statergy RateLimitStrategy) *RLServer {
	return &RLServer{
		statergy: statergy,
	}
}

func (s *RLServer) HandleRequest(key string) bool {
	if s.statergy.GiveAccess(key) {
		return true
	}
	return false
}
