#!/bin/bash

# Bash will expand * and make a list for us
for sc in Screenshot*; do
    echo Opening "$sc" with Eye of GNOME
    eog "$sc"
    echo -n "Enter a new filename for $sc, or type 'q' to quit or 'd' to delete: "
    read dest

    if [[ "$dest" == "q" ]]; then
        exit 0
    elif [[ "$dest" == "d" ]]; then
        rm -v "$sc"
    else
        mv -v "$sc" ./screenshots/$dest.png
    fi

    echo

done
