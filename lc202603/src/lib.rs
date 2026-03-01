pub fn min_partitions(n: String) -> i32 {
    // 1689
    // The answer is the maximum digit in the string.
    n.chars().map(|c| (c as u8 - b'0') as i32).max().unwrap_or(0)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_min_partitions() {
        assert_eq!(min_partitions("32".to_string()), 3);
        assert_eq!(min_partitions("82734".to_string()), 8);
        assert_eq!(min_partitions("27346209830709182346".to_string()), 9);
    }
}
