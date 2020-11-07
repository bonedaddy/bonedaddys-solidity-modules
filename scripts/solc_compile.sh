#! /bin/bash

CONTRACT="$1"
OUTPUT="$2"

solc \
    --bin \
    --abi \
    --optimize \
    --optimize-runs 200 \
    --hashes \
    --devdoc \
    --userdoc \
    --pretty-json \
    --output-dir="$OUTPUT" \
    --overwrite \
    "$CONTRACT"
	# solc --bin --abi --optimize --optimize-runs 200 --hashes --devdoc --userdoc --pretty-json --output-dir="$OUJTPUT" "$CONTRACT"