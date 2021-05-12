package main
import (
  
  "fmt"
 )
 
func abs(x int) int {
  y :=0
  if x<0 {
   y=-x
  } else {
   y =x
  }
  return y
}

func mFor(a,b,bz []int, d,fs int) {
  var totalA=1
  var totalB=1
  var n = make([]int, len(a))
  var tempArry =make([]int, len(a))
  var tempSum1 int
  tempSum2 := 10000000
  lastSum :=0
  at := copy(n,a)
  at += copy(tempArry,a)
  for t:=0;t<=d;t++ {
	  totalA = totalA*(b[t]-a[t]+1)
  }
  
  for t:=d+1;t<len(a);t++ {
	  totalB = totalB*(b[t]-a[t]+1)
      n[t]=a[t]
      tempArry[t]=a[t]
  }
	
	for w:=0;w<totalA;w++ {        
		for t:=d;t>=0;t-- {
		    if n[t]>b[t] {
			  n[t-1]++
			  n[t]=a[t]
			}
	    }
		for v:=d+1;v<totalB;v++ {
			for t:=len(a)-1;t>d;t-- {
		      if(n[t]>b[t]) {
			    n[t-1]++
			    n[t]=a[t]
			  }
	        }
			
			for k:=0;k<=d;k++ {
			  lastSum = lastSum + n[k]*bz[k]
			}
			for k:=d+1;k<(d+len(a)-1)/2+1;k++ {
			  lastSum=lastSum + n[k]*n[k+(len(a)-d-1)/2]*bz[k]		
			}
			tempSum1 = abs(lastSum - fs)
			if ((lastSum - fs)>tempSum2) {
				for t:=d+1;t<len(a);t++ {	       
                  n[t]=a[t]
	            }
				
				break
			}
			if (tempSum2 > tempSum1) {
				tempSum2 = tempSum1
				for m:=0;m<len(a);m++{
				tempArry[m]=n[m]
				}
			}
			n[len(a)-1]++
			lastSum = 0
	    }
		n[d]++
	}
	fmt.Println(tempArry)
	fmt.Println(tempSum2)
	fmt.Println(lastSum)
	fmt.Println(fs)
	
}


func main() {
 var maxArr = []int{2600,2400,980,2455,5,5}
 var minArr = []int{1800,1500,456,1922,1,1}
 var perQuanArr = []int{16000,1000,480,200}
 var h=1
 var firstSum = 37180000
 mFor(minArr,maxArr,perQuanArr,h,firstSum)
	
}

