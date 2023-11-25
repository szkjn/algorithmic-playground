# Algorithmic Playground

## Overview

This repository is a collection of algorithmic exercises and code challenges ([Rendezvous with cassidoo](https://cassidoo.co/newsletter/) / [HackerRank](https://www.hackerrank.com/profile/junseraphinsuzu1) / [Codewars](https://www.codewars.com/users/szkjn)) to master best practices in **Python**, **Go**, **Typescript** and other side-projects languages. *All exercises problems are written in comments at the beginning of the corresponding script.*

## Repository structure

- Each language has its own directory (`Python`, `Go`, `TypeScript`, `Solidity`).
- Within the `Go` directory, every exercise is contained in its own subdirectory, each with `go.mod`, `main.go` and `main_test.go` files.
- Other languages are organized similarly, following the best practices of each language.

## Content summary

||<img src='https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg' width='24'>|<img src='https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png' width='36'>|<img src='https://upload.wikimedia.org/wikipedia/commons/4/4c/Typescript_logo_2020.svg' width='24'>|<img src='https://upload.wikimedia.org/wikipedia/commons/9/98/Solidity_logo.svg' width='22'>|
| --- | --- | --- | --- | --- |
| How many opened doors ? |[python](../main/Python/how_many_opened_doors.py)||||
| Anti-Divisor |[python](../main/Python/anti_divisor.py)|[go](../main/Go/anti_divisor/main.go)|||
| Combine strings |[python](../main/Python/combine_strings.py)||||
| Replace consecutive zeros by its length |[python](../main/Python/replace_consecutive_zeros_by_its_length.py)||||
| Max Subarray |[python](../main/Python/max_subarray.py)||[typescript](../main/TypeScript/maxSubarray.ts)|[solidity](../main/Solidity/maxSubarray.sol)|
| Swap Pairs |[python](../main/Python/swap_pairs.py)|[go](../main/Go/swap_pairs/main.go)|[typescript](../main/TypeScript/swapPairs.ts)|[solidity](../main/Solidity/swapPairs.sol)|
| Decode Morse Code |[python](../main/Python/decode_morse_code.py)|[go](../main/Go/decode_morse_code/main.go)|[typescript](../main/TypeScript/decodeMorseCode.ts)||
| Sum every other digit |[python](../main/Python/sum_every_other.py)|[go](../main/Go/sum_every_other/main.go)|[typescript](../main/TypeScript/sumEveryOther.ts)||
| Get A1 Reference number |[python](../main/Python/get_a1_ref_num.py)|[go](../main/Go/get_a1_ref_num/main.go)|||
| Are two strings isomorphic ? |[python](../main/Python/is_isomorphic.py)|[go](../main/Go/is_isomorphic/main.go)|||
| Ceaser Cipher ||[go](../main/Go/ceaser_cipher/main.go)|||
| Mars exploration ||[go](../main/Go/mars_exploration/main.go)|||
| Subsequence string search ||[go](../main/Go/subsequence_string_search/main.go)|[typescript](../main/TypeScript/subsequenceStringSearch.ts)||
| Diagonal difference ||[go](../main/Go/diagonal_difference/main.go)|||
| Min-max sum ||[go](../main/Go/min_max_sum/main.go)|||
| Highest candles |[python](../main/Python/highest_candles.py)|[go](../main/Go/highest_candles/main.go)|||
| Time conversion ||[go](../main/Go/time_conversion/main.go)|[typescript](../main/TypeScript/timeConversion.ts)||
| Phone Book ||[go](../main/Go/phone_book/main.go)|||
| Factorial ||[go](../main/Go/factorial/main.go)|||
| Between two sets ||[go](../main/Go/between_two_sets/main.go)|||
| Binary numbers ||[go](../main/Go/binary_numbers/main.go)|||
| Hourglass in 2-d array ||[go](../main/Go/hourglass_in_2d_arr/main.go)|||
| Breaking records ||[go](../main/Go/breaking_records/main.go)|||
| Subarray division ||[go](../main/Go/subarray_division/main.go)|||

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