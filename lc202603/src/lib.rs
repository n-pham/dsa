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
