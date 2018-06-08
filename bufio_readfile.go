package main
import (
    "fmt"
	"strings"
	"bufio"
)

func main() {
s := strings.NewReader("ABCDEFG")
br := bufio.NewReader(s) //返回一个带缓冲区的reader，还没开始读，只是指向了s

b, _ := br.Peek(5) //读入5个
fmt.Printf("%s\n", b)
// ABCDE

b[0] = 'a'
b, _ = br.Peek(5)
fmt.Printf("%s\n", b)
// aBCDE

}




func main() {
s := strings.NewReader("ABCDEFG")
br := bufio.NewReader(s)

c, _ := br.ReadByte() //调用一次读一个字节
fmt.Printf("%c\n", c)
// A

c, _ = br.ReadByte()
fmt.Printf("%c\n", c)
// B

br.UnreadByte()
c, _ = br.ReadByte()
fmt.Printf("%c\n", c)
// B
}





func main() {
s := strings.NewReader("你好，世界！")
br := bufio.NewReader(s)

c, size, _ := br.ReadRune() //按UTF8字符往外读
fmt.Printf("%c %v\n", c, size)
// 你 3

c, size, _ = br.ReadRune()
fmt.Printf("%c %v\n", c, size)
// 好 3

br.UnreadRune()
c, size, _ = br.ReadRune()
fmt.Printf("%c %v\n", c, size)
// 好 3
}





func main() {
s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
br := bufio.NewReader(s)
b := make([]byte, 20)

n, err := br.Read(b) //缓冲区大小为20，一次只读入20个
fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
// ABCDEFGHIJKLMNOPQRST 20 <nil>

n, err = br.Read(b)  //这次再读的时候再读20个，不句的话有多少读多少
fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
// UVWXYZ1234567890 16 <nil>

n, err = br.Read(b)  //读完再调用一次才返回EOF
fmt.Printf("%-20s %-2v %v\n", b[:n], n, err)
// 0 EOF
}



func main() {
s := strings.NewReader("ABC DEF GHI JKL")
br := bufio.NewReader(s)

w, _ := br.ReadSlice(' ') //按指定分隔符读取
fmt.Printf("%q\n", w)
// "ABC "

w, _ = br.ReadSlice(' ')
fmt.Printf("%q\n", w)
// "DEF "

w, _ = br.ReadSlice(' ')
fmt.Printf("%q\n", w)
// "GHI "
}



func main() {
s := strings.NewReader("ABC\nDEF\r\nGHI\r\nJKL")
br := bufio.NewReader(s)

w, isPrefix, _ := br.ReadLine()  //按行读取，其实就是按\n或者\r\n来分隔读取，返回的是切片引用
fmt.Printf("%q %v\n", w, isPrefix) //如果在缓存中找不到行尾标记，则设置isPrefix为 true，表示查找未完成
// "ABC" false

w, isPrefix, _ = br.ReadLine()
fmt.Printf("%q %v\n", w, isPrefix)
// "DEF" false

w, isPrefix, _ = br.ReadLine()
fmt.Printf("%q %v\n", w, isPrefix)
// "GHI" false
}

func main() {
  //原型func (b *Writer) Write(p []byte) (nn int, err error)

  // WriteString 同 Write，只不过写入的是字符串
func (b *Writer) WriteString(s string) (int, error)
    b := bytes.NewBuffer(make([]byte, 0))
    bw := bufio.NewWriter(b)
    fmt.Println(bw.Available()) // 4096
    fmt.Println(bw.Buffered())  // 0

    bw.WriteString("ABCDEFGH")
    fmt.Println(bw.Available()) // 4088
    fmt.Println(bw.Buffered())  // 8
    fmt.Printf("%q\n", b)       // ""

    bw.Flush()
    fmt.Println(bw.Available()) // 4096
    fmt.Println(bw.Buffered())  // 0
    fmt.Printf("%q\n", b)       // "ABCEFG"
}




