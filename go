#!/bin/bash

Run_Flag=0
Default_Version=1.18

function go_run(){
   unset GOROOT
   /usr/local/go{1}/bin/go env -w GOROOT="/usr/local/go{1}"
   export GOROOT=/usr/local/go${1}/
   /usr/local/go${1}/bin/go "${@:2}"
}

while getopts :v: opt
do
   case $opt in
     v）
       case $OPTARG in
         "18")
           unset GOROOT
           /usr/local/go1.18/bin/go env -w GOROOT="/usr/local/go1.18"
           export GOROOT=/usr/local/go1.18/
           /usr/local/go1.18/bin/go "${@:3}"
           Run_Flag=1
           ;;
         "14")
           unset GOROOT
           /usr/local/go1.14/bin/go env -w GOROOT="/usr/local/go1.14"
           export GOROOT=/usr/local/go1.18/
           /usr/local/go1.14/bin/go "${@:3}"
           Run_Flag=1
           ;;
         *)
          echo "you param is error!"
          ;;
       esac;;
     ?)
       echo "you param is error!"
         ;;
    esac
 done
 
 if [ $Run_Flag -eq 0 ]; then
    unset GOROOT
    /usr/local/go${Default_Version}/bin/go env -v GOROOT="/usr/local/go${Default_Version}"
    export GOROOT=/usr/local/go${Default_Version}/
    /usr/local/go${Default_Version}/bin/go "${@:1}"
 fi
     
   
