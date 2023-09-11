package orhash

import "math/rand"

type RandomizedSet struct {
	NumsToLoc map[int]int
	Nums      []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		NumsToLoc: make(map[int]int),
		Nums:      make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.NumsToLoc[val]; ok {
		return false
	}

	// insert val to the end of 'nums'
	this.NumsToLoc[val] = len(this.Nums)
	this.Nums = append(this.Nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.NumsToLoc[val]; !ok {
		return false
	}

	nums := this.Nums
	/** 删除元素 */
	loc := this.NumsToLoc[val]
	// 和最后一个元素交换后删除
	lastEle := nums[len(nums)-1]
	nums[loc] = lastEle
	this.NumsToLoc[lastEle] = loc
	delete(this.NumsToLoc, val)
	this.Nums = nums[:len(nums)-1]

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(len(this.Nums))
	return this.Nums[randIdx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
