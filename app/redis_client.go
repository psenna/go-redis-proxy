package app

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type RedisClient struct {
	clientWriter *redis.Client
	clientReader *redis.Client
	funcionando  bool
}

func getRedisUrlMaster() string {
	urlMaster := os.Getenv("REDIS_MASTER_URL")

	if urlMaster != "" {
		return urlMaster
	}

	return "localhost:6379"
}

func getRedisUrlSlave() string {
	urlSlave := os.Getenv("REDIS_SLAVE_URL")

	if urlSlave != "" {
		return urlSlave
	}

	return "localhost:6379"
}

func getRedisPasswordMaster() string {
	passwordMaster := os.Getenv("REDIS_MASTER_PASSWORD")

	return passwordMaster
}

func getRedisPasswordSlave() string {
	passwordSlave := os.Getenv("REDIS_SLAVE_PASSWORD")

	return passwordSlave
}

func (r *RedisClient) NewRedisWrite() {
	r.clientWriter = redis.NewClient(&redis.Options{
		Addr:     getRedisUrlMaster(),
		Password: getRedisPasswordMaster(), // no password set
		DB:       0,                        // use default DB
	})

	pong, err := r.clientWriter.Ping().Result()

	fmt.Println(pong)

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

	fmt.Println(pong)

	if err != nil || pong != "PONG" {
		fmt.Println(err)
		fmt.Println(getRedisPasswordSlave())
		r.funcionando = false
	}

	return
}

func (r RedisClient) GetFuncionando() bool {
	return r.funcionando
}

func (r *RedisClient) Write(chave string, valor string) (sucesso bool) {
	err := r.clientWriter.Set(chave, valor, 0).Err()
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
