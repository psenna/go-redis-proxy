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
	authClients := &AuthClient{
		clients: make(map[string]time.Time),
	}

	return authClients
}

func (a *AuthClient) SetClient(ip string, password string) bool {
	if password == a.password {
		a.mutex.Lock()
		a.clients[ip] = time.Now()
		a.mutex.Unlock()
		return true
	}

	return false
}

func (a *AuthClient) VerifyClient(ip string) bool {
	return true
}
