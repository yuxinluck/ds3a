package main

/*
146. LRU 缓存
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type LRUCache struct {
	Cap  int
	Vmap map[int]*Node
	Head *Node
	Tail *Node
}

type Node struct {
	Val  int
	Key  int
	Next *Node
	Pre  *Node
}

func delNode(n *Node) {
	n.Pre.Next, n.Next.Pre = n.Next, n.Pre
}

func addNode(pos *Node, n *Node) {
	n.Pre = pos
	n.Next = pos.Next
	pos.Next.Pre = n
	pos.Next = n
}

func Constructor(capacity int) LRUCache {

	ProtectHead := &Node{
		Val: 0,
	}
	ProtectTail := &Node{
		Val: 0,
	}
	ProtectHead.Next = ProtectTail
	ProtectTail.Pre = ProtectHead
	vMap := make(map[int]*Node, 0)
	return LRUCache{Cap: capacity, Vmap: vMap, Head: ProtectHead, Tail: ProtectTail}
}

func (this *LRUCache) Get(key int) int {
    //如果有缓存，返回值，更新到队头
	if n, ok := this.Vmap[key]; ok {
		delNode(n)
		addNode(this.Head, n)
		return n.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {

	//如果已经存在，更新值，位置更新到队头
	if n, ok := this.Vmap[key]; ok {
		n.Val = value
		delNode(n)
		addNode(this.Head, n)
		return
	}

	//新建node，插入到队头
	n := &Node{
		Val: value,
		Key: key,
	}
	this.Vmap[key] = n
	addNode(this.Head, n)

	//插入后如果超出容量，删除最旧的
	if len(this.Vmap) > this.Cap {
		d := this.Tail.Pre.Key
		delNode(this.Tail.Pre)
		delete(this.Vmap, d)
	}
}

func main() {

}
