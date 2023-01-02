package auth

import "context"

// Authentication todo
type Authentication struct {
	clientID     string
	clientSecret string
}

func NewAuthentication(clientId, clientSecret string) *Authentication {
	return &Authentication{
		clientID:     clientId,
		clientSecret: clientSecret,
	}
}

// 实现PerRPCCredentials接口，返回map认证数据
func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{ClientIdKey: a.clientID, ClientSecretKey: a.clientSecret}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
