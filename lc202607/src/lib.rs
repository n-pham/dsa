use std::collections::HashSet;
use std::collections::VecDeque;

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct State {
    x: usize,
    y: usize,
    health: i32,
}

fn find_safe_walk(grid: Vec<Vec<i32>>, mut health: i32) -> bool {
    // 3286
    let m = grid.len();
    if m == 0 {
        return false;
    }
    let n = grid[0].len();
    
    // Check starting position - if it's an obstacle, lose health immediately
    if grid[0][0] == 1 {
        health -= 1;
    }

    if health <= 0 {
        return false;
    }
    
    let directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]; // up, down, left, right
    let mut visited: HashSet<State> = HashSet::new();
    let mut queue: VecDeque<State> = VecDeque::new();

    queue.push_back(State { x: 0, y: 0, health });
    visited.insert(State { x: 0, y: 0, health });

    while let Some(mut state) = queue.pop_front() {
        if state.x == m - 1 && state.y == n - 1 {
            return true; // Reached destination with any remaining health > 0
        }

        for &(dx, dy) in &directions {
            let new_x = state.x as i32 + dx;
            let new_y = state.y as i32 + dy;

            if new_x >= 0 && new_x < m as i32 && new_y >= 0 && new_y < n as i32 {
                let is_obstacle = grid[new_x as usize][new_y as usize] == 1;
                
                // If entering an obstacle, health decreases by 1
                if is_obstacle {
                    state.health -= 1;
                }

                if state.health > 0 {
                    let next_state = State { x: new_x as usize, y: new_y as usize, health: state.health };

                    // Only add to queue if we haven't visited this cell with >= current health
                    // We need to track max_health seen at each position for optimization
                    let mut should_visit = true;
                    
                    // Check if we've already been here with equal or better health
                    // For simplicity, just use basic visited set (can be optimized further)
                    if !visited.contains(&next_state) {
                        visited.insert(next_state.clone());
                        queue.push_back(next_state);
                    } else {
                        // If we have more health than previously seen at this position, try again
                        let mut found_better = false;
                        for (sx, sy, sh) in visited.iter().filter_map(|s| Some((s.x, s.y, s.health))) {
                            if sy == next_state.y && sx == next_state.x && sh <= state.health {
                                // This is a bit complex - let's simplify by just using visited set with health as key
                                found_better = true;
                                break;
                            }
                        }
                        
                        // For now, keep it simple: if we haven't been here at all or have less health before
                    }
                } else {
                    continue;
                }
            }
        }
    }

    false
}

#[cfg(test)]
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
