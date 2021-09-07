func threeSum(nums []int) [][]int {
	// create 2d slice to return values in
	ret := make([][]int, 0)
	
	// sort the input array.  This makes a hashmap unnecessary, as we can fast forward over duplicate
	// values that are next to each other.  sort.Ints works better than sort.Slice
	sort.Ints(nums)
	
	// i will traverse the whole array, we fast forward over contigious duplicate values in nums
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		
		// In the balance of the array not referenced by i, we keep an index at the left most
		// and right most elements.
		for l, r := i+1, len(nums)-1; l < r; {
			if nums[l]+nums[r] == -nums[i] {  // this is a match where the 3sum == 0
				// store the match
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				
				// decrementR will move r one place to the left, and then continue until it sees a new value
				r = decrementR(l, r, nums)
				
				// incrementL will advance L until it sees a new value
				l = incrementL(l, r, nums)
			} else if nums[l]+nums[r] > -nums[i] {
				// if the sum is too big, decrementR until we have a new value to check
				r = decrementR(l, r, nums)
			} else if nums[l]+nums[r] < -nums[i] {
				// if the sum is too small, incrementL until we have a new value to check
				l = incrementL(l, r, nums)
			}
		}
	}
	return ret
}

func decrementR(l, r int, nums []int) int {
	r--
	for l < r && nums[r] == nums[r+1] {
		r--
	}
	return r
}

func incrementL(l, r int, nums []int) int {
	l++
	for l < r && nums[l] == nums[l-1] {
		l++
	}
	return l
}

