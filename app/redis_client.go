package app

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	clientWriter *redis.Client
	clientReader *redis.Client
	funcionando  bool
}

func getRedisUrlMaster() string {
	urlMaster := os.Getenv("REDIS_PRIMARY_NODE_URL")

	if urlMaster != "" {
		return urlMaster
	}

	return "localhost:6379"
}

func getRedisUrlSlave() string {
	urlSlave := os.Getenv("REDIS_REPLICA_NODE_URL")

	if urlSlave != "" {
		return urlSlave
	}

	return "localhost:6379"
}

func getRedisPasswordMaster() string {
	passwordMaster := os.Getenv("REDIS_PRIMARY_NODE_PASSWORD")

	return passwordMaster
}

func getRedisPasswordSlave() string {
	passwordSlave := os.Getenv("REDIS_REPLICA_NODE_PASSWORD")

	return passwordSlave
}

func (r *RedisClient) NewRedisWrite() {
	r.clientWriter = redis.NewClient(&redis.Options{
		Addr:     getRedisUrlMaster(),
		Password: getRedisPasswordMaster(), // no password set
		DB:       0,                        // use default DB
	})

	pong, err := r.clientWriter.Ping().Result()

	if err != nil || pong != "PONG" {
		fmt.Println(err)
		r.funcionando = false
	}

	return
}

func (r *RedisClient) NewRedisReader() {
	r.clientReader = redis.NewClient(&redis.Options{
		Addr:     getRedisUrlSlave(),
		Password: getRedisPasswordSlave(), // no password set
		DB:       0,                       // use default DB
	})

	pong, err := r.clientReader.Ping().Result()

	if err != nil || pong != "PONG" {
		fmt.Println(err)
		r.funcionando = false
	}

	return
}

func (r RedisClient) GetFuncionando() bool {
	return r.funcionando
}

// Write Write a key value in redis and set its expiration in milliseconds (set 0 for no expiration)
func (r *RedisClient) Write(chave string, valor string, expiration int) (sucesso bool) {
	err := r.clientWriter.Set(chave, valor, time.Duration(expiration*1000000)).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (r *RedisClient) Delete(chave string) (sucesso bool) {
	err := r.clientWriter.Del(chave).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (r RedisClient) Read(chave string) (valor string, sucesso bool) {
	valor, err := r.clientReader.Get(chave).Result()
	fmt.Println(err)
	if err != nil {
		sucesso = false
		return
	}

	sucesso = true
	return
}

func NewRedisClient() *RedisClient {
	cliente := new(RedisClient)
	cliente.funcionando = true
	cliente.NewRedisWrite()
	cliente.NewRedisReader()

	return cliente
}
