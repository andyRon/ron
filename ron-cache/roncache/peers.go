package roncache

type PeerPicker interface {
	// PickPeer 根据传入的 key 选择相应节点 PeerGetter
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 对应HTTP 客户端
type PeerGetter interface {
	// Get 从对应 group 查找缓存值
	Get(group string, key string) ([]byte, error)
}
