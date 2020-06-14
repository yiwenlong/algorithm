package main

import "fmt"

//
//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//进阶:
//
//你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/lru-cache
//
type CacheNode struct {
	Key  int
	Val  int
	Prev *CacheNode
	Next *CacheNode
}

type LRUCache struct {
	KvMap    map[int]*CacheNode
	Root     *CacheNode
	Tail     *CacheNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	r := &CacheNode{}
	t := &CacheNode{}
	r.Next = t
	t.Prev = r
	return LRUCache{
		KvMap:    make(map[int]*CacheNode, capacity),
		Root:     r,
		Tail:     t,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.KvMap[key]
	if !ok {
		return -1
	}
	if this.Root.Next == node {
		return node.Val
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = this.Root.Next
	this.Root.Next.Prev = node
	this.Root.Next = node
	node.Prev = this.Root
	return this.KvMap[key].Val
}

func (this *LRUCache) Put(key int, value int) {
	v := this.Get(key)
	if v != -1 {
		if v != value {
			this.KvMap[key].Val = value
		}
		return
	}
	if this.capacity == len(this.KvMap) {
		delete(this.KvMap, this.Tail.Prev.Key)
		this.Tail.Prev = this.Tail.Prev.Prev
		this.Tail.Prev.Next = this.Tail
	}
	node := &CacheNode{
		Key:  key,
		Val:  value,
		Prev: this.Root,
		Next: this.Root.Next,
	}
	this.KvMap[key] = node
	this.Root.Next.Prev = node
	this.Root.Next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
