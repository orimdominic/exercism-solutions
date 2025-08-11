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
  main () {
    NUM=$1
    local RESULT=""

    if [ $(($NUM % 3)) -eq 0 ]; then
      RESULT="${RESULT}Pling"
    fi

    if [ $(($NUM % 5)) -eq 0 ]; then
      RESULT="${RESULT}Plang"
    fi

    if [ $(($NUM % 7)) -eq 0 ]; then
      RESULT="${RESULT}Plong"
    fi

    if [ "$RESULT" == "" ] ; then
      echo "$NUM"
      exit 0
    fi

    echo $RESULT
  }
#
#   # call main with all of the positional arguments
  main "$@"
#
# *** PLEASE REMOVE THESE COMMENTS BEFORE SUBMITTING YOUR SOLUTION ***
