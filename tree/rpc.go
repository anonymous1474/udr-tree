package tree

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/anonymous1474/udr-tree/protos"
)

// RPC Call
func (node *ReplicaNode) Propogate(ctx context.Context, in *protos.UpdateValue) (*protos.Void, error) {

	crdt.Lock() // Acquire Lock
	start := time.Now()
	totalops++
	//fmt.Printf(Yellow+"Replica %v : %s -> %s --- Time = %v\n", in.ID, in.Key, in.Value, in.Clock)
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
	return &protos.Void{}, nil
}

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

func (node *ReplicaNode) Redo(mid int) {
	logsize := len(node.LogHistory)
	for i := mid + 1; i < logsize; i++ {
		updateOldParent := node.store[node.LogHistory[i].currentNode]
		node.Apply(node.LogHistory[i].currentNode, node.LogHistory[i].newParent)
		node.LogHistory[i].oldParent = updateOldParent

	}
}

func (node *ReplicaNode) CheckAnswer(ctx context.Context, in *protos.Verify) (*protos.Void, error) {
	crdt.Lock()
	ans := 0
	for i := 1; i <= int(in.ID); i++ {
		key := strconv.Itoa(i)
		if node.store[key] != in.Vertices[i-1] {
			ans += 1
			//fmt.Printf(Purple + "Not Matching with replica\n")
			//crdt.Unlock()
			//return &protos.Void{}, nil
		}
	}
	fmt.Printf(Red+"Matching with replica %v\n", ans)
	crdt.Unlock()
	return &protos.Void{}, nil
}
