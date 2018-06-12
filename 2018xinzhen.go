package main


import (

	"fmt"

)

func jizhi (q [4]int) (int, int) {
    max := 0
	min := 10
    for _ , v := range q {
        if max < v {
            max = v
		}
		if min > v {
		    min = v
        }
    }
	
    return max, min
}


func checkResult(a [10]int) {
    var i = 1
	var sorted [4]int
	for _, m := range a {
	    switch m {
		    case 1:
			    sorted[0]++
			case 2:
			    sorted[1]++
			case 3:
			    sorted[2]++
			case 4:
			    sorted[3]++
		}
	}
	
	max, min := jizhi(sorted)

	
	switch a[1] {
	    case 1:
		    if a[4] == 3 {i++}
		case 2:
		    if a[4] == 4 {i++}
		case 3:
		    if a[4] == 1 {i++}
		case 4:
		    if a[4] == 2 {i++}
	}
	switch a[2] {
	    case 1:
		    if (a[2] != a[5]) && (a[2] != a[1]) && (a[2] != a[3]){i++}
		case 2:
		    if (a[5] != a[2]) && (a[5] != a[1]) && (a[5] != a[3]){i++}
		case 3:
		    if (a[1] != a[2]) && (a[1] != a[5]) && (a[1] != a[3]){i++}
		case 4:
		    if (a[3] != a[2]) && (a[3] != a[1]) && (a[3] != a[5]){i++}
	}
	
	switch a[3] {
	    case 1:
		    if a[0] == a[4]{i++}
		case 2:
		    if a[1] == a[6]{i++}
		case 3:
		    if a[0] == a[8]{i++}
		case 4:
		    if a[5] == a[9]{i++}
	}
	
	switch a[4] {
	    case 1:
		    if a[7] == 1{i++}
		case 2:
		    if a[3] == 2{i++}
		case 3:
		    if a[8] == 3{i++}
		case 4:
		    if a[6] == 4{i++}
	}
		
	switch a[5] {
	    case 1:
		    if a[7] == a[1] && a[7] == a[3]{i++}
		case 2:
		    if a[7] == a[0] && a[7] == a[5]{i++}
		case 3:
		    if a[7] == a[2] && a[7] == a[9]{i++}
		case 4:
		    if a[7] == a[4] && a[7] == a[8]{i++}
	}
	
	switch a[6] {
	    case 1:
		    if sorted[2] == min{i++}
		case 2:
		    if sorted[1] == min{i++}
		case 3:
		    if sorted[0] == min{i++}
		case 4:
		    if sorted[3] == min{i++}
	}
	
	switch a[7] {
	    case 1:
		    if (a[6]-a[0])>1 || (a[6]-a[0])< -1 {i++}
		case 2:
		    if (a[4]-a[0])>1 || (a[4]-a[0])< -1  {i++}
		case 3:
		    if (a[1]-a[0])>1 || (a[1]-a[0])< -1 {i++}
		case 4:
		    if (a[9]-a[0])>1 || (a[9]-a[0])< -1 {i++}
	}
	
	switch a[8] {
	    case 1:
		    if (a[0]==a[5])!=(a[4]==a[5]){i++}
		case 2:
		    if (a[0]==a[5])!=(a[4]==a[9]){i++}
		case 3:
		    if (a[0]==a[5])!=(a[4]==a[1]){i++}
		case 4:
		    if (a[0]==a[5])!=(a[4]==a[8]){i++}
	}
	
	switch a[9] {
	    case 1:
		    if max-min==3{i++}
		case 2:
		    if max-min==2{i++}
		case 3:
		    if max-min==4{i++}
		case 4:
		    if max-min==1{i++}
	}
	
	if i == 10 {
	    for _, zf := range(a) {
	        switch zf {
	            case 1:
		            fmt.Printf("A")
		        case 2:
		            fmt.Printf("B")
		        case 3:
		            fmt.Printf("C")
		        case 4:
		            fmt.Printf("D")
	        }
		}
	}
}
	



func main() {

	var answer [10]int
	for answer[0] = 1; answer[0]<5; answer[0]++ {
		for answer[1] = 1; answer[1]<5; answer[1]++ {
			for answer[2] = 1; answer[2]<5; answer[2]++ {
				for answer[3] = 1; answer[3]<5; answer[3]++ {
					for answer[4] = 1; answer[4]<5; answer[4]++ {
						for answer[5] = 1; answer[5]<5; answer[5]++ {
							for answer[6] = 1; answer[6]<5; answer[6]++ {
								for answer[7] = 1; answer[7]<5; answer[7]++ {
									for answer[8] = 1; answer[8]<5; answer[8]++ {
										for answer[9] = 1; answer[9]<5; answer[9]++ {
										checkResult(answer)
										
										
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
    
}