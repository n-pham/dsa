pub fn count_odds(low: i32, high: i32) -> i32 {
    // 1523
	let both_count = high - low + 1;
	if both_count%2 == 0 {
		return both_count / 2
	}
    both_count/2 + (low % 2)
	
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_count_odds() {
        assert_eq!(count_odds(3, 7), 3);
        assert_eq!(count_odds(8, 10), 1);
    }
}
