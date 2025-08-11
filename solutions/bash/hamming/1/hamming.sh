#!/usr/bin/env bash

# The following comments should help you get started:
# - Bash is flexible. You may use functions or write a "raw" script.
#
# - Complex code can be made easier to read by breaking it up
#   into functions, however this is sometimes overkill in bash.
#
# - You can find links about good style and other resources
#   for Bash in './README.md'. It came with this exercise.
#
#   Example:
#   # other functions here
#   # ...
#   # ...
#
main() {
  if [ $# -lt 2 ]; then
    echo "Usage: hamming.sh <string1> <string2>"
    exit 1
  fi

  START_STRING=$1
  END_STRING=$2

  if [ ${#START_STRING} -eq 0 ] && [ ${#END_STRING} -eq 0 ]; then
    echo "0"
    exit 0
  fi

  if [ ${#START_STRING} -ne ${#END_STRING} ]; then
    echo "error: strands must be of equal length"
    exit 1
  fi

  DIFF=0
  for ((i = 0; i < ${#START_STRING}; i++)); do
    if [ ${START_STRING:$i:1} != ${END_STRING:$i:1} ]; then
      DIFF=$(expr $DIFF + 1)
    fi
  done

  echo "$DIFF"
}
#
#   # call main with all of the positional arguments
main "$@"
#
# *** PLEASE REMOVE THESE COMMENTS BEFORE SUBMITTING YOUR SOLUTION ***
