pub fn minimum_operations(nums: &[i32]) -> usize {
    // 3190
    nums.iter()
        .filter(|&&x| x % 3 != 0)
        .count() as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_minimum_operations() {
        let nums = vec![1, 3, 4, 6, 7, 9];
        assert_eq!(minimum_operations(&nums), 3); // 1,4,7
    }
}