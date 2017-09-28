#!/bin/bash

###
#  usage: source $0
#
#         Do not execute this directly ... 'source' the script
#
#         This script will create a function named 'drp' with your
#         Digital Rebar Provision endpoint, username, and password.
#
#         We attempt to prevent the password from leaking to standard
#         shell places (eg. do not use an alias, disable history
#         recording of the actual command, "silent" input to the
#         Password read
#
#         You can set the following variables in your environment
#         and they'll be honored:
#
#            DRP_USER       username of the DRP user - defaults to
#                           empty - equivalent to "rocketskates"
#            DRP_ENDPOINT   DRP endpoint defaults to empty, which is
#                           equivalent to "https://127.0.0.1:8092/"
#
#         OR - modify the below two variables to your preferred
#              default user and endpoint URL
###
#MY_DRP_USER="drp"
#MY_DRP_ENDPOINT="https://drp.example.com:8092"

function set_endpoint_info() {
  export _pass
  export DRP_USER
  export DRP_ENDPOINT

  [[ -n ${MY_DRP_USER} ]]     && DRP_USER=$MY_DRP_USER
  [[ -n ${MY_DRP_ENDPOINT} ]] && DRP_ENDPOINT=$MY_DRP_ENDPOINT
  
  [[ -z "$DRP_USER" || -z "$DRP_ENDPOINT" ]] \
    && printf "\nEnter username/password/endpoint info.  For defaults, just hit <Enter>.\n"

  if [[ -z "$DRP_USER" ]]
  then
    printf "\n%65s" "Enter DRP username (eg 'rocketskates'):  "
    read DRP_USER
  fi

  if [[ -z "$DRP_ENDPOINT" ]]
  then
    printf "\n%65s" "Enter DRP endpoint (eg 'https://127.0.0.1:8092'):  "
    read DRP_ENDPOINT
  fi

  DRP_USER=${DRP_USER:-"rocketskates"}
  DRP_ENDPOINT=${DRP_ENDPOINT:-"https://127.0.0.1:8092"}

  if [[ -z "$_pass" ]]
  then
    printf "\n%65s" "Enter '$DRP_USER' password:  "
    read -s _pass
    echo ""
  fi

  _pass=${_pass:-"r0cketsk8ts"}
}

# make sure we are sourced, not executed directly
if [[ "$0" != "-bash" ]] 
then
  echo ""
  echo "Source this file, do not execute it - example:  source $0"
  echo ""
  echo "exiting ..."
  echo ""
  exit 3
fi

set_endpoint_info

function drp() {
	local _drp_cmd
  set_endpoint_info
  HC=$HISTCOMMAND
  export HISTCOMMAND=ignorespace
  ( which drpcli > /dev/null 2>&1 ) \
    || { echo "ERROR: Unable to find 'drpcli' binary in \$PATH"; return; }
	_drp_cmd=" drpcli -E \"${DRP_ENDPOINT}\" -U \"${DRP_USER}\" -P "
  eval $_drp_cmd $( printenv _pass ) $*
  export HISTCOMMAND=$HC
}

function drp_unset() {
  unset DRP_USER; unset DRP_ENDPOINT; unset _pass
}

echo ""
echo "example usage:     drp isos list"
echo "                   drp_unset       # wipe user/pass/endpoint settings"
echo ""

set_endpoint_info
