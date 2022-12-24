import numpy as np
import re

# Update this with your input
input_lines = """
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
""".split('\n')[1:-1]

cols = 0
rows = len(input_lines)-2
for i in range(rows):
    line = input_lines[i]
    cols = max(cols, len(line))

# Some of these are not really required
grid = []
row_ranges = []
for i in range(rows):
    line = input_lines[i]
    line += ' ' * (cols - len(line))
    row_start = None
    row_end = None
    for j in range(cols):
        if line[j] != ' ':
            if row_start is None:
                row_start = j
        if line[j] == ' ':
            if row_start is not None:
                row_end = j
                break
    if row_end == None:
        row_end = j + 1
    row_ranges.append((row_start, row_end))
    grid.append(line)

col_ranges = []
for i in range(cols):
    line = ''.join(grid[j][i] for j in range(rows))
    col_start = None
    col_end = None
    for j in range(rows):
        if line[j] != ' ':
            if col_start is None:
                col_start = j
        if line[j] == ' ':
            if col_start is not None:
                col_end = j
                break
    if col_end == None:
        col_end = j + 1
    col_ranges.append((col_start, col_end))

lengths = re.split('[RL]', input_lines[-1])
directions = re.split('[0-9]+', input_lines[-1])
sequence = [(x, y) for x, y in zip(lengths, directions[1:])]

location = [col_ranges[row_ranges[0][0]][0], row_ranges[0][0]]
facing = 0
MAPPING = {0: [0, 1], 1: [1, 0], 2: [0, -1], 3: [-1, 0]}
for amount, direction in sequence:
    # print(amount, facing, location, direction)
    for i in range(int(amount)):
        new_location = [
            location[0] + MAPPING[facing][0],
            location[1] + MAPPING[facing][1]
        ]
        if (facing == 0 or facing == 2) and (new_location[1] == row_ranges[new_location[0]][1]):
            new_location[1] = row_ranges[new_location[0]][0]
        if (facing == 0 or facing == 2) and (new_location[1] == row_ranges[new_location[0]][0] - 1):
            new_location[1] = row_ranges[new_location[0]][1] - 1
        if (facing == 1 or facing == 3) and (new_location[0] == col_ranges[new_location[1]][1]):
            new_location[0] = col_ranges[new_location[1]][0]
        if (facing == 1 or facing == 3) and (new_location[0] == col_ranges[new_location[1]][0] - 1):
            new_location[0] = col_ranges[new_location[1]][1] - 1
        if grid[new_location[0]][new_location[1]] == '#':
            break
        else:
            location = new_location
    if direction != '':
        facing += 1 if direction == 'R' else -1
    facing %= 4

part1 = (location[0]+1)*1000+(location[1]+1)*4+facing
print(part1)

side_length = min([y-x for x, y in row_ranges])
sides = []
for i in range(rows//side_length):
    row = i * side_length
    for j in range((row_ranges[row][1]-row_ranges[row][0])//side_length):
        sides.append((i, row_ranges[row][0] // side_length + j))

def fill(location, visited, sides, cube, shadow_cube, shadow_mapping):
    for m in range(side_length):
        for n in range(side_length):
            cube[0][m+1][n+1] = 2 if grid[location[0]*side_length+m][location[1]*side_length+n] == '#' else 1
    for m in range(side_length):
        for n in range(side_length):
            shadow_cube[0][m+1][n+1] = side_length * side_length * len(visited) + side_length * m + n
            shadow_mapping[side_length * side_length * len(visited) + side_length * m + n] = [location[0]*side_length+m, location[1]*side_length+n]
    visited.add(location)
    for side in sides:
        if side in visited: continue
        flip_axis = None
        if side[0] == location[0]:
            if side[1] == location[1] - 1:
                flip_axis = (2, 0)
            elif side[1] == location[1] + 1:
                flip_axis = (0, 2)
        elif side[1] == location[1]:
            if side[0] == location[0] - 1:
                flip_axis = (1, 0)
            elif side[0] == location[0] + 1:
                flip_axis = (0, 1)
        if flip_axis is None: continue
        cube = np.rot90(cube, k=1, axes=flip_axis)
        shadow_cube = np.rot90(shadow_cube, k=1, axes=flip_axis)
        cube, shadow_cube, shadow_mapping = fill(side, visited, sides, cube, shadow_cube, shadow_mapping)
        cube = np.rot90(cube, k=-1, axes=flip_axis)
        shadow_cube = np.rot90(shadow_cube, k=-1, axes=flip_axis)

    return cube, shadow_cube, shadow_mapping
cube, shadow_cube, shadow_mapping = fill((0, 2), set(), sides, np.zeros([side_length+2, side_length+2, side_length+2]), np.zeros([side_length+2, side_length+2, side_length+2]), dict())

lengths = re.split('[RL]', input_lines[-1])
directions = re.split('[0-9]+', input_lines[-1])
sequence = [(x, y) for x, y in zip(lengths, directions[1:])]

location = [1, 1]
facing = 0
MAPPING = {0: [0, 1], 1: [1, 0], 2: [0, -1], 3: [-1, 0]}

flips = []
for amount, direction in sequence:
    for i in range(int(amount)):
        new_location = [
            location[0] + MAPPING[facing][0],
            location[1] + MAPPING[facing][1]
        ]
        
        edge = False

        if new_location[0] == 0:
            new_location[0] = side_length
            edge = True
        if new_location[0] == side_length + 1:
            new_location[0] = 1
            edge = True
        if new_location[1] == 0:
            new_location[1] = side_length
            edge = True
        if new_location[1] == side_length + 1:
            new_location[1] = 1
            edge = True

        if facing == 0 or facing == 2:
            axes = (0, 2)
            k = 1 if facing == 0 else -1
        else:
            axes = (0, 1)
            k = 1 if facing == 1 else -1

        if edge:
            new_cube = np.rot90(cube, k=k, axes=axes)
            if new_cube[0][new_location[0]][new_location[1]] == 2:
                continue
            cube = new_cube
            shadow_cube = np.rot90(shadow_cube, k=k, axes=axes)
            flips.append((axes, k))
        
        if cube[0][new_location[0]][new_location[1]] == 2:
            continue
        location = new_location

    if direction != '':
        facing += 1 if direction == 'R' else -1
    facing %= 4

smallest = shadow_cube[0][1:-1, 1:-1].min()
shadow_location = shadow_cube[0][location[0]][location[1]]
if shadow_cube[0][1][-2] == smallest: facing += 3
elif shadow_cube[0][-2][-2] == smallest: facing += 2
elif shadow_cube[0][-2][1] == smallest: facing += 1
else: facing += 0
facing %= 4
location = shadow_mapping[int(shadow_location)]
part2 = (location[0]+1)*1000+(location[1]+1)*4+facing

print(part2)
