#!/bin/bash

. set-env.sh acme



# jq -n --slurpfile arr ./convertcsv.json '$arr'
# length=$( cat convertcsv.json | jq length)
# arr=$( cat convertcsv.json | jq '.[0]')

# Read the elements in an array
arr=$(cat employee-records.json | jq -c '.[]')

# Document type for the data
docType="\"EmployeeRecords\""
RANDOM=$$
# echo $arr
COUNTER=1
for item in $arr;
do

  id=$(echo $item | jq .employeeId)
  randomId=$(echo $RANDOM)
  employeeId=$(($id + $randomId))
  employeeName=$(echo $item | jq .employeeName)
  employeeType=$(echo $item | jq .employeeType)
  parentCompany=$(echo $item | jq .parentCompany)
  currentCompany=$(echo $item | jq .currentCompany)
  salary=$(echo $item | jq .salary)
  workingHours=$(echo $item | jq .workingHours)

  echo "$COUNTER  $employeeName"
  COUNTER=$((COUNTER+1))

  args="{\"Args\":[\"createEmployee\",\"$employeeId\", $employeeName,\"$salary\",\"$workingHours\",$employeeType,$parentCompany,$currentCompany]}"
  set-chain-env.sh -i  "$args"
  chain.sh invoke

  echo $args

done
