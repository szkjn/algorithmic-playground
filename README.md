# Algorithmic Playground

## Overview

This repository is a collection of algorithmic exercises and code challenges ([rendezvous with cassidoo](https://cassidoo.co/newsletter/)/[HackerRank](https://www.hackerrank.com/profile/junseraphinsuzu1)/[Codewars](https://www.codewars.com/users/szkjn)) to master best practices in **Python**, **Go**, **Typescript** and other side-projects languages. *All exercises problems are written in comments at the beginning of the corresponding script.*

## Repository structure

- Each language has its own directory (e.g. `Go`, `Python`, `TypeScript`).
- Within the `Go` directory, every exercise if contained in its own subdirectory, each with a `main.go` file and a `go.mod` file.
- Other languages are organized similarly, following the best practices of each language.

## Content summary

||<img src='https://upload.wikimedia.org/wikipedia/commons/c/c3/Python-logo-notext.svg' width='24'>|<img src='https://upload.wikimedia.org/wikipedia/commons/6/6a/JavaScript-logo.png' width='22'>|<img src='https://upload.wikimedia.org/wikipedia/commons/4/4c/Typescript_logo_2020.svg' width='24'>|<img src='https://upload.wikimedia.org/wikipedia/commons/9/98/Solidity_logo.svg' width='22'>|<img src='https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Aqua.png' width='36'>|
| --- | --- | --- | --- | --- | --- |
| How many opened doors ? | [python](../main/Python/how_many_opened_doors.py) | | | | |
| Anti-Divisor | [python](../main/Python/anti_divisor.py) | | | | [go](../main/Go/anti_divisor.go)|
| Combine strings | [python](../main/Python/combine_strings.py) | | | | |
| Replace consecutive zeros by its length | [python](../main/Python/replace_consecutive_zeros_by_its_length.py) | | | | |
| Max Subarray | [python](../main/Python/max_subarray.py) | [javascript](../main/JavaScript/maxSubarray.js) | | [solidity](../main/Solidity/maxSubarray.sol)| |
| Swap Pairs | [python](../main/Python/swap_pairs.py) | [javascript](../main/JavaScript/swapPairs.js) | | [solidity](../main/Solidity/swapPairs.sol)| [go](../main/Go/swap_pairs.go)
| Decode Morse Code | [python](../main/Python/decode_morse_code.py) | | | | [go](../main/Go/decode_morse_code.go)|
| Sum every other digit | [python](../main/Python/sum_every_other.py) | [javascript](../main/JavaScript/sumEveryOther.js) | [typescript](../main/TypeScript/sumEveryOther.ts) | |[go](../main/Go/sum_every_other.go)|
| Get A1 Reference number | [python](../main/Python/get_a1_ref_num.py) | | | | [go](../main/Go/get_a1_ref_num.go)|
| Are two strings isomorphic ? | [python](../main/Python/is_isomorphic.py) | | | | [go](../main/Go/is_isomorphic.go)|
| Ceaser Cipher | | | | | [go](../main/Go/ceaser_cipher.go)|
| Mars exploration | | | | | [go](../main/Go/mars_exploration.go)|
| Subsequence string search | |[javascript](../main/JavaScript/subsequenceStringSearch.js)|[typescript](../main/TypeScript/subsequenceStringSearch.ts)| | [go](../main/Go/subsequence_string_search.go)|
| Diagonal difference | | | | | [go](../main/Go/diagonal_difference.go)|
| Min-max sum | | | | | [go](../main/Go/min_max_sum.go)|
| Highest candles |[python](../main/Python/highest_candles.py)| | | | [go](../main/Go/highest_candles.go)|
| Time conversion ||[javascript](../main/JavaScript/timeConversion.js)|[typescript](../main/TypeScript/timeConversion.ts)| | [go](../main/Go/time_conversion.go)|
| Phone Book |||| | [go](../main/Go/phone_book.go)|

## CI/CD Pipeline with GitHub Actions

### Linting and code quality

I've set up a Continuous Integration (CI) pipeline using GitHub Actions to ensure code quality and consistency. Here's how it works:

1. **Linting on Push**: Each time code is pushed to the repository, the CI pipeline automatically lints the code.
2. **Tools Used**: I use `golint` to check the Go code for style mistakes and maintain a consistent code style.
3. **GitHub Actions Workflow**: The linting process is defined in the `.github/workflows/go-lint.yml` file.

