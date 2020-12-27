import enum
from collections import defaultdict
from functools import reduce
from operator import mul
from typing import List, Tuple, Optional, Dict, Set

import numpy as np

EMPTY = 0
FILLED = 1
SINGLE_SHOT_SIDE = 10
VALUABLE_SHOT_SIDE = SINGLE_SHOT_SIDE - 2
TileType = np.array


# np.rot90 - rotates counter-clockwise

class MatchType(enum.IntEnum):
    # Relative position
    right = 0
    left = 1
    bottom = 2
    top = 3


tiles = dict()
with open("in.txt", "r") as f:
    tile_id = None
    tile_array = list()
    for line in f.readlines() + [""]:
        line = line.strip()
        if not tile_id:
            tile_id = int(line[len("Tile "): -1])
        else:
            if line:
                # line not empty
                tile_array.append([
                    FILLED if el == "#" else EMPTY
                    for el in line
                ])
            else:
                tiles[tile_id] = np.array(tile_array)
                tile_id = None
                tile_array = list()


def do_match(tile1: TileType, tile2: TileType) -> Optional[MatchType]:
    placements = [
        np.count_nonzero(tile1[:, -1] - tile2[:, 0]),  # right
        np.count_nonzero(tile2[:, -1] - tile1[:, 0]),  # left
        np.count_nonzero(tile1[-1, :] - tile2[0, :]),  # bottom
        np.count_nonzero(tile2[-1, :] - tile1[0, :]),  # top
    ]
    for i in range(len(placements)):
        if placements[i] == 0:
            return MatchType(i)
    return None


def rotate_tile(tile: TileType, rotation_count: int) -> TileType:
    # Counter-clockwise rotation
    return np.rot90(tile, rotation_count)


def can_match_rotate(tile1: TileType, tile2: TileType) -> Tuple[
    Optional[MatchType], Optional[int]
]:
    """
    :rtype: Tuple[
        how_match: Optional[MatchType]
        rotation_count: Optional[int]
    ]
    """
    for rotation_count in range(4):
        match = do_match(tile1, rotate_tile(tile2, rotation_count))
        if match is not None:
            return match, rotation_count
    return None, None


def can_match(
        tile1: TileType,
        tile2: TileType,
        return_except: Optional[Set[int]] = None
) -> Tuple[
    Optional[MatchType], Optional[int], Optional[bool]
]:
    """
    :rtype: Tuple[
        how_matched: Optional[MatchType],
        rotation_count Optional[int],
        flipud: Optional[bool]
    ]
    """
    if return_except is None:
        return_except = [None]
    match, rotation_count = can_match_rotate(tile1, tile2)
    if match is not None and match not in return_except:
        return match, rotation_count, False

    tile2_vertical_flip = np.flipud(tile2)
    match, rotation_count = can_match_rotate(tile1, tile2_vertical_flip)
    if match is not None and match not in return_except:
        return match, rotation_count, True

    return None, None, None


def apply_transformation(
        tile: TileType,
        rotation_count: int,
        flip: bool
) -> TileType:
    tile = rotate_tile(tile, rotation_count)
    if flip:
        tile = np.flipud(tile)
    return tile


def can_match_to_direction(
        base: TileType,
        other: TileType,
        match_type: MatchType
) -> Optional[TileType]:

    return_except = set(map(int, MatchType)).difference({match_type})
    match, rotation_count, flipped_other = can_match(base, other, return_except)
    if match == match_type:
        return apply_transformation(other, rotation_count, flipped_other)
    else:
        return None


def can_match_to_right(base: TileType, other: TileType) -> Optional[TileType]:
    return can_match_to_direction(base, other, MatchType.right)


def can_match_to_bottom(base: TileType, other: TileType) -> Optional[TileType]:
    return can_match_to_direction(base, other, MatchType.bottom)


def part1() -> Tuple[int, Dict[int, List[int]]]:
    # matches = defaultdict(list)
    matches = defaultdict(dict)
    tile_ids = tiles.keys()
    for t1 in tile_ids:
        for t2 in tile_ids:
            if t1 != t2:
                match, rot_count, flipped = can_match(tiles[t1], tiles[t2])
                if match is not None:
                    matches[t1].update({t2: (match)})

    corners = [
        piece_id
        for piece_id, matched_piece in matches.items()
        if len(matched_piece) == 2
    ]
    assert len(corners) == 4
    return reduce(mul, corners), matches


def paste_to_image(
        image: TileType,
        shot: TileType,
        row: int,
        col: int
) -> TileType:
    image[
    row * VALUABLE_SHOT_SIDE:(row + 1) * VALUABLE_SHOT_SIDE,
    col * VALUABLE_SHOT_SIDE:(col + 1) * VALUABLE_SHOT_SIDE
    ] = np.flipud(shot[1:-1, 1:-1])
    return image


def get_image_part(
        image: TileType,
        row: int,
        col: int
) -> TileType:
    return image[
           row * VALUABLE_SHOT_SIDE:(row + 1) * VALUABLE_SHOT_SIDE,
           col * VALUABLE_SHOT_SIDE:(col + 1) * VALUABLE_SHOT_SIDE
           ]


def draw_image(image: TileType, verbose: bool = True):
    result = ""
    for y in image:
        for val in y:
            if val:
                result += "#"
            else:
                result += "."
        result += "\n"
    if verbose:
        print(result)


def reconstruct_image(
        matches: Dict[int, Dict[int, MatchType]]
) -> TileType:
    grid_side = int(np.sqrt(len(matches)))
    image_side = grid_side * VALUABLE_SHOT_SIDE
    image = np.zeros((image_side, image_side), dtype=np.int)
    grid = np.zeros((grid_side, grid_side), dtype=np.int)
    corners = [
        piece_id
        for piece_id, matched_piece in matches.items()
        if len(matched_piece) == 2
    ]

    # ========================================
    # start = corners[0]
    # adjacent_positions = [
    #     v for _, v in matches[start].items()
    # ]
    # # Rotate till right, bottom is reached
    # if MatchType.right in adjacent_positions and MatchType.bottom:
    #     rotate_count = 0
    # elif MatchType.left in adjacent_positions and MatchType.bottom:
    #     rotate_count = 1
    # elif MatchType.left in adjacent_positions and MatchType.top:
    #     rotate_count = 2
    # else:
    #     rotate_count = 3
    # tiles[start] = rotate_tile(tiles[start], rotate_count)
    # grid[0][0] = start
    # image = paste_to_image(image, tiles[start], 0, 0)
    #
    # row, col = 0, 0
    # used = {start}
    # while len(used) < len(tiles):
    #     assert col >= 0
    #     assert row >= 0
    #     col_delta = 0
    #     row_delta = 0
    #     adjacent_ids = set(matches[grid[row][col]].keys()).difference(used)
    #     adjacent_positions = dict([
    #         (matches[grid[row][col]][adj_id], adj_id)
    #         for adj_id in adjacent_ids
    #     ])
    #     adjacent_position_directions = adjacent_positions.keys()
    #     if MatchType.right in adjacent_position_directions:
    #         col_delta += 1
    #         next = adjacent_positions[MatchType.right]
    #     elif MatchType.left in adjacent_position_directions:
    #         col_delta -= 1
    #         next = adjacent_positions[MatchType.left]
    #     else:
    #         row_delta += 1
    #         next = adjacent_positions[MatchType.bottom]
    #
    #     grid[row + row_delta][col + col_delta] = next
    #     _, rotation_count, flip = can_match(tiles[grid[row][col]], tiles[next])
    #     row += row_delta
    #     col += col_delta
    #     image = paste_to_image(
    #         image,
    #         apply_transformation(
    #             tiles[next],
    #             rotate_count,
    #             flip
    #         ),
    #         row,
    #         col
    #     )
    #     used.add(next)
    # a = 1

    # =======================================================================

    leftmost_tile_id = corners[0]  # just find top left by luck
    leftmost_tile_id = 1171
    used = {leftmost_tile_id}
    need_to_flip_first = False
    for row in range(grid_side):
        current_tile_id = leftmost_tile_id
        image = paste_to_image(image, tiles[current_tile_id], row, 0)
        grid[row, 0] = current_tile_id

        # Fill current row
        for col in range(1, grid_side):
            for neighbour_tile_id in matches[current_tile_id]:
                next_tile = can_match_to_right(
                    tiles[current_tile_id],
                    tiles[neighbour_tile_id]
                )
                if next_tile is not None:
                    tiles[neighbour_tile_id] = next_tile
                    current_tile_id = neighbour_tile_id
                    image = paste_to_image(image, next_tile, row, col)
                    grid[row, col] = neighbour_tile_id
                    break
            else:
                raise Exception("Didn't find proper neighbour")

        # Find start for next row
        for neighbour_tile_id in matches[leftmost_tile_id]:
            next_tile = can_match_to_bottom(
                tiles[leftmost_tile_id],
                tiles[neighbour_tile_id]
            )
            if next_tile is not None and neighbour_tile_id not in used:
                tiles[neighbour_tile_id] = next_tile
                leftmost_tile_id = neighbour_tile_id
                break
        else:
            if row == 0:
                # If at the first row we didn't find next row start
                need_to_flip_first = True
            tiles[leftmost_tile_id] = np.flipud(tiles[leftmost_tile_id])
            for neighbour_tile_id in matches[leftmost_tile_id]:
                next_tile = can_match_to_bottom(
                    tiles[leftmost_tile_id],
                    tiles[neighbour_tile_id]
                )
                if next_tile is not None and neighbour_tile_id not in used:
                    tiles[neighbour_tile_id] = next_tile
                    leftmost_tile_id = neighbour_tile_id
                    break
        used.add(leftmost_tile_id)

    print(grid)
    draw_image(image)
    print(need_to_flip_first)
    return image


def part2(matches: Dict[int, List[int]]):
    """
    :param matches: Dict[tile_id: {
        matched_id: (match_type, transformed_tile)
    }]
    :return:
    """
    image = reconstruct_image(matches)
    a = 1


if __name__ == '__main__':
    part1_answer, tile_matches = part1()
    print("Part 1:", part1_answer)
    part2(tile_matches)
