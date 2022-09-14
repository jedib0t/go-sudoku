#!/bin/bash
cd $(dirname $0)/..

# ensure the binary is available to run
if [[ ! -e ./go-sudoku ]]; then
  echo "ERROR: ./go-sudoku is missing; try 'make build' first?"
  exit 1
fi

# setup seed value to make the results reproducible
export SEED=42

# read the template line by line and execute commands if any
while read line
do
  echo $line
  # if the line contains a command to be executed, execute it
  if [[ $line =~ ^\$\ (.*)\|(.*)$ ]]; then
    # commands with a single pipe
    (${BASH_REMATCH[1]} | ${BASH_REMATCH[2]}) 2>&1
  elif [[ $line =~ ^\$\ (.*)$ ]]; then
    # simple command without a pipe
    (${BASH_REMATCH[1]}) 2>&1
  fi
done <docs/README.md.template
