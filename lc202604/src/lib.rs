use std::collections::HashMap;

pub fn is_leap_year(year: u64) -> bool {
    year.is_multiple_of(400) || (year.is_multiple_of(4) && !year.is_multiple_of(100))
}

pub fn decode_ciphertext_mem(encoded_text: String, rows: i32) -> String {
    let n = encoded_text.len();
    let rows = rows as usize;
    let cols = n / rows;
    let chars: Vec<char> = encoded_text.chars().collect();
    let mut result = String::new();
    for c in 0..cols {
        let mut row = 0;
        let mut col = c;
        // (0, c) -> (1, c+1) -> (2, c+2) ...
        while row < rows && col < cols {
            let index = row * cols + col;
            result.push(chars[index]);
            row += 1;
            col += 1;
        }
    }
    result.trim_end().to_string()
}

pub fn decode_ciphertext(encoded_text: String, rows: i32) -> String {
    // 2075
    let bytes = encoded_text.as_bytes(); // No collection/allocation!
    let rows = rows as usize;
    let cols = bytes.len() / rows;
    let mut result = Vec::new();
    for c in 0..cols {
        let mut r = 0;
        let mut current_col = c;
        // (0, c) -> (1, c+1) -> (2, c+2) ...
        while r < rows && current_col < cols {
            let index = r * cols + current_col;
            if index < bytes.len() {
                result.push(bytes[index]);
            }
            r += 1;
            current_col += 1;
        }
    }
    String::from_utf8_lossy(&result).trim_end().to_string()
}

pub fn judge_circle(moves: String) -> bool {
    // 657
    let (mut x, mut y) = (0, 0);
    for c in moves.chars() {
        match c {
            'R' => x += 1,
            'L' => x -= 1,
            'U' => y -= 1,
            'D' => y += 1,
            _ => return false,
        }
    }
    x == 0 && y == 0
}

pub fn xor_after_queries(nums: Vec<i32>, queries: Vec<Vec<i32>>) -> i32 {
    let n = nums.len();
    if n == 0 {
        return 0;
    }
    let m = 1_000_000_007i64;
    let b = (n as f64).sqrt() as usize + 1;
    let mut nums_64: Vec<i64> = nums.into_iter().map(|x| x as i64).collect();

    let mut small_k: Vec<Vec<&[i32]>> = vec![vec![]; b];

    for q in &queries {
        let k = q[2] as usize;
        if k >= b {
            let l = q[0] as usize;
            let r = q[1] as usize;
            let v = q[3] as i64;
            for i in (l..=r).step_by(k) {
                nums_64[i] = (nums_64[i] * v) % m;
            }
        } else if k > 0 {
            small_k[k].push(q);
        }
    }

    fn mod_pow(mut base: i64, mut exp: i64, m: i64) -> i64 {
        let mut res = 1;
        base %= m;
        while exp > 0 {
            if exp % 2 == 1 {
                res = (res * base) % m;
            }
            base = (base * base) % m;
            exp /= 2;
        }
        res
    }

    fn mod_inverse(n: i64, m: i64) -> i64 {
        mod_pow(n, m - 2, m)
    }

    let mut final_multipliers = vec![1i64; n];
    let mut diff = vec![1i64; n];
    for k in 1..b {
        if small_k[k].is_empty() {
            continue;
        }

        diff.fill(1);
        for q in &small_k[k] {
            let l = q[0] as usize;
            let r = q[1] as usize;
            let v = q[3] as i64;
            let inv_v = mod_inverse(v, m);

            diff[l] = (diff[l] * v) % m;
            let last = l + ((r - l) / k) * k;
            let next = last + k;
            if next < n {
                diff[next] = (diff[next] * inv_v) % m;
            }
        }

        for i in k..n {
            diff[i] = (diff[i] * diff[i - k]) % m;
        }

        for i in 0..n {
            if diff[i] != 1 {
                final_multipliers[i] = (final_multipliers[i] * diff[i]) % m;
            }
        }
    }

    let mut res = 0;
    for i in 0..n {
        let val = (nums_64[i] * final_multipliers[i]) % m;
        res ^= val as i32;
    }
    res
}

pub fn minimum_distance(nums: Vec<i32>) -> i32 {
    // 3741
    let mut indices_map: HashMap<i32, Vec<usize>> = HashMap::new();
    
    for (index, &value) in nums.iter().enumerate() {
        indices_map.entry(value).or_insert_with(Vec::new).push(index);
    }

    let mut min_dist: Option<i32> = None;

    for indices in indices_map.values() {
        // A good tuple requires at least 3 occurrences.
        if indices.len() < 3 {
            continue;
        }

        for r in 0..=(indices.len() - 3) {
            let i_idx = indices[r];
            let k_idx = indices[r + 2]; // k_idx is the maximum index (i_max)
            
            // The distance D = 2 * (i_max - i_min)
            // Since the indices are usize, we cast the difference to i32.
            // We must ensure that i32 can hold the difference (which it should 
            // given typical LeetCode constraints on array size).
            let span = (k_idx.saturating_sub(i_idx)) as i32;
            let current_dist = 2 * span;

            // Update the minimum distance found so far.
            min_dist = match min_dist {
                Some(min_d) => Some(min_d.min(current_dist)),
                None => Some(current_dist),
            };
        }
    }

    min_dist.unwrap_or(-1)
}

pub fn get_min_distance(nums: Vec<i32>, target: i32, start: i32) -> i32 {
    // 1848
    let mut target_indices: Vec<usize> = Vec::new();
    for (index, &value) in nums.iter().enumerate() {
        if value == target {
            target_indices.push(index);
        }
    }
    let distances: Vec<i32> = target_indices
        .iter()
        .map(|&index| {
            let diff = (index as i32 - start as i32).abs();
            diff
        })
        .collect();
    *distances.iter().min().unwrap()
}

pub fn closest_target(words: Vec<String>, target: String, start_index: i32) -> i32 {
    // 2515
    let len = words.len();
    for dist in 0..=(len / 2) {
        // Look forward (with wrap)
        if words[(start_index as usize+ dist) % len] == target {
            return dist as i32;
        }
        // Look backward (with wrap)
        if words[(start_index as usize + len - dist) % len] == target {
            return dist as i32;
        }
    }
    -1
}

pub fn solve_queries(nums: Vec<i32>, queries: Vec<i32>) -> Vec<i32> {
    // 3488
    let n = nums.len();
    
    let mut indices_map: HashMap<i32, Vec<usize>> = HashMap::new();
    for (i, &val) in nums.iter().enumerate() {
        indices_map.entry(val).or_default().push(i);
    }
    
    let mut min_distances = vec![-1; n];
    
    for indices in indices_map.values() {
        let k = indices.len();
        if k < 2 {
            continue;
        }
        
        for i in 0..k {
            let curr = indices[i];
            let prev = indices[(i + k - 1) % k];
            let next = indices[(i + 1) % k];
            
            let d1 = (curr as i32 - prev as i32).abs();
            let dist1 = d1.min(n as i32 - d1);
            
            let d2 = (curr as i32 - next as i32).abs();
            let dist2 = d2.min(n as i32 - d2);
            
            min_distances[curr] = dist1.min(dist2);
        }
    }
    
    queries.iter().map(|&q| min_distances[q as usize]).collect()
}

pub fn mirror_distance(n: i32) -> i32 {
    // 3783
    let (mut tmp_n, mut rev_n) = (n, 0);
    while tmp_n > 0 {
        let digit = tmp_n % 10;
        rev_n = (rev_n * 10) + digit;
        tmp_n /= 10;
    }
    (n - rev_n).abs()
}

pub fn max_distance(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
    // 1855
    let mut max_d = 0;
    let n2 = nums2.len();
    for (i, value) in nums1.iter().enumerate() {
        if i >= n2 { break; }
        let (mut left, mut right) = (i, n2 - 1);
        let mut last_valid_j = i; // Default to i, as j must be >= i
        let mut found = false;
        while left <= right {
            let mid = left + (right - left) / 2;
            if nums2[mid] >= *value {
                last_valid_j = mid;
                found = true;
                left = mid + 1;
            } else {
                if mid == 0 { break; } // Prevent underflow
                right = mid - 1;
            }
        }
        if found && last_valid_j >= i {
            max_d = max_d.max(last_valid_j - i);
        }
    }
    max_d as i32
}

pub fn max_distance_single_group(colors: Vec<i32>) -> i32 {
    let mut max_d = 0;
    let mut last_color = colors[0];
    let mut first_index = 0;
    for (index, color) in (1..).zip(colors.iter()) {
        if *color != last_color {
            max_d = max_d.max(index - first_index - 1);
            first_index = index;
            last_color = *color;
        }
    }
    max_d
}

pub fn max_distance_2078(colors: Vec<i32>) -> i32 {
    // 2078
    let n = colors.len();
    let mut max_d = 0;
    // Compare everything to the first color (looking from the right)
    for i in (0..n).rev() {
        if colors[i] != colors[0] {
            max_d = max_d.max(i as i32);
            break;
        }
    }
    // Compare everything to the last color (looking from the left)
    for i in 0..n {
        if colors[i] != colors[n - 1] {
            max_d = max_d.max((n - 1 - i) as i32);
            break;
        }
    }
    max_d
}

pub fn furthest_distance_from_origin_long(moves: String) -> i32 {
    // 2833
    let (mut d_right, mut d_left) = (0i32, 0i32);
    for c in moves.chars() {
        match c {
            'L' => {
                d_right -= 1;
                d_left -= 1;
            }
            'R' => {
                d_right += 1;
                d_left += 1;
            }
            '_' => {
                d_right += 1;
                d_left -= 1;
            }
            _ => return 0,
        }
    }
    d_right.abs().max(d_left.abs())
}

pub fn furthest_distance_from_origin(moves: String) -> i32 {
    let (dist, underscore_count): (i32, i32) = moves.chars().fold((0, 0), |(d, u), c| match c {
        'L' => (d - 1, u),
        'R' => (d + 1, u),
        _ => (d, u + 1),
    });

    dist.abs() + underscore_count
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solve_queries() {
        let nums = vec![1, 2, 3, 1];
        let queries = vec![0, 3, 1, 2];
        assert_eq!(solve_queries(nums, queries), vec![1, 1, -1, -1]);

        let nums = vec![1, 1, 1, 1];
        let queries = vec![0, 1, 2, 3];
        assert_eq!(solve_queries(nums, queries), vec![1, 1, 1, 1]);

        let nums = vec![1, 2, 3, 4, 1];
        let queries = vec![0, 4];
        assert_eq!(solve_queries(nums, queries), vec![1, 1]);

        let nums = vec![1, 5, 5, 1];
        let queries = vec![0, 1, 2, 3];
        assert_eq!(solve_queries(nums, queries), vec![1, 1, 1, 1]);
    }
}