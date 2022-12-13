from collections import deque

grid = [list(x) for x in open(0).read().strip().splitlines()]

for r, row in enumerate(grid):
    for c, item in enumerate(row):
        if item == "S":
            grid[r][c] = "a"
        if item == "E":
            endR = r
            endC = c
            grid[r][c] = "z"

queue = deque()
queue.append((0, endR, endC))
visited = {(endR, endC)}

while queue:
    dist, row, col = queue.popleft()
    for newR, newC in [(row + 1, col), (row - 1, col), (row, col + 1), (row, col - 1)]:
        if newR < 0 or newC < 0 or newR >= len(grid) or newC >= len(grid[0]):
            continue
        if (newR, newC) in visited: # If we've already visited the target cell on this path
            continue
        if ord(grid[newR][newC]) - ord(grid[row][col]) < -1: # If the height of the next is more than one lower
            continue
        if grid[newR][newC] == 'a': # Found a valid starting point
            print(dist + 1)
            exit(0)
        visited.add((newR, newC))
        queue.append((dist + 1, newR, newC))
