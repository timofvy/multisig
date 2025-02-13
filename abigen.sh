#!/bin/bash

ABI_DIRS=(
abi
abi
)
PACKAGES=(
safe_proxy_factory_abi
safe_abi
)
ABI_FILES=(
safe_proxy_factory.abi
safe.abi
)

for ((i=0; i<${#ABI_DIRS[@]}; i++)); do
    ABI_FILE="${ABI_DIRS[$i]}/${ABI_FILES[$i]}"
    PACKAGES_FILE="${ABI_DIRS[$i]}/${PACKAGES[$i]}.go"

    if [ ! -f "$PACKAGES_FILE" ] || [ "$(md5sum $ABI_FILE | awk '{print $1}')" != "$(md5sum $PACKAGES_FILE | awk '{print $1}')" ]; then
        abigen --abi=$ABI_FILE --pkg=${PACKAGES[$i]} --out=$PACKAGES_FILE --alias=_totalSupply=totalSupplyPrivate
    fi
done