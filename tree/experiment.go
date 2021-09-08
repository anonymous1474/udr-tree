package tree

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/anonymous1474/udr-tree/protos"
)

// Make random updates
func (node *ReplicaNode) experiment(rounds, rate, siz int) {
	time.Sleep((5 - time.Duration(node.replicaID)) * time.Second)

	done := make(chan bool, 10)

	var peerStream []pb.StreamService_FetchResponseClient
	peerStream = make([]pb.StreamService_FetchResponseClient, 3)
	//vc := make([]int, 3)

	for i := 0; i < 3; i++ {
		if int32(i) == node.replicaID {
			continue
		}
		in := &pb.Request{Id: node.replicaID}
		var err error
		peerStream[i], err = node.peerReplica[i].FetchResponse(context.Background(), in)
		CheckFatalError(err)

		id := i
		go func(id int) {
			for {
				in, err := peerStream[id].Recv()
				if err == io.EOF {
					done <- true //means stream is finished
					return
				}
				if err != nil {
					log.Fatalf("cannot receive %v", err)
				}

				//start := time.Now()
				crdt.Lock() // Acquire Lock
				start := time.Now()
				totalops2++
				//vc[id]++
				//fmt.Println(Reset, vc)
				midpoint := node.Undo(in.Clock, in.ID)
				compensate += midpoint

				prevParent := node.store[in.Key]
				node.Do(in.Key, in.Value, prevParent, in.Clock, in.ID, midpoint)
				node.Redo(midpoint)

				if in.Clock > node.lamport {
					node.lamport = in.Clock
				}

				elapsed := time.Since(start)
				measure2 += elapsed
				crdt.Unlock() // Release Lock
			}
		}(id)
	}

	//Run experiment mentioned number of times

	for i := 0; i < rounds; i++ {

		// Generate random operations & Sleep for interval time
		time.Sleep(time.Duration(rate) * time.Microsecond)
		genKey := (rand.Intn(10000) % siz) + 1
		genVal := (rand.Intn(10000) % siz) + 1
		sendKey := strconv.Itoa(genKey)
		sendVal := strconv.Itoa(genVal)

		if sendKey == sendVal {
			//fmt.Printf(Red + "Invalid key value pair\n")
			continue
		}

		//start := time.Now()
		crdt.Lock() // Acquire Lock
		start := time.Now()
		totalops++
		//vc[node.replicaID]++
		//fmt.Println(Reset, vc)

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

		msg := Entry{
			currentNode: sendKey,
			newParent:   sendVal,
			oldParent:   node.store[sendKey],
			lamportTime: node.lamport,
			originID:    node.replicaID,
		}

		// check if key value pair is valid
		localCycle := node.isCycle(sendKey, sendVal)
		if localCycle != true {
			node.store[sendKey] = sendVal
		}

		// Send to other peers
		ch1 <- msg
		ch2 <- msg
		ch3 <- msg

		elapsed := time.Since(start)
		measure1 += elapsed
		crdt.Unlock() // Release Lock

	}

	time.Sleep(5 * time.Second)

	ch1 <- Entry{"zero", "zero", "zero", -1, node.replicaID}
	ch2 <- Entry{"zero", "zero", "zero", -1, node.replicaID}
	ch3 <- Entry{"zero", "zero", "zero", -1, node.replicaID}

	<-done
	<-done
	fmt.Printf(Green+"Local time = %s, Remote Time = %s, totalops = %v, Conflicts = %v\n", measure1/time.Duration(totalops), measure2/time.Duration(totalops2), totalops, compensate)

	vertices := make([]string, 0, siz)

	for i := 1; i <= siz; i++ {
		key := strconv.Itoa(i)
		vertices = append(vertices, node.store[key])
	}

	terminate := &pb.Verify{
		Vertices: vertices,
		ID:       int32(siz),
	}

	node.SendOps(terminate)
}

// step 1
func (node *ReplicaNode) Undo(clk, id int32) int {
	logsize := len(node.LogHistory)
	midpoint := logsize
	for i := logsize - 1; i >= 0; i-- {
		if (node.LogHistory[i].lamportTime < clk) ||
			(node.LogHistory[i].lamportTime == clk && node.LogHistory[i].originID < id) {
			break
		}

		node.Apply(node.LogHistory[i].currentNode, node.LogHistory[i].oldParent)
		midpoint = i
	}
	//fmt.Printf("Midpoint = %v, Total = %v\n", midpoint, logsize)
	return midpoint
}

// step 2
func (node *ReplicaNode) Do(cur, newp, oldp string, clck, id int32, mid int) {
	node.LogHistory = append(node.LogHistory,
		Entry{
			currentNode: cur,
			newParent:   newp,
			oldParent:   oldp,
			lamportTime: clck,
			originID:    id,
		})

	node.Apply(cur, newp)

	total := len(node.LogHistory)
	temp := node.LogHistory[total-1]

	for i := total - 1; i > mid; i-- {
		node.LogHistory[i] = node.LogHistory[i-1]
	}
	node.LogHistory[mid] = temp
}

// step 3
func (node *ReplicaNode) Redo(mid int) {
	logsize := len(node.LogHistory)
	for i := mid + 1; i < logsize; i++ {
		updateOldParent := node.store[node.LogHistory[i].currentNode]
		node.Apply(node.LogHistory[i].currentNode, node.LogHistory[i].newParent)
		node.LogHistory[i].oldParent = updateOldParent

	}
}

func (node *ReplicaNode) Apply(cur, newp string) {
	// check if key value pair is valid
	localCycle := node.isCycle(cur, newp)
	if localCycle == true {
		return
	}

	node.store[cur] = newp
}

func (node *ReplicaNode) SendOps(msg *pb.Verify) {
	for i := 0; i < node.peers; i++ {
		if int32(i) == node.replicaID {
			continue
		}
		clientObj := node.final[i]

		go func(node *ReplicaNode, clientObj pb.ChatServiceClient, msg *pb.Verify) {

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
