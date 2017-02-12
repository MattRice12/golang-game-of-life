# Game of Life
Conway's Game of Life:

  # General:
  	Every cell interacts with its eight neighbors, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time (generation), the following transitions occur:

  # Rules:
  	1) Any live cell with fewer than two live neighbors die, as if by isolation
  	2) Any live cell with two or three live neighbors lives on to the next generation
  	3) Any live cell with more than three live neighbors dies, as if by overcrowding
  	4) Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction

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

  3. Structure -- Get Rows and Columns together to form a grid.

  4. Rules -- Still basic rules. Same as in (2) above, but now life cascades down the grid (1 row at a time)

  5. Play around with x and y axis, trying to figure out how to implement simple rules based on the status of neighboring cells.

  6. Implement starting positions

  7. Set up a counter for amount of living neighbors

  8. Add game of life rules

  9. Wait to update lifecycle until ALL cells have been accounted for (don't update as you go)

  10. Game works.

  11. Create new colony to spawn every 20 generations

  12. When colony reaches bottom of grid, it should reappear at the top (ditto for sides)

  13. Put `Grid` as a value in the Game struct.
    - This didn't worked as much as I'd hoped because during every generation 2 grids should exist:
      - 1 grid that holds the current generation's info (which I use to apply the rules to in order to determine life/death in the next generation)
      - 1 grid that builds the next generation.
    - So I'd either need 2 Grid values in the Game struct, or 2 functions/methods for each grid (one that references Game and one that outputs a value). Neither option seems efficient, so I'm not including Grid as a value in the struct.
      - The downside is that for most functions I end up calling (game Game, grid [][]int) rather than just (g Game).
