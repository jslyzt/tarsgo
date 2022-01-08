#!/bin/sh

for file in `ls *.tars`;do
    #sfile=${file%%.*}
    #sfile=$(echo $sfile | tr 'A-Z' 'a-z')
    #echo $sfile $file
    #tars2go --outdir=$sfile $file
    tars2go $file
done