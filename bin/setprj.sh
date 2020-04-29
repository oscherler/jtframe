#!/bin/bash
# Define JTROOT before sourcing this file

export JTFRAME=$JTROOT/modules/jtframe
PATH=$PATH:$JTFRAME/bin:.
#unalias jtcore
alias jtcore="$JTFRAME/bin/jtcore"

# derived variables
if [ -e $JTROOT/cores ]; then
    CORES=$JTROOT/cores
else
    CORES=$JTROOT
fi

export ROM=$JTROOT/rom
CC=$JTROOT/cc
MRA=$ROM/mra
export MODULES=$JTROOT/modules
JT12=$MODULES/jt12
JT51=$MODULES/jt51

function swcore {
    IFS=/ read -ra string <<< $(pwd)
    j="/"
    next=0
    good=
    for i in ${string[@]};do
        if [ $next = 0 ]; then
            j=${j}${i}/            
        else
            next=0
            j=${j}$1/
        fi
        if [ "$i" = cores ]; then
            next=1
            good=1
        fi
    done
    if [[ $good && -d $j ]]; then
        cd $j
    else       
        cd $JTROOT/cores/$1
    fi
    pwd
}

if [ "$1" != "--quiet" ]; then
    echo "Use swcore <corename> to switch to a different core once you are"
    echo "inside the cores folder"
fi