package tree

func (node *ReplicaNode) isCycle(a, b string) bool {
	for b != "ROOT" {
		if b == a {
			return true
		}
		b = node.store[b]
	}
	return false
}
