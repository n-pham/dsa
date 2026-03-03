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

#[cfg(test)]
mod tests {
    use super::*;

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
