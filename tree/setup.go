package tree

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/anonymous1474/udr-tree/protos"
	"google.golang.org/grpc"
)

type Entry struct {
	currentNode string
	newParent   string
	oldParent   string
	lamportTime int32
	originID    int32
}

type ReplicaNode struct {
	replicaID   int32
	peerReplica []protos.StreamServiceClient
	final       []protos.ChatServiceClient
	LogHistory  []Entry
	address     string
	myServer    *grpc.Server
	peers       int
	lamport     int32
	store       map[string]string
}

var crdt sync.Mutex
var measure1, measure2 time.Duration
var compensate, totalops, totalops2 int

var ch1, ch2, ch3 chan Entry

func SetupReplica(num, id, rounds, rate, siz int) {
	grpc_address := ":500" + strconv.Itoa(id)
	lis, err := net.Listen("tcp", grpc_address)
	CheckFatalError(err)
	measure1 = 0
	measure2 = 0
	compensate = 0
	totalops = 0
	totalops2 = 0

	grpcServer := grpc.NewServer()
	node := &ReplicaNode{
		replicaID:   int32(id),
		peerReplica: make([]protos.StreamServiceClient, num),
		final:       make([]protos.ChatServiceClient, num),
		LogHistory:  make([]Entry, 0),
		address:     grpc_address,
		myServer:    grpcServer,
		peers:       int(num),
		lamport:     0,
		store:       make(map[string]string),
	}
	protos.RegisterStreamServiceServer(grpcServer, node)
	protos.RegisterChatServiceServer(grpcServer, node)

	rep_addrs := make([]string, num)
	// azure vm ip address
	rep_addrs[0] = "10.1.0.6:5000"
	rep_addrs[1] = "10.2.0.5:5001"
	rep_addrs[2] = "10.0.0.7:5002"

	// Loclal run ip address
	/*for i := 0; i < num; i++ {
		rep_addrs[i] = ":500" + strconv.Itoa(i)
	}*/
	node.connectRest(rep_addrs)

	node.store["ROOT"] = "ROOT"
	node.store["TRASH"] = "ROOT"
	node.store["CONFLICT"] = "ROOT"

	for i := 1; i <= siz; i++ {
		key := strconv.Itoa(i)
		node.store[key] = "ROOT"
	}

	ch1 = make(chan Entry, 10000)
	ch2 = make(chan Entry, 10000)
	ch3 = make(chan Entry, 10000)

	go node.experiment(rounds, rate, siz)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// Connect to other Peers
func (node *ReplicaNode) connectRest(rep_addrs []string) {
	client_obj := make([]protos.StreamServiceClient, node.peers)
	final_obj := make([]protos.ChatServiceClient, node.peers)

	for i := 0; i < node.peers; i++ {
		if int32(i) == node.replicaID {
			continue
		}

		connxn, err := grpc.Dial(rep_addrs[i], grpc.WithInsecure())
		CheckFatalError(err)
		fmt.Println("Connected to replica ", i)
		cli := protos.NewStreamServiceClient(connxn)
		cli2 := protos.NewChatServiceClient(connxn)
		client_obj[i] = cli
		final_obj[i] = cli2
	}
	node.peerReplica = client_obj
	node.final = final_obj
}
