use std::collections::HashSet;

pub fn repeated_n_times(nums: Vec<i32>) -> i32 {
    let mut seen = HashSet::new();
    for num in nums {
        if seen.contains(&num) {
            return num;
        }
        seen.insert(num);
    }
    // According to the problem description, there will always be an N-repeated element,
    // so this line should theoretically not be reached.
    // However, Rust requires a return value for all paths.
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_repeated_n_times() {
        assert_eq!(repeated_n_times(vec![1, 2, 3, 3]), 3);
        assert_eq!(repeated_n_times(vec![2, 1, 2, 5, 3, 2]), 2);
        assert_eq!(repeated_n_times(vec![5, 1, 5, 2, 5, 3, 5, 4]), 5);
    }
}