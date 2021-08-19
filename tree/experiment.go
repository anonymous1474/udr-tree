package tree

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/anonymous1474/udr-tree/protos"
)

func (node *ReplicaNode) experiment(rounds, rate, siz int) {
	time.Sleep((25 - time.Duration(node.replicaID)) * time.Second)

	// Generate random operations
	for i := 0; i < rounds; i++ {
		genKey := (rand.Intn(10000) % siz) + 1
		genVal := (rand.Intn(10000) % siz) + 1
		sendKey := strconv.Itoa(genKey)
		sendVal := strconv.Itoa(genVal)

		// Sleep for random time
		time.Sleep(time.Duration(rate) * time.Microsecond)

		if genKey == genVal {
			//fmt.Printf(Red + "Invalid key value pair\n")
			continue
		}

		crdt.Lock() // Acquire Lock
		start := time.Now()
		totalops++

		node.lamport += 1
		//fmt.Printf(Yellow+"Replica %v : %s -> %s --- Time = %v\n", node.replicaID, sendKey, sendVal, node.lamport)

		// Add new entry to log
		node.LogHistory = append(node.LogHistory,
			Entry{
				currentNode: sendKey,
				newParent:   sendVal,
				oldParent:   node.store[sendKey],
				lamportTime: node.lamport,
				originID:    node.replicaID,
			})

		msg := &protos.UpdateValue{
			Key:   sendKey,
			Value: sendVal,
			Prev:  node.store[sendKey],
			Clock: node.lamport,
			ID:    node.replicaID,
		}

		node.Apply(sendKey, sendVal)

		elapsed := time.Since(start)
		measure1 += elapsed
		crdt.Unlock() // Release Lock

		go node.SendOps(msg) // Send to other peers
	}

	time.Sleep(5 * time.Second)
	var done int
	fmt.Printf(Cyan + "Finished operations")
	fmt.Scanf("%v", &done)
	crdt.Lock() // Acquire Lock
	vertices := make([]string, 0, siz)

	//vertices = append(vertices, "ROOT")
	//vertices = append(vertices, "TRASH")
	//vertices = append(vertices, "CONFLICT")

	for i := 1; i <= siz; i++ {
		key := strconv.Itoa(i)
		vertices = append(vertices, node.store[key])
		//fmt.Printf(Green+"%s -> %s\n", key, node.store[key])
	}

	terminate := &protos.Verify{
		Vertices: vertices,
		ID:       int32(siz),
	}

	node.SendOps2(terminate)

	/*for _, v := range vertices {
		fmt.Printf(Green+"%s -> %s\n", v, node.store[v])
	}*/

	// Release Lock
	fmt.Printf(Green+"Local time = %s, Remote Time = %s, Total ops = %v, Total undo-redo = %v\n", measure1, measure2, totalops, compensate)
	crdt.Unlock()
}

func (node *ReplicaNode) SendOps(msg *protos.UpdateValue) {
	for i := 0; i < node.peers; i++ {
		if int32(i) == node.replicaID {
			continue
		}
		clientObj := node.peerReplica[i]

		go func(node *ReplicaNode, clientObj protos.ChatServiceClient, msg *protos.UpdateValue) {

			ctx, cancel := context.WithCancel(context.Background())
			var ack error
			_, ack = clientObj.Propogate(ctx, msg)
			cancel()

			for ack != nil {
				_, ack = clientObj.Propogate(ctx, msg)
			}
		}(node, clientObj, msg)
	}
}

func (node *ReplicaNode) SendOps2(msg *protos.Verify) {
	for i := 0; i < node.peers; i++ {
		if int32(i) == node.replicaID {
			continue
		}
		clientObj := node.peerReplica[i]

		go func(node *ReplicaNode, clientObj protos.ChatServiceClient, msg *protos.Verify) {

			ctx, cancel := context.WithCancel(context.Background())
			var ack error
			_, ack = clientObj.CheckAnswer(ctx, msg)
			cancel()

			for ack != nil {
				_, ack = clientObj.CheckAnswer(ctx, msg)
			}
		}(node, clientObj, msg)
	}
}

func (node *ReplicaNode) Apply(cur, newp string) {
	// check if key value pair is valid
	localCycle := node.isCycle(cur, newp)
	if localCycle == true {
		//fmt.Printf(Red + "Cycle Found\n")
		return
	}

	node.store[cur] = newp
}
