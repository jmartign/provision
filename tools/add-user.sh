#!/bin/bash

###### WARNING: this script is very simple - it'll nuke an account first
######          if it exists already .... then (re)create the acount with
######          a user collected password.
######          
######          OR - set DRP_USER and/or DRP_PW at command line, like:
######          
######          DRP_USER=drp DRP_PW=drp ./add-user.sh

function xiterr() { [[ $1 =~ ^[0-9]+$ ]] && { XIT=$1; shift; } || XIT=255; echo "FATAL: $*"; exit $XIT; }

DRP_USER=${DRP_USER:-""}
DRP_PW=${DRP_PW:-""}

set +o history

if [[ -z "${DRP_USER}" ]]
then
  echo -n "Enter username of new user:  "
  read DRP_USER
  echo ""
fi

if [[ -z "${DRP_PW}" ]]
then
  echo -n "Enter password for user '$DRP_USER':  "
  read -s DRP_PW
  echo ""
  echo -n "  Re-enter password for '$DRP_USER':  "
  read -s DRP_PW2
  echo ""

  [[ "${DRP_PW}" != "${DRP_PW2}" ]] && xiterr 2 "passwords do not match"
fi

echo ""

if ( drpcli $ENDPOINT users exists $DRP_USER > /dev/null 2>&1 ) 
then
  drpcli $ENDPOINT users destroy $DRP_USER
fi

echo ""
echo "Creating user '$DRP_USER':"
drpcli $ENDPOINT users create $DRP_USER || xiterr 1 "unable to create user '$DRP_USER'"

# create a token for that account that only can set the password and sends that token to new user.
#TOKEN_JSON=`drpcli users token $DRP_USER scope users action password ttl 3600`
#TOKEN=`echo $TOKEN_JSON | jq -r '.Token' `

echo "Setting user password:"
drpcli $ENDPOINT users password "$DRP_USER" "$DRP_PW" || xiterr 1 "failed setting pasword for '$DRP_USER'"

set -o history

exit 0
