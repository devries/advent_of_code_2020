#!/bin/sh

for day in $(seq -f "%02g" 1 25); do
  finecho=0
  for part in 1 2; do
    dir="day${day}_p${part}"
    if [ -d $dir ]; then
      if [ $part -eq 1 ]; then
        printf "Day $day "
        finecho=1
      fi
      cwd=$(pwd)
      cd $dir
      printf "part $part: "
      retval="$(./$dir)"
      printf "$retval "
      cd $cwd
    fi
  done
  if [ $finecho -eq 1 ]; then
    echo ""
  fi
done
