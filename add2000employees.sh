#!/bin/bash

for ((i = 0; i<200; i++))
do
    ./upload-employee-records.sh
    sleep 1
done
