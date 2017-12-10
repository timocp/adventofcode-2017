#! /bin/sh

go test && go build || exit 1
for d in $(jot 8); do
    if [ -f input/day$d.in ]; then
        input=$(cat input/day$d.in)
    else
        input=input/day$d.txt
    fi
    echo -n "Day $d first puzzle  => "
    ./adventofcode ${d}a $input
    echo -n "Day $d second puzzle => "
    ./adventofcode ${d}b $input
done
