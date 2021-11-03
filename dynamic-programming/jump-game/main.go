package main

func main() {

	//var hasSubstring = isSubsequence("abx", "ahbgdc")

	var res = jump([]int{2, 3, 1, 1, 4})
	print(res)

	//fmt.Println(hasSubstring)
}

func jump(nums []int) int {
	var numsCount = len(nums)
	if numsCount <= 1 {
		return 0
	}
	var res = make([]int, numsCount)
	res[0] = 0
	for i:=0; i < numsCount; i++ {
		for j:=1; j <= nums[i] && i + j < numsCount; j++ {
			if res[i+j] == 0 {
				res[i+j] = res[i] + 1
			}
			res[i+j] = minInt(res[i + j], res[i] + 1)
		}
	}
	return res[numsCount-1]
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	var jumpSize = 0
	for i:=0; i < len(nums); i++ {
		jumpSize = maxInt(jumpSize, nums[i])
		if jumpSize == 0 {
			return false
		}
		jumpSize--
	}
	return true
}

func minInt(a int, b int) int  {
	if a > b {
		return b
	}
	return a
}

func maxInt(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
