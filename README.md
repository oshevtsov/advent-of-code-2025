# Advent of Code - 2025

My solutions to [Advent of Code](https://adventofcode.com/) 2025. Let's Go! 🎅

## Structure of solutions

Solutions are separated by the programming language used to solve it. The solution for
each day is contained in a folder called, e.g. `day-01`. The folder contains three
subfolders:

- `inputs`: contains all the inputs specified in the day's puzzle
- `part-1`: contains the solution to the first part of the puzzle
- `part-2`: contains the solution to the second part of the puzzle

Inside each day's folder, there will also be a `README.md` file describing the puzzle.

## Running solutions

In the root of the repository you will find a file called `Taskfile.yml`. In order to run
a solution, make sure you have the [Task CLI](https://taskfile.dev/) installed, and run

```sh
task LANG=go DAY=day-01 PART=part-1
```

The parameters above are obviously inputs that choose which day and part should the
solution be run for.
