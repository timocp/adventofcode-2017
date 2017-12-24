#! /bin/sh

go test -v ./... && go build || exit 1
for d in $(seq 1 24); do
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
