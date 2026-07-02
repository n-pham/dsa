use std::collections::{HashSet, VecDeque};

fn find_safe_walk(grid: Vec<Vec<i32>>, mut health: i32) -> bool {
    // 3286 - Find if we can reach bottom-right from top-left with >0 health
    
    let m = grid.len();
    if m == 0 || grid[0].is_empty() {
        return false;
    }
    
    let n = grid[0].len();
    
    // Check starting position - if it's an obstacle, lose health immediately
    if grid[0][0] == 1 {
        health -= 1;
    }

    if health <= 0 || m == 1 && n == 1 {
        return false;
    }
    
    let directions = [(-1i32, 0), (1i32, 0), (0i32, -1), (0i32, 1)]; // up, down, left, right
    
    // Track max health seen at each position for optimization
    let mut visited: Vec<Vec<i32>> = vec![vec![-1; n]; m];
    
    let mut queue: VecDeque<(usize, usize)> = VecDeque::new();

    if grid[0][0] == 0 {
        // Starting cell is safe, no health loss needed (already handled above)
        visited[0][0] = health;
    } else {
        // Starting cell has obstacle, we already deducted health
        if health > 0 {
            visited[0][0] = health;
        }
    }

    queue.push_back((0, 0));

    while let Some((x, y)) = queue.pop_front() {
        if x == m - 1 && y == n - 1 {
            return true; // Reached destination with any remaining health > 0
        }

        for &(dx, dy) in &directions {
            let new_x = (x as i32 + dx).max(0i32);
            let new_y = (y as i32 + dy).max(0i32);

            if new_x < m as i32 && new_y < n as i32 {
                // Check bounds again after clamping
                if new_x >= 0 && new_x < m as usize && new_y >= 0 && new_y < n as usize {
                    let is_obstacle = grid[new_x][new_y] == 1;
                    
                    // If entering an obstacle, health decreases by 1
                    let mut current_health = if is_obstacle { 
                        visited[x][y].saturating_sub(1) 
                    } else { 
                        visited[x][y] 
                    };

                    if current_health > 0 && new_x == m - 1 && new_y == n - 1 {
                        return true; // Reached destination with health > 0
                    }

                    let next_pos = (new_x as usize, new_y as usize);
                    
                    // Only visit if we have more health than before at this position
                    if current_health > visited[new_x][new_y] {
                        visited[new_x][new_y] = current_health;
                        
                        // Add to queue only for non-destination cells or first time visiting destination
                        let is_destination = new_x == m - 1 && new_y == n - 1;
                        if !is_destination || (x, y) != next_pos {
                            queue.push_back(next_pos);
                        } else {
                            return true; // Reached destination with health > 0
                        }
                    }
                }
            }
        }
    }

    false
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
