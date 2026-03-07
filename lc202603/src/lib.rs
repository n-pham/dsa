pub fn min_partitions(n: String) -> i32 {
    // 1689
    // The answer is the maximum digit in the string.
    n.chars().map(|c| (c as u8 - b'0') as i32).max().unwrap_or(0)
}

pub fn find_kth_bit(n: i32, k: i32) -> char {
    // 1545
    let mut chars: Vec<char> = vec!['0'];
    for _ in 1..n {
        let mut part2: Vec<char> = chars.iter()
            .rev()
            .map(|&c| if c == '0' { '1' } else { '0' })
            .collect();
        chars.push('1');
        chars.append(&mut part2);
    }
    chars[k as usize - 1]
}

pub fn num_special(mat: Vec<Vec<i32>>) -> i32 {
    // 1582
    let m = mat.len();
    let n = mat[0].len();

    // Count ones in each row and column
    let mut row_count = vec![0; m];
    let mut col_count = vec![0; n];

    for i in 0..m {
        for j in 0..n {
            if mat[i][j] == 1 {
                row_count[i] += 1;
                col_count[j] += 1;
            }
        }
    }

    // Check for special positions
    let mut result = 0;
    for i in 0..m {
        for j in 0..n {
            if mat[i][j] == 1 && row_count[i] == 1 && col_count[j] == 1 {
                result += 1;
            }
        }
    }

    result
}

pub fn min_operations(s: String) -> i32 {
    // 1758
    let mut count_if_0_first = 0;
    for (i, c) in s.chars().enumerate() {
        if i % 2 == 0 {
            if c != '0' {
                count_if_0_first += 1;
            }
        } else {
            if c != '1' {
                count_if_0_first += 1;
            }
        }
    }
    let count_if_1_first = s.len() as i32 - count_if_0_first;
    count_if_0_first.min(count_if_1_first)
}

pub fn check_ones_segment_fail(s: String) -> bool {
    // 1784
    let mut prev_is_1 = false;
    for char in s.chars() {
        if char == '1' && prev_is_1 {
            return true
        }
        prev_is_1 = char == '1'
    }
    false
}

pub fn check_ones_segment(s: String) -> bool {
    // 1784
    !s.contains("01")
}

pub fn min_flips(s: String) -> i32 {
    // 1888
    let n = s.len();
    if n == 0 { return 0; }
    let t = s.repeat(2);
    let bytes = t.as_bytes();
    let mut even0 = 0;
    let mut even1 = 0;
    let mut odd0 = 0;
    let mut odd1 = 0;

    // initial window: indices 0..n
    for idx in 0..n {
        let b = bytes[idx]; // FIX: use square brackets
        if idx % 2 == 0 {
            if b == b'0' { even0 += 1; } else { even1 += 1; }
        } else {
            if b == b'0' { odd0 += 1; } else { odd1 += 1; }
        }
    }

    let mut ans = i32::MAX;
    let p0_initial = 0 % 2; // always 0
    let flips0_initial = if p0_initial == 0 { even1 + odd0 } else { odd1 + even0 };
    let flips1_initial = if p0_initial == 0 { even0 + odd1 } else { odd0 + even1 };
    ans = flips0_initial.min(flips1_initial);

    // iterate i from 1 to n-1
    for i in 1..n {
        // remove old index i-1
        let old = bytes[i - 1]; // FIX
        let old_parity = (i - 1) % 2;
        if old == b'0' {
            if old_parity == 0 { even0 -= 1; } else { odd0 -= 1; }
        } else {
            if old_parity == 0 { even1 -= 1; } else { odd1 -= 1; }
        }

        // add new index i+n-1
        let new_idx = i + n - 1;
        let newb = bytes[new_idx]; // FIX
        let new_parity = new_idx % 2;
        if newb == b'0' {
            if new_parity == 0 { even0 += 1; } else { odd0 += 1; }
        } else {
            if new_parity == 0 { even1 += 1; } else { odd1 += 1; }
        }

        let p0 = i % 2;
        let flips0 = if p0 == 0 { even1 + odd0 } else { odd1 + even0 };
        let flips1 = if p0 == 0 { even0 + odd1 } else { odd0 + even1 };
        ans = ans.min(flips0.min(flips1));
    }

    ans
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_min_operations() {
        assert_eq!(min_operations("0100".to_string()), 1);
        assert_eq!(min_operations("10".to_string()), 0);
        assert_eq!(min_operations("1111".to_string()), 2);
    }

    #[test]
    fn test_find_kth_bit() {
        assert_eq!(find_kth_bit(3, 1), '0');
        assert_eq!(find_kth_bit(4, 11), '1');
    }

    #[test]
    fn test_min_partitions() {
        assert_eq!(min_partitions("32".to_string()), 3);
        assert_eq!(min_partitions("82734".to_string()), 8);
        assert_eq!(min_partitions("27346209830709182346".to_string()), 9);
    }
}
