package twopointers

import "sort"

func threeSumDefault(nums []int) [][]int {
    vals := [][]int{}
    n := len(nums)
    for i :=0 ; i< n; i++{
        for j :=i+1; j< n ;j++ {
            for k := j+ 1; k < n ;k++{
                if nums[i] + nums[j] + nums [k] ==0 {
                    val := []int {nums[i],nums[j],nums[k]}
                    alreadyExists := isExists(val,vals)
                    if !alreadyExists {
                        vals = append(vals,val)
                    }    
                }
            }
        }
    }
    return vals
}

func isExists(val []int,vals [][]int) bool {
    for _, v := range vals {
        valCopy := append([]int(nil), val...)
        vCopy := append([]int(nil), v...)
        sort.Ints(valCopy)
        sort.Ints(vCopy)
        equal := true
        for i := range valCopy {
            if valCopy[i] != vCopy[i] {
                equal = false
                break
            }
        }
        if equal {
            return true
        }
    }
    return false
}

func threeSumOptimized(nums []int) [][]int {

    sort.Ints(nums)
    values := [][]int{}
    n := len(nums)
    
    for i := 0 ;i < n-2; i++ {
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        index := i
        value := nums[i]
        left := index + 1
        right := n - 1
        for left < right {
            leftRightSum := nums[left] + nums[right]
            if leftRightSum == (-1) * value{
                val := []int {nums[index], nums[left], nums[right]}
                values = append(values,val)
                
                for left < right && nums[left] == nums[left+1]{
                    left++
                }
                for left < right && nums[right] == nums[right-1]{
                    right--
                }
                left++
                right--
            } else if leftRightSum > (-1) * value{
                right = right - 1
            } else {
                left = left + 1
            }
        }
    }
    return values
}