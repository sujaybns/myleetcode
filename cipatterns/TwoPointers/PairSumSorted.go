package twopointers

import "fmt"

type pair struct {
	SIndex int
	LIndex int
}

func PairSumDefault(target int, mylist []int) *pair {
	fmt.Printf("\nHello World : target: \n%d in list \n%v", target, mylist)
	for i := range mylist {
		for j := i + 1; j < len(mylist); j++ {
			fmt.Println("\nmylist[i]", mylist[i])
			fmt.Println("mylist[j]", mylist[j])
			if mylist[i]+mylist[j] == target {
				return &pair{i, j}
			}
		}
	}
	return nil
}



func PairSumOptimized(target int, mylist []int) *pair {
	fmt.Printf("\nHello World : target: \n%d in list \n%v", target, mylist)
	left := 0
	right := len(mylist) -1 

	for left < len(mylist) && right > 0 {
		if mylist[left] + mylist[right] > target{
			right = right - 1
		} else if mylist[left] + mylist[right] < target{
			left = left + 1
		} else {
			return &pair{left,right}
		}
	}
	return nil
	/*
	for i := range mylist {
		for j := i + 1; j < len(mylist); j++ {
			fmt.Println("\nmylist[i]", mylist[i])
			fmt.Println("mylist[j]", mylist[j])
			if mylist[i]+mylist[j] == target {
				return &pair{i, j}
			}
		}
	}
	
	return nil
	*/
}