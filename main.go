package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	// "strconv"
	// "strings"
	"io/ioutil"
)

// Interpreter structure to manage the Brainfuck state
type Interpreter struct {
	code        string
	input       string
	memory      []byte
	ptr         int
	codePtr     int
	inputPtr    int
	loopMap     map[int]int
	debugMode   bool
	memorySize  int
}

// NewInterpreter initializes a new Brainfuck interpreter
func NewInterpreter(code, input string, memorySize int, debugMode bool) *Interpreter {
	return &Interpreter{
		code:       code,
		input:      input,
		memory:     make([]byte, memorySize),
		loopMap:    make(map[int]int),
		debugMode:  debugMode,
		memorySize: memorySize,
	}
}

// BuildLoopMap preprocesses the code to map loop brackets for faster execution
func (bf *Interpreter) BuildLoopMap() error {
	stack := []int{}
	for i, cmd := range bf.code {
		switch cmd {
		case '[':
			stack = append(stack, i)
		case ']':
			if len(stack) == 0 {
				return fmt.Errorf("unmatched closing bracket at position %d", i)
			}
			start := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			bf.loopMap[start] = i
			bf.loopMap[i] = start
		}
	}
	if len(stack) > 0 {
		return fmt.Errorf("unmatched opening bracket at position %d", stack[0])
	}
	return nil
}

// DebugPrint prints the current state of the interpreter for debugging
func (bf *Interpreter) DebugPrint() {
	fmt.Printf("\n-- Debug State --\n")
	fmt.Printf("Code Pointer: %d\n", bf.codePtr)
	fmt.Printf("Memory Pointer: %d\n", bf.ptr)
	fmt.Printf("Current Cell Value: %d ('%c')\n", bf.memory[bf.ptr], bf.memory[bf.ptr])
	fmt.Printf("Memory Dump (0-15): %v\n", bf.memory[:16])
	fmt.Printf("Loop Map: %v\n", bf.loopMap)
	fmt.Println("------------------\n")
}

// Execute runs the Brainfuck program
func (bf *Interpreter) Execute() error {
	for bf.codePtr < len(bf.code) {
		cmd := bf.code[bf.codePtr]
		switch cmd {
		case '>':
			bf.ptr++
			if bf.ptr >= bf.memorySize {
				return fmt.Errorf("memory pointer overflow at position %d", bf.ptr)
			}
		case '<':
			bf.ptr--
			if bf.ptr < 0 {
				return fmt.Errorf("memory pointer underflow at position %d", bf.ptr)
			}
		case '+':
			bf.memory[bf.ptr]++
		case '-':
			bf.memory[bf.ptr]--
		case '.':
			fmt.Printf("%c", bf.memory[bf.ptr])
		case ',':
			if bf.inputPtr < len(bf.input) {
				bf.memory[bf.ptr] = bf.input[bf.inputPtr]
				bf.inputPtr++
			} else {
				bf.memory[bf.ptr] = 0
			}
		case '[':
			if bf.memory[bf.ptr] == 0 {
				bf.codePtr = bf.loopMap[bf.codePtr]
			}
		case ']':
			if bf.memory[bf.ptr] != 0 {
				bf.codePtr = bf.loopMap[bf.codePtr]
			}
		case '#':
			bf.DebugPrint()
		default:
			// Ignore any non-Brainfuck characters (comments or invalid characters)
		}

		bf.codePtr++

		if bf.debugMode {
			bf.DebugPrint()
			fmt.Print("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
	return nil
}

// LoadCodeFromFile reads the Brainfuck code from a file
func LoadCodeFromFile(filePath string) (string, error) {
	codeBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}
	return string(codeBytes), nil
}

func main() {
	// Define command-line arguments
	filePath := flag.String("file", "", "Path to Brainfuck (.bf) file")
	inputString := flag.String("input", "", "Input string for ',' commands")
	memorySize := flag.Int("memory", 30000, "Memory size (default is 30000)")
	debugMode := flag.Bool("debug", false, "Enable debug mode (default is false)")

	// Parse command-line arguments
	flag.Parse()

	// Validate required arguments
	if *filePath == "" {
		fmt.Println("Error: Brainfuck file path is required. Use -file=<path_to_code.bf>")
		return
	}

	// Load Brainfuck code from file
	code, err := LoadCodeFromFile(*filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create and initialize the interpreter
	interpreter := NewInterpreter(code, *inputString, *memorySize, *debugMode)

	// Build loop mappings
	if err := interpreter.BuildLoopMap(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Execute the Brainfuck code
	if err := interpreter.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
