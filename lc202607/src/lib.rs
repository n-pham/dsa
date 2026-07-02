use std::collections::HashSet;
use std::collections::VecDeque;

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct State {
    x: usize,
    y: usize,
    health: i32,
}

fn find_safe_walk(grid: Vec<Vec<i32>>, health: i32) -> bool {
    // 3286
    let m = grid.len();
    if m == 0 {
        return false;
    }
    let n = grid[0].len();
    let directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]; // up, down, left, right
    let mut visited: HashSet<State> = HashSet::new();
    let mut queue: VecDeque<State> = VecDeque::new();

    queue.push_back(State { x: 0, y: 0, health });
    visited.insert(State { x: 0, y: 0, health });

    while let Some(state) = queue.pop_front() {
        if state.x == m - 1 && state.y == n - 1 {
            return state.health > 0;
        }

        for &(dx, dy) in &directions {
            let new_x = state.x as i32 + dx;
            let new_y = state.y as i32 + dy;

            if new_x >= 0 && new_x < m as i32 && new_y >= 0 && new_y < n as i32 {
                let new_health = if grid[new_x as usize][new_y as usize] == 1 {
                    state.health - 1
                } else {
                    state.health
                };

                if new_health > 0 {
                    let next_state = State { x: new_x as usize, y: new_y as usize, health: new_health };

                    if !visited.contains(&next_state) {
                        visited.insert(next_state.clone());
                        queue.push_back(next_state);
                    }
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
