# go-redis-proxy
Go redis proxy


Enviroment variables:

DEBUG: If true, print the received commands.

REDIS_PRIMARY_NODE_URL: Principal redis node URL (Write).

REDIS_PRIMARY_NODE_PASSWORD: Principal redis node password.

REDIS_REPLICA_NODE_URL: Replica redis node URL (Read). Inside k8s, the replica's service URL. Todo: Accept a list of replicas. 

REDIS_REPLICA_NODE_PASSWORD: Replica node passwors.

SERVER_ADDR: Address of the proxy's redis server.

SERVER_PASSWORD: Password of the proxy's redis server.
