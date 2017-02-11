# Game of Life


# Plan

  I. Structure:
    1. Get single cell to alternate through life and death

    2. Get a single row working. Cell[0] is alive on first run, Cell[1] alive, Cell[0] dead on second run. Repeat down the row for each run.

    3. Get rows and columns to light up; simple rules.

  II. Rules:
    1. Implement simple rules (May already be completed by I.3.)

    2. Implement complex rules

# Actual Process:

  1. Structure -- Get single cell to alternate between life and death (1 and 0, respectively)
     - This also bootstrapped getting a row to alternate (in unison) between life and death.

  2. Rules -- Get single cell in row to alternate between life and death based on the state of itself and its neighbors.
     - First cell in the row dies if any other cell is alive
     - First cell comes to life if all other cells are dead
     - Second cell (and n+1) comes to life if it has a living neighbor (living neighbor dies)
