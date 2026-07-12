use std::collections::{VecDeque};
use std::collections::BTreeMap;

pub fn find_safe_walk(grid: Vec<Vec<i32>>, health: i32) -> bool {
    // 3286
    let m = grid.len();
    if m == 0 || grid[0].is_empty() { return false; }
    let n = grid[0].len();

    // Track the maximum health remaining when reaching each cell
    let mut max_health = vec![vec![-1; n]; m];
    let mut queue = VecDeque::new();

    // Initialize starting position
    let start_health = health - grid[0][0];
    if start_health <= 0 {
        return false;
    }
    
    max_health[0][0] = start_health;
    queue.push_back((0, 0));

    let directions = [(-1, 0), (1, 0), (0, -1), (0, 1)];

    while let Some((x, y)) = queue.pop_front() {
        let current_h = max_health[x][y];

        // If we reached the bottom-right corner with > 0 health, we succeed
        if x == m - 1 && y == n - 1 && current_h > 0 {
            return true;
        }

        for &(dx, dy) in &directions {
            let nx = x as i32 + dx;
            let ny = y as i32 + dy;

            // Check grid bounds
            if nx >= 0 && nx < m as i32 && ny >= 0 && ny < n as i32 {
                let nx = nx as usize;
                let ny = ny as usize;

                let next_h = current_h - grid[nx][ny];

                // Only proceed if this path yields strictly more health than previously found
                if next_h > max_health[nx][ny] && next_h > 0 {
                    max_health[nx][ny] = next_h;
                    
                    // 0-1 BFS optimization: 
                    // If moving to an empty cell (cost 0), prioritize it by pushing to front
                    if grid[nx][ny] == 0 {
                        queue.push_front((nx, ny));
                    } else {
                        queue.push_back((nx, ny));
                    }
                }
            }
        }
    }

    max_health[m - 1][n - 1] > 0
}

pub fn sum_and_multiply(n: i32) -> i64 {
    // 3754
    let mut x = 0i64;
    let mut sum = 0i64;
    let mut t = n;
    let mut digit_count = 0;
    while t > 0  {
        let d = (t % 10) as i64;
        t = t / 10;
        if d > 0 {
            x = (d * 10_i64.pow(digit_count)) + x;
            digit_count += 1;
            sum += d;
        }
    }
    x * sum
}

pub fn array_rank_transform(arr: Vec<i32>) -> Vec<i32> {
    // 1331
    let mut map: BTreeMap<i32, i32> = arr.iter().map(|&x| (x, 0)).collect();
    for (rank, val) in map.values_mut().enumerate() {
        *val = (rank + 1) as i32;
    }
    arr.into_iter().map(|num| map[&num]).collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_find_safe_walk_1() {
        let grid = vec![
            vec![0, 1, 0, 0, 0],
            vec![0, 1, 0, 1, 0],
            vec![0, 0, 0, 1, 0],
        ];
        assert_eq!(find_safe_walk(grid, 1), true);
    }

    #[test]
    fn test_find_safe_walk_2() {
        let grid = vec![vec![1, 1, 1], vec![1, 0, 1], vec![1, 1, 1]];
        assert_eq!(find_safe_walk(grid, 5), true);
    }

    #[test]
    fn test_find_safe_walk_3() {
        let grid = vec![vec![1,1,1,1]];
        assert_eq!(find_safe_walk(grid, 4), false);
    }
}
