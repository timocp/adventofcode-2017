#! /bin/sh

go build || exit
./adventofcode 1a $(cat input/day1.txt)
./adventofcode 1b $(cat input/day1.txt)
./adventofcode 2a input/day2.txt
