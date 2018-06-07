package main

import (
	"fmt"
	"math/big"  //提供了两个无限精度的整数类型big.Int大整数，和big.Rat大有理数
	"os"
	"path/filepath"
	"strconv"
)

// Pi with Machin's Formula
// pi = 4 X (4 X arccot(5) - arccot(239))
// arcot(x) = 1/x - 1/3x^3 + 1/5x^5 - 1/7x^7 + ...
func main() {
	places := handleCommandLine(1000)  //默认小数位为1000位
	scaledPi := fmt.Sprint(π(places))
	fmt.Printf("3.%s\n", scaledPi[1:])
}

func handleCommandLine(defaultValue int) int {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			usage := "usage: %s [digits]\n e.g.: %s 10000"
			app := filepath.Base(os.Args[0])
			fmt.Fprintln(os.Stderr, fmt.Sprintf(usage, app, app))
			os.Exit(1)
		}

		if x, err := strconv.Atoi(os.Args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "ignoring invalid number of digits: will display %d\n", defaultValue)
		} else {
			return x
		}
	}

	return defaultValue
}

func π(places int) *big.Int {
    //创建4个大整数，返回的是指针其实
	digits := big.NewInt(int64(places)) //初始化为places的值
	unity := big.NewInt(0)  //初始化为0
	ten := big.NewInt(10)
	exponent := big.NewInt(0)
	
	//unity变成10的digits+10次方，注意Exp的三个参数，如果最后一个不是nil，假设是Exp(a,b,c)的话，结果是a的b次方再模z
	unity.Exp(ten, exponent.Add(digits, ten), nil)
	pi := big.NewInt(4)
	left := arccot(big.NewInt(5), unity)
	left.Mul(left, big.NewInt(4))
	right := arccot(big.NewInt(239), unity)
	left.Sub(left, right)
	pi.Mul(pi, left)
	return pi.Div(pi, big.NewInt(0).Exp(ten, ten, nil)) //把pi除以10的10次方，还原为正常的数量级
}

func arccot(x, unity *big.Int) *big.Int {
	sum := big.NewInt(0)
	sum.Div(unity, x)
	xpower := big.NewInt(0)
	xpower.Div(unity, x)
	n := big.NewInt(3)
	sign := big.NewInt(-1)
	zero := big.NewInt(0)
	square := big.NewInt(0)
	square.Mul(x, x)
	for {
		xpower.Div(xpower, square)
		term := big.NewInt(0)
		term.Div(xpower, n)
		if term.Cmp(zero) == 0 {
			break
		}
		addend := big.NewInt(0)
		sum.Add(sum, addend.Mul(sign, term))
		sign.Neg(sign)
		n.Add(n, big.NewInt(2))
	}
	return sum
}
