func twoSum(nums []int, target int) []int {
    i := 0;
    j   := i+1;
    n := len(nums);


for i <n {
    j = i + 1

for  j <n{
         
        if nums[i] + nums[j] == target  {
      return []int{i, j}
        }else {
         j++
        }
}
    i++
}
return []int{}
}

