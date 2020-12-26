import enum
from collections import defaultdict
from functools import reduce
from operator import mul
from typing import List, Tuple, Optional, Dict

import numpy as np

EMPTY = 0
FILLED = 1
TileType = np.array


class MatchType(enum.IntEnum):
    left_right = 0
    right_left = 1
    top_bottom = 2
    bottom_top = 3


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
        np.count_nonzero(tile1[:, -1] - tile2[:, 0]),  # left_right
        np.count_nonzero(tile2[:, -1] - tile1[:, 0]),  # right_left
        np.count_nonzero(tile1[-1, :] - tile2[0, :]),  # top_bottom
        np.count_nonzero(tile2[-1, :] - tile1[0, :]),  # bottom_top
    ]
    for i in range(len(placements)):
        if placements[i] == 0:
            return MatchType(i)
    return None


def rotate_tile(tile: TileType, rotation_count: int) -> TileType:
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


def can_match(tile1: TileType, tile2: TileType) -> Tuple[
    Optional[MatchType], Optional[int], Optional[bool]
]:
    """
    :rtype: Tuple[
        how_matched: Optional[MatchType],
        rotation_count Optional[int],
        flipud: Optional[bool]
    ]
    """
    match, rotation_count = can_match_rotate(tile1, tile2)
    if match is not None:
        return match, rotation_count, False

    tile2_vertical_flip = np.flipud(tile2)
    match, rotation_count = can_match_rotate(tile1, tile2_vertical_flip)
    if match is not None:
        return match, rotation_count, True

    return None, None, None


def apply_transformation(
        tile: TileType,
        rotation_count: int,
        flip: bool
) -> TileType:
    tile = np.rot90(tile, rotation_count)
    tile = np.flipud(tile)
    return tile


def part1() -> Tuple[int, Dict[int, TileType]]:
    matches = defaultdict(dict)
    tile_ids = tiles.keys()
    for t1 in tile_ids:
        for t2 in tile_ids:
            if t1 != t2:
                match, rot_count, flipped = can_match(tiles[t1], tiles[t2])
                if match is not None:
                    matches[t1].update({
                        t2: (
                            match,
                            apply_transformation(
                                tiles[t2],
                                rot_count,
                                flipped
                            )
                        )
                    })

    corners = [
        piece_id
        for piece_id, matched_piece in matches.items()
        if len(matched_piece) == 2
    ]
    assert len(corners) == 4
    return reduce(mul, corners), matches


# def part2(matches: Dict[int, TileType])

if __name__ == '__main__':
    part1_answer, tile_matches = part1()
    print("Part 1:", part1_answer)
