# Day 1 
## Part 1
Part one of day one was all about iterating an array of values. You need to find 2 entries that sum up to 2020, then multiply them. The dataset was quite small so I decided to just brute force iterate the array. First create an array from the input, read the file as a string and split by new line. Then iterate!
### Loop 1
Grab the current entry, make it a number from a string. save the number and start the next loop!
### Loop 2
This loop starts at the next index, grab the current entry, make it a number, add it to the previous number. If it's 2020, multiply the numbers and retuen the result! Boom, done!

## Part2
Part two was the same thing except now there are 3 numbers that equal 2020. So all I really did was add a third loop b