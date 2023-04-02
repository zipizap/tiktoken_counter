#!/usr/bin/env bash
# Paulo Aleixo Campos
__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
function shw_info { echo -e '\033[1;34m'"$1"'\033[0m'; }
function error { echo "ERROR in ${1}"; exit 99; }
trap 'error $LINENO' ERR
PS4='████████████████████████${BASH_SOURCE}@${FUNCNAME[0]:-}[${LINENO}]>  '
set -o errexit
set -o pipefail
set -o nounset
#set -o xtrace

cd "${__dir}"

export GO111MODULES=on

if ! [ -r go.mod ]
then
  #published in internet would be: go mod init github.com/zipizap/myMod
  MY_UNIQUE_MODULE_NAME=$(basename $PWD)
  go mod init $MY_UNIQUE_MODULE_NAME
fi

# updates dependencies to latest
go get -u ./... 
# dont updates dependencies to latest
#go get /... 

if [[ "${1:-}" == "vendor" ]]; then
  go mod vendor
fi
