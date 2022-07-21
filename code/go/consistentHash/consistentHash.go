package consistentHash

import (
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type uints []uint32

func (u uints) Len() int {
	return len(u)
}

func (u uints) Less(i, j int) bool {
	return u[i] < u[j]
}

func (u uints) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// 生成hash函数
type Hash func(data []byte) uint32

// hash一致性信息
type Map struct {
	hash     Hash              // 计算 hash 的函数
	replicas int               // 这个是副本数，这里影响到虚拟节点的个数,虚拟节点数=replicas*len(keys)
	keys     uints             // 有序的列表，从大到小排序的，这个很重要
	hashMap  map[uint32]string // 可以理解成用来记录虚拟节点和物理节点元数据关系的
	mux      sync.Mutex
}

// New 初始化
func New(replicas int, hash Hash) (m *Map) {
	m = &Map{
		hash:     hash,
		replicas: replicas,
		keys:     []uint32{},
		hashMap:  make(map[uint32]string),
	}
	// 默认使用crc32来计算hash
	if hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return
}

// IsEmpty 检查是否为空
func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// Add 添加节点
func (m *Map) Add(keys ...string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := m.hash([]byte(strconv.Itoa(i) + key))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Sort(m.keys)
}

// Delete 删除节点
func (m *Map) Delete(keys ...string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := m.hash([]byte(strconv.Itoa(i) + key))
			delete(m.hashMap, hash)
		}
	}
	m.updateKeys()
}

// 更新排序hash
func (m *Map) updateKeys() {
	keyLen := len(m.hashMap)
	keys := make([]uint32, keyLen)
	for hash := range m.hashMap {
		keys = append(keys, hash)
	}
	k := uints(keys)
	sort.Sort(k)
	m.keys = k

}

// Get 根据key获取对应的节点信息
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}
	hash := m.hash([]byte(key))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	if idx == len(m.keys) {
		idx = 0
	}
	return m.hashMap[m.keys[idx]]
}
