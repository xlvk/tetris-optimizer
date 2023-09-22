#!/bin/bash

echo "Building the Go project..."
cd /home/xlvk/reboot01/tetris-optimizer  # Change to the correct directory
go build -o tetris-optimiser 
if [ $? -ne 0 ]; then
    echo "Error building the project. Exiting."
    exit 1
fi

echo "Testing against input files..."
for file in ./TestFile/*; do
    echo "Testing with input file: $file"
    
    TIME_TAKEN=$( { time ./tetris-optimiser "$file" ; } 2>&1 | grep real | awk '{print $2}')
    output=$(./tetris-optimiser "$file")
    echo "$output"
    echo "----------------------------------"
    
    spaces_count=$(echo "$output" | grep -o '\.' | wc -l)
    echo "Number of spaces in output: $spaces_count"
    echo "Time taken: $TIME_TAKEN"
    echo "----------------------------------"
done

echo "All tests completed."