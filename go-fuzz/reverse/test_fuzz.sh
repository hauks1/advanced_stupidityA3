#!/bin/bash

for time in {"5s","10s","15s","30s","45s"}
do 
    go test -fuzz=FuzzReverse -fuzztime "${time}" > results/mult_inputs/res_"${time}".txt
done
