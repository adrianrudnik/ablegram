#!/bin/bash -e

for fname in "$@"; do
  if [ ! -f "$fname" ]; then
      echo "File $fname not found."
      exit 1
  fi

  FP=$(realpath "${fname}")

  zcat "${FP}" >"${FP%.*}.xml"

  echo "${FP%.*}.xml"
done
