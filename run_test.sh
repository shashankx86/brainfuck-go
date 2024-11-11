#!/bin/bash

# Define variables
BUILD_OUTPUT="bf-interpreter"
TEST_FILE="test.bf"
DEBUG=${DEBUG:-0}  # Default to 0 if DEBUG is not set

echo "==============================="
echo "Brainfuck Interpreter - Go"
echo "==============================="

# Function to print debug information
debug_echo() {
    if [[ $DEBUG -eq 1 ]]; then
        echo "[DEBUG] $1"
    fi
}

# Check if Go is installed
debug_echo "Checking if Go is installed..."
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    exit 1
fi

# Step 1: Build the project
echo "Step 1: Building the Brainfuck interpreter..."
debug_echo "Running: go build -o $BUILD_OUTPUT"
if go build -o $BUILD_OUTPUT; then
    echo "Build successful!"
else
    echo "Build failed. Please check the error messages."
    exit 1
fi

# Step 2: Check if test file exists
echo "Step 2: Checking for test file '$TEST_FILE'..."
debug_echo "Checking file existence: $TEST_FILE"
if [[ ! -f $TEST_FILE ]]; then
    echo "Error: Test file '$TEST_FILE' not found."
    exit 1
else
    echo "Test file found!"
fi

# Step 3: Run the test file
echo "Step 3: Running the test program..."
if [[ $DEBUG -eq 1 ]]; then
    echo "Running interpreter with debug mode enabled..."
    debug_echo "Running: ./$BUILD_OUTPUT -file=$TEST_FILE"
    ./$BUILD_OUTPUT -file=$TEST_FILE 
else
    debug_echo "Running: ./$BUILD_OUTPUT -file=$TEST_FILE"
    ./$BUILD_OUTPUT -file=$TEST_FILE
fi

if [[ $? -eq 0 ]]; then
    echo ""
    echo "Test run completed successfully!"
else
    echo "Error: Test run failed."
    exit 1
fi

# Cleanup: Optionally remove the build output
echo "Step 4: Cleanup (Optional)"
read -p "Do you want to delete the build output file '$BUILD_OUTPUT'? (y/n): " choice
if [[ "$choice" == "y" || "$choice" == "Y" ]]; then
    debug_echo "Removing build output file: $BUILD_OUTPUT"
    rm -f $BUILD_OUTPUT
    echo "Build output file removed."
else
    echo "Build output file retained."
fi

echo "==============================="
echo "Script execution completed."
echo "==============================="
