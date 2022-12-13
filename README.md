# `extract` - JSON Key Extractor

JSON Key Extractor written in [Go](https://go.dev/) 1.19.

Does in fact what is so easy to do in Linux, but not in Windows :/

    jq 'keys' config.json | base64 | rev

## Usage

    Usage: extract [OPTION]... -i file -o file
    -h, --help
    -i, --input file     json input file
    -o, --output file    json output file
    -b, --base64         base64 encoded output string
    -r, --reverse        reverse the output string

## Examples

    extract -i config.json -o config.json
    extract -i config.json -o config.json -b
    extract -i config.json -o config.json -b -r

## Build

    make build

---
Copyright © 2022 Typomedia Foundation. All rights reserved.
