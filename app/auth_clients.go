package app

import (
	"sync"
	"time"
)

type AuthClient struct {
	clients  map[string]time.Time
	mutex    sync.RWMutex
	password string
}

func GetAuthClient() *AuthClient {
	clients := &AuthClient{
		clients: make(map[string]time.Time),
	}

	return clients
}

func (a *AuthClient) SetClient(ip string, password string) bool {
	return true
}

func (a *AuthClient) VerifyClient(ip string) bool {
	return true
}
