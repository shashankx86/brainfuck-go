# Brainfuck Interpreter in Go

A Brainfuck interpreter written in Go with additional features like a customizable memory size and optional debug mode.

## Features

- **Standard Brainfuck Commands**: Supports the full set of Brainfuck commands:
  - `>` : Increment the memory pointer.
  - `<` : Decrement the memory pointer.
  - `+` : Increment the byte at the memory pointer.
  - `-` : Decrement the byte at the memory pointer.
  - `.` : Output the byte at the memory pointer as a character.
  - `,` : Read a byte of input into the byte at the memory pointer.
  - `[` : Jump forward to the matching `]` if the byte at the memory pointer is zero.
  - `]` : Jump back to the matching `[` if the byte at the memory pointer is non-zero.
  - `#` : (Extra) Print debug information at the current state if debug mode is enabled.
- **Debug Mode**: Enables step-by-step execution with memory dumps and pointer tracking.
- **Customizable Memory Size**: Specify the memory size for the Brainfuck tape via a command-line argument.
- **Error Handling**: Includes checks for unmatched brackets and memory pointer overflows/underflows.

## Prerequisites

- [Go](https://go.dev/dl/) 1.18 or higher

## Installation

Clone the repository:

```bash
git clone https://github.com/shashankx86/brainfuck-go.git
cd brainfuck-go
```

Build the project:

```bash
go build -o bf-interpreter
```

## Usage

```bash
./bf-interpreter -file=<path_to_code.bf> [-input=<input_string>] [-memory=<memory_size>] [-debug]
```

- `-file` (required): Path to the Brainfuck file to execute (e.g., `test.bf`).
- `-input` (optional): Input string for the Brainfuck `,` command.
- `-memory` (optional): Size of the memory tape (default is `30000`).
- `-debug` (optional): Enable debug mode to print the interpreter state after each command.

### Example

1. Create a sample Brainfuck file `test.bf`:

    ```brainfuck
    ++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
    ```

2. Run the interpreter:

    ```bash
    ./bf-interpreter -file=test.bf
    ```

    **Output**:

    ```
    Hello World!
    ```

3. Run with debug mode enabled:

    ```bash
    ./bf-interpreter -file=test.bf -debug
    ```

    This will display detailed debug information about the interpreter's internal state during execution.

## Project Structure

```
brainfuck-go/
├── go.mod              # Go module file
├── main.go             # Main program containing the interpreter logic
├── README.md           # Project documentation
└── test.bf             # Example Brainfuck program
```

## Customizing and Testing

### Building From Source

You can build the project with:

```bash
go build
```

### Running Tests

You can test the interpreter using a sample `.bf` file like `test.bf`:

```bash
go run main.go -file=test.bf
```

### Sample Brainfuck Code

Here are some sample programs you can try:

1. **Hello World**:
    ```brainfuck
    ++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
    ```
2. **Multiply two numbers (input: 2 and 3)**:
    ```brainfuck
    ,>,<[>[->+>+<<]>>[-<<+>>]<<<-]>>.
    ```

For the multiplication program, you can run:

```bash
go run main.go -file=<your_file>.bf -input="\x02\x03"
```

## Contributing

Contributions are welcome! Feel free to fork the repository, open issues, and submit pull requests. For major changes, please open an issue first to discuss your ideas.

## License

This project is licensed under the MIT License

## Acknowledgments

- Inspired by the [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) programming language.
