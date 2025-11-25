pub fn minimum_operations(nums: &[i32]) -> i32 {
    // 3190
    nums.iter()
        .filter(|&&x| x % 3 != 0)
        .count() as i32
}

pub fn prefixes_div_by5(nums: &[i32]) -> Vec<bool> {
    // 1018
    let mut divisibles = Vec::with_capacity(nums.len());
    let mut remainder = 0;
	for bit in nums {
		remainder = (remainder * 2 + bit) % 5;
        divisibles.push(remainder == 0);
	}
	divisibles
}

pub fn smallest_repunit_div_by_k(k: i32) -> i32 {
    // 1015
    let mut remainder = 0;
    for length in 1..=k {
        remainder = (remainder*10 + 1) % k;
        if remainder == 0 {
			return length
		}
    }
    -1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_smallest_repunit_div_by_k() {
        assert_eq!(smallest_repunit_div_by_k(1), 1);
        assert_eq!(smallest_repunit_div_by_k(2), -1);
        assert_eq!(smallest_repunit_div_by_k(3), 3);
    }

    #[test]
    fn test_prefixes_div_by5() {
        let nums = vec![0, 1, 1, 1, 1, 1, 0];
        let divisibles = vec![true, false, false, false, true, false, false];
        assert_eq!(prefixes_div_by5(&nums), divisibles);
    }

    #[test]
    fn test_minimum_operations() {
        let nums = vec![1, 3, 4, 6, 7, 9];
        assert_eq!(minimum_operations(&nums), 3); // 1,4,7
    }
}