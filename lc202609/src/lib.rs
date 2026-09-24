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

pub fn uniform_array(nums1: Vec<i32>) -> bool {
    // 3876
    let min_odd = nums1.iter().filter(|&&x| x % 2 != 0).min();
    !nums1.iter().any(|&x| x % 2 == 0 && min_odd.is_some_and(|&m| x < m))
}

pub fn first_stable_index(nums: Vec<i32>, k: i32) -> i32 {
    // 3903, 3904
    // iter is faster (no boundary checking)
    let mut suffix_min = vec![0; nums.len()];
    suffix_min[nums.len() - 1] = nums[nums.len() - 1];
    for (index, &number) in nums.iter().enumerate().rev().skip(1) {
        suffix_min[index] = suffix_min[index + 1].min(number);
    }
    let mut max = i32::MIN;
    for (index, (min, number)) in suffix_min.into_iter().zip(nums).enumerate() {
        max = max.max(number);
        if max - min <= k {
            return index as i32;
        }
    }
    -1
}

pub fn count_commas(n: i32) -> i32 {
    // 3870
    if n < 1000 {
        return 0
    }
    n - 999
}

pub fn count_commas_3871(n: i64) -> i64 {
    // 3871
    let mut total_commas = 0;
    let mut lower_bound = 1000;
    while n >= lower_bound {
        total_commas += n - lower_bound + 1;
        // Move to the next comma tier (multiply by 1000)
        if let Some(next_bound) = lower_bound.checked_mul(1000) {
            lower_bound = next_bound;
        } else {
            break;
        }
    }
    total_commas
}

pub fn total_numbers(digits: Vec<i32>) -> i32 {
    // 3483
    let mut pool_counts = [0; 10];
    for &digit in &digits {
        pool_counts[digit as usize] += 1;
    }
    let valid_three_digit_evens = (100..1000).step_by(2);
    let mut ans = 0;
    for num in valid_three_digit_evens {
        let mut current_counts = [0; 10];
        let d1 = num / 100;          
        let d2 = (num / 10) % 10;    
        let d3 = num % 10;           
        current_counts[d1 as usize] += 1;
        current_counts[d2 as usize] += 1;
        current_counts[d3 as usize] += 1;
        let mut can_form = true;
        for i in 0..10 {
            if pool_counts[i] < current_counts[i] {
                can_form = false;
                break;
            }
        }
        if can_form {
            ans += 1;
        }
    }
    ans
}

pub fn is_rectangle_overlap(rec1: Vec<i32>, rec2: Vec<i32>) -> bool {
    // 836
    if rec1[0] == rec1[2] || rec1[1] == rec1[3] || rec2[0] == rec2[2] || rec2[1] == rec2[3] {
        return false;
    }
    rec1[0] < rec2[2] && rec2[0] < rec1[2] && rec1[1] < rec2[3] && rec2[1] < rec1[3]

}

pub fn check_overlap(radius: i32, x_center: i32, y_center: i32, x1: i32, y1: i32, x2: i32, y2: i32) -> bool {
    // 1401
    let closest_x = x_center.clamp(x1, x2);
    let closest_y = y_center.clamp(y1, y2);
    let distance_x = x_center - closest_x;
    let distance_y = y_center - closest_y;
    let squared_distance = distance_x * distance_x + distance_y * distance_y;
    squared_distance <= radius * radius
}

pub fn reverse_degree(s: String) -> i32 {
    // 3498
    let mut sm = 0i32;
    for (i, ch) in s.chars().enumerate() {
        let reverse_alphabet_pos = 26 - (ch as i32 - 'a' as i32);
        sm += reverse_alphabet_pos * (i as i32 +1)
    }
    sm
}

pub fn smallest_index(nums: Vec<i32>) -> i32 {
    // 3550
    for (i, num) in nums.into_iter().enumerate() {
        let (mut sm, mut tmp) = (0, num);
        while tmp > 0 {
            sm += tmp % 10;
            tmp = tmp / 10;
        }
        if sm == i as i32 {
            return i as i32
        }
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