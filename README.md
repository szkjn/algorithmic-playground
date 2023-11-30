# Algorithmic Playground

![example workflow](https://github.com/szkjn/algorithmic-playground/actions/workflows/python-ci.yml/badge.svg)
![example workflow](https://github.com/szkjn/algorithmic-playground/actions/workflows/go-ci.yml/badge.svg)
![example workflow](https://github.com/szkjn/algorithmic-playground/actions/workflows/typescript-ci.yml/badge.svg)

![Python (tests/scripts)](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/szkjn/d63aa1a5139a8720d9142cb6c1c734c2/raw/playground-python-badge.json)
![Go (tests/scripts)](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/szkjn/d63aa1a5139a8720d9142cb6c1c734c2/raw/playground-go-badge.json)
![TypeScript (tests/scripts)](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/szkjn/d63aa1a5139a8720d9142cb6c1c734c2/raw/playground-ts-badge.json)

## Overview

This repository is a collection of algorithmic exercises and code challenges ([Rendezvous with cassidoo](https://cassidoo.co/newsletter/) / [HackerRank](https://www.hackerrank.com/profile/junseraphinsuzu1) / [Codewars](https://www.codewars.com/users/szkjn)) to master best practices in **Python**, **Go**, **Typescript** and other side-projects languages. *All exercises problems are written in comments at the beginning of the corresponding script.*

## Repository structure

- Each language has its own directory (`Python`, `Go`, `TypeScript`, `Solidity`).
- Within the `Go` directory, every exercise is contained in its own subdirectory, each with `go.mod`, `main.go` and `main_test.go` files.
- Other languages are organized similarly, following the best practices of each language.

## Content summary

||<img src='https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg' width='24'>|<img src='https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png' width='36'>|<img src='https://upload.wikimedia.org/wikipedia/commons/4/4c/Typescript_logo_2020.svg' width='24'>|<img src='https://upload.wikimedia.org/wikipedia/commons/9/98/Solidity_logo.svg' width='22'>|
| --- | --- | --- | --- | --- |
| How many opened doors ? |[python](../main/Python/scripts/how_many_opened_doors.py)||||
| Anti-Divisor |[python](../main/Python/scripts/anti_divisor.py)|[go](../main/Go/anti_divisor/main.go)|||
| Combine strings |[python](../main/Python/scripts/combine_strings.py)||||
| Zero run replacer |[python](../main/Python/scripts/zero_run_replacer.py)||||
| Max Subarray |[python](../main/Python/scripts/max_subarray.py)||[typescript](../main/TypeScript/maxSubarray.ts)|[solidity](../main/Solidity/maxSubarray.sol)|
| Swap Pairs |[python](../main/Python/scripts/swap_pairs.py)|[go](../main/Go/swap_pairs/main.go)|[typescript](../main/TypeScript/swapPairs.ts)|[solidity](../main/Solidity/swapPairs.sol)|
| Decode Morse Code |[python](../main/Python/scripts/decode_morse_code.py)|[go](../main/Go/decode_morse_code/main.go)|[typescript](../main/TypeScript/decodeMorseCode.ts)||
| Sum every other digit |[python](../main/Python/scripts/sum_every_other.py)|[go](../main/Go/sum_every_other/main.go)|[typescript](../main/TypeScript/sumEveryOther.ts)||
| Get A1 Reference number |[python](../main/Python/scripts/get_a1_ref_num.py)|[go](../main/Go/get_a1_ref_num/main.go)|||
| Are two strings isomorphic ? |[python](../main/Python/scripts/is_isomorphic.py)|[go](../main/Go/is_isomorphic/main.go)|||
| Ceaser Cipher ||[go](../main/Go/ceaser_cipher/main.go)|||
| Mars exploration ||[go](../main/Go/mars_exploration/main.go)|||
| Subsequence string search ||[go](../main/Go/subsequence_string_search/main.go)|[typescript](../main/TypeScript/subsequenceStringSearch.ts)||
| Diagonal difference ||[go](../main/Go/diagonal_difference/main.go)|||
| Min-max sum ||[go](../main/Go/min_max_sum/main.go)|||
| Highest candles |[python](../main/Python/scripts/highest_candles.py)|[go](../main/Go/highest_candles/main.go)|||
| Time conversion ||[go](../main/Go/time_conversion/main.go)|[typescript](../main/TypeScript/timeConversion.ts)||
| Phone Book ||[go](../main/Go/phone_book/main.go)|||
| Factorial ||[go](../main/Go/factorial/main.go)|||
| Between two sets ||[go](../main/Go/between_two_sets/main.go)|||
| Binary numbers ||[go](../main/Go/binary_numbers/main.go)|||
| Hourglass in 2-d array ||[go](../main/Go/hourglass_in_2d_arr/main.go)|||
| Breaking records ||[go](../main/Go/breaking_records/main.go)|||
| Subarray division ||[go](../main/Go/subarray_division/main.go)|||
| Max difference solver |[python](../main/Python/scripts/max_difference_solver.py)||||
| Linked list construction |[python](../main/Python/scripts/linked_list_construction.py)||||

## Run, lint and test

### Python

- To run scripts :

        cd algorithmic-playground/Python/scripts
        python <name_of_script>.py

- To run linter :

        cd algorithmic-playground/Python/
        black .

- To run tests :

        cd algorithmic-playground/Python/
        pytest

### Go

- To run scripts :

        cd algorithmic-playground/Go/<name_of_folder>
        go run main.go

- To run tests :

        cd algorithmic-playground/Go/<name_of_folder>
        go test

### TypeScript

- To run scripts :

        cd algorithmic-playground/TypeScript/src
        ts-node <name_of_script>

- To run linter :

        cd algorithmic-playground/TypeScript/
        ./node_modules/.bin/eslint .

- To run tests :

        cd algorithmic-playground/TypeScript/tests
        npx jest

## CI/CD Pipeline with GitHub Actions

A Continuous Integration (CI) pipeline has been setup using GitHub Actions to ensure code quality, consistency, and successful builds for both our Go and TypeScript code.

### Go linting and testing

1. **Linting on Push and Pull Requests**: Whenever code is pushed or a pull request is made, the CI pipeline automatically lints the Go code to ensure coding standards are met.
2. **Automated Testing**: Along with linting, the pipeline runs tests for the Go code to validate functionality.
2. **Tools Used**: I use `golint` for linting and Go's built-in testing framework for running tests.
3. **GitHub Actions Workflow**: The linting process is defined in the `.github/workflows/go-lint-and-test.yml` file, utilizing GitHub Actions to automatically execute linting and testing on the latest version of Go.

### TypeScript Build

1. **Build on Push and Pull Requests**: The CI pipeline is configured to automatically build TypeScript files whenever changes are pushed or a pull request is made to the TypeScript directory.
2. **Node.js Setup**: The workflow sets up a Node.js environment to manage dependencies and run the build scripts.
3. **Build Script**: TypeScript files are compiled using the TypeScript compiler configured in the package.json file.
4. **Workflow Details**: The TypeScript build process is defined in the `.github/workflows/typescript-build.yml` file. It ensures that TypeScript code is compiled successfully, preventing runtime errors due to compilation issues.
