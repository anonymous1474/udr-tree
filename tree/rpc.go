package tree

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "github.com/anonymous1474/udr-tree/protos"
)

func (node *ReplicaNode) FetchResponse(in *pb.Request, srv pb.StreamService_FetchResponseServer) error {

	//log.Printf(Reset+"rpc from id : %d", in.Id)

	if in.Id == 0 {
		for {
			count := <-ch1
			if count.lamportTime == -1 {
				break
			}
			//time sleep to simulate network latency
			resp := pb.Response{Key: count.currentNode, Value: count.newParent, Prev: count.oldParent, Clock: count.lamportTime, ID: count.originID}
			//time.Sleep(1 * time.Millisecond)
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}
	} else if in.Id == 1 {
		for {
			count := <-ch2
			if count.lamportTime == -1 {
				break
			}
			//time sleep to simulate network latency
			resp := pb.Response{Key: count.currentNode, Value: count.newParent, Prev: count.oldParent, Clock: count.lamportTime, ID: count.originID}
			//time.Sleep(1 * time.Millisecond)
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}
	} else {
		for {
			count := <-ch3
			if count.lamportTime == -1 {
				break
			}
			//time sleep to simulate network latency
			resp := pb.Response{Key: count.currentNode, Value: count.newParent, Prev: count.oldParent, Clock: count.lamportTime, ID: count.originID}
			//time.Sleep(1 * time.Millisecond)
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}
	}

	return nil
}

func (node *ReplicaNode) CheckAnswer(ctx context.Context, in *pb.Verify) (*pb.Void, error) {
	crdt.Lock()
	count := 0
	for i := 1; i <= int(in.ID); i++ {
		key := strconv.Itoa(i)
		if node.store[key] != in.Vertices[i-1] {
			count++
		}
	}
	fmt.Printf(Red+"Matching with replica %v\n", count)
	crdt.Unlock()
	return &pb.Void{}, nil
}
