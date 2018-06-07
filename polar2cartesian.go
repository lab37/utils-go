package main
//极坐标转为笛卡尔坐标
//用法 polar2cartesian 回车，然后输入半径空格角度回车，角度可以是负数
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

const result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, or %s to quit."

type polar struct {
	radius float64
	θ      float64
}

type cartesian struct {
	x float64
	y float64
}

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)

	answers := createSolver(questions)
	defer close(answers)

	interact(questions, answers)
}

// begins by creating an answers channel to which it will send the answers (i.e., cartesian coordinates)
// to the questions (i.e., polar coordinates) that it receives fromt he questions channel.
// The go statemet has an infinite loop that waits (blocking its own goroutine, but not any other goroutines,
// and not the function in which the goroutine was started), until it receives a question. When a polar coordinate
// arrives the anonymous function computes the corresponding cartesian coordinates using some simple math, and
// then send the answer as a cartesian struct to the answers channel.
// Once the call to createSolver() returns we have reached the point where we have two communication channels set
// up and where a separate goroutine is waiting for polar coordinates to be sent on the questions channel - and
// without any other goroutine, including the one executing main(), being blocked.
func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)

	go func() {
		for {
			polarCoord := <-questions
			θ := polarCoord.θ * math.Pi / 180.0 // degrees to radians
			x := polarCoord.radius * math.Cos(θ)
			y := polarCoord.radius * math.Sin(θ)
			answers <- cartesian{x, y}
		}
	}()

	return answers
}

// If valid numbers were input and sent to the questions channel (in a polar struct), we block the main
// goroutine waiting for a response on the answers channel. The additional goroutine created in createSolver()
// function is itself blocked waiting for a polar on the questions channel, so when we send the polar, the
// additional goroutine performs the computation, sends the resultant cartesian to the answers channel, and
// then waits (blocking only itself) for another question to arrive. Once the cartesian answer is received in
// the interact() function on the answers channel, interact() is no longer blocked. At this point we print the
// result.
func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)

	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		var radius, θ float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &θ); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}

		questions <- polar{radius, θ}
		coord := <-answers
		fmt.Printf(result, radius, θ, coord.x, coord.y)
	}

	fmt.Println()
}
