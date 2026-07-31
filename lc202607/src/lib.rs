use std::collections::{VecDeque};
use std::collections::BTreeMap;
use common::gcd;

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

pub fn gcd_of_odd_even_sums(n: i32) -> i32 {
    // 3658
    let (sum_odd, sum_even) = (n*n, n*(n-1));
    gcd(sum_odd, sum_even)
}

pub fn gcd_sum(nums: Vec<i32>) -> i64 {
    // 3867
    let mut prefix_gcd = Vec::new();
    let mut mx = nums[0];
    for num in nums {
        mx = mx.max(num);
        prefix_gcd.push(gcd(num, mx))
    }
    prefix_gcd.sort_unstable();
    let mut sum: i64 = 0;
    let mut left = 0;
    let mut right = prefix_gcd.len() - 1;
    while left < right {
        let pair_gcd = gcd(prefix_gcd[left], prefix_gcd[right]);
        sum += pair_gcd as i64;
        left += 1;
        right -= 1;
    }
    sum
}

pub fn find_gcd(nums: Vec<i32>) -> i32 {
    // 1979
    let (mut mn, mut mx) = (i32::MAX, i32::MIN);
    for num in nums {
        mn = mn.min(num);
        mx = mx.max(num);
    }
    gcd(mn, mx)
}

pub fn max_product(n: i32) -> i32 {
    // 3536
    let mut t = n;
    let (mut m1, mut m2) = (0, 0); // m1 <= m2
    while t > 0 {
        let d = t % 10;
        t /= 10;
        if d > m2 {
            (m1, m2) = (m2, d)
        } else if d > m1 {
            m1 = d;
        }
    }
    m1 * m2
}

pub fn max_product_2(nums: Vec<i32>) -> i32 {
    // 1464
    let (mut mx1, mut mx2) = (0, 0); // mx1 <= mx2
    for num in nums {
        if num > mx2 {
            (mx1, mx2) = (mx2, num);
        } else if num > mx1 {
            mx1 = num;
        }
    }
    (mx1-1) * (mx2-1)
}

pub fn minimum_pushes(word: String) -> i32 {
    // 3014
    let mut n = word.len();
    let mut cnt = 0;
    while n > 0 {
        cnt += n;
        n = n.saturating_sub(8);
    }
    cnt as i32
}

pub fn minimum_pushes_3016(word: String) -> i32 {
    // 3016
    let mut counts = [0; 26];
    for byte in word.bytes() {
        counts[(byte - b'a') as usize] += 1;
    }
    
    // Sort frequencies in descending order to prioritize frequent letters
    counts.sort_unstable_by(|a, b| b.cmp(a));
    
    let mut pushes = 0;
    for (i, &count) in counts.iter().enumerate() {
        if count == 0 { break; }
        // First 8 keys take 1 push, next 8 take 2 pushes, etc.
        pushes += count * ((i / 8) as i32 + 1);
    }
    
    pushes
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
