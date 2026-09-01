pub fn min_moves(classroom: Vec<String>, energy: i32) -> i32 {
    // 3568
    let m = classroom.len();
    let n = classroom[0].len();
    let grid: Vec<Vec<char>> = classroom.iter().map(|s| s.chars().collect()).collect();

    let mut litter_id = vec![vec![0; n]; m];
    let mut start_x = 0;
    let mut start_y = 0;
    let mut litter_count = 0;

    for i in 0..m {
        for j in 0..n {
            match grid[i][j] {
                'S' => {
                    start_x = i;
                    start_y = j;
                }
                'L' => {
                    litter_id[i][j] = litter_count;
                    litter_count += 1;
                }
                _ => {}
            }
        }
    }

    if litter_count == 0 {
        return 0;
    }

    let max_mask = 1 << litter_count;
    let max_energy = energy as usize;

    // visited[x][y][energy][mask]
    let mut visited = vec![vec![vec![vec![false; max_mask]; max_energy + 1]; n]; m];

    use std::collections::VecDeque;
    let mut queue = VecDeque::new();
    let initial_mask = (1 << litter_count) - 1;
    queue.push_back((start_x, start_y, energy, initial_mask));
    visited[start_x][start_y][energy as usize][initial_mask] = true;

    let dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)];
    let mut moves = 0;

    while !queue.is_empty() {
        let level_size = queue.len();
        for _ in 0..level_size {
            let (x, y, cur_energy, mask) = queue.pop_front().unwrap();

            if mask == 0 {
                return moves;
            }

            if cur_energy <= 0 {
                continue;
            }

            for (dx, dy) in dirs {
                let nx = x as i32 + dx;
                let ny = y as i32 + dy;

                if nx < 0 || nx >= m as i32 || ny < 0 || ny >= n as i32 {
                    continue;
                }

                let nx = nx as usize;
                let ny = ny as usize;

                if grid[nx][ny] == 'X' {
                    continue;
                }

                let nxt_energy = if grid[nx][ny] == 'R' {
                    energy
                } else {
                    cur_energy - 1
                };

                let mut nxt_mask = mask;
                if grid[nx][ny] == 'L' {
                    nxt_mask &= !(1 << litter_id[nx][ny]);
                }

                if !visited[nx][ny][nxt_energy as usize][nxt_mask] {
                    visited[nx][ny][nxt_energy as usize][nxt_mask] = true;
                    queue.push_back((nx, ny, nxt_energy, nxt_mask));
                }
            }
        }
        moves += 1;
    }

    -1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_min_moves_example1() {
        let classroom = vec!["S.".to_string(), "XL".to_string()];
        let energy = 2;
        assert_eq!(min_moves(classroom, energy), 2);
    }

    #[test]
    fn test_min_moves_example2() {
        let classroom = vec!["LS".to_string(), "RL".to_string()];
        let energy = 4;
        assert_eq!(min_moves(classroom, energy), 3);
    }

    #[test]
    fn test_min_moves_example3() {
        let classroom = vec!["L.S".to_string(), "RXL".to_string()];
        let energy = 3;
        assert_eq!(min_moves(classroom, energy), -1);
    }

    #[test]
    fn test_min_moves_no_litter() {
        let classroom = vec!["S.".to_string(), ".R".to_string()];
        let energy = 2;
        assert_eq!(min_moves(classroom, energy), 0);
    }
}