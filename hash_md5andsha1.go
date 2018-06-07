package main
//使用md5和sha1加密字符串
import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	TestString := "Hi, pandaman!" //要加密的字符串

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
	
	TestFile := "a.txt"  //要加密的文件
	infile, err := os.Open(TestFile)
	if err == nil {
	    md5f := md5.New()
		io.Copy(md5f, infile)
		fmt.Printf("%x %s\n", md5f.Sum([]byte("")), TestFile)
		
	    sha1f := sha1.New()
		io.Copy(sha1h, infile)
		fmt.Printf("%x %s\n", sha1h.Sum([]byte("")), TestFile)
    } else {
	    fmt.Println(err)
		os.Exit(1)
	}
	
}
