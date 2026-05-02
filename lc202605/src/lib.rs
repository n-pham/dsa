#![allow(dead_code)]

pub fn max_rotate_function(nums: Vec<i32>) -> i32 {
    // 396
    let n = nums.len() as i64;
    let mut sum = 0i64;
    let mut f = 0i64;
    // F(0)
    for (i, &num) in nums.iter().enumerate() {
        sum += num as i64;
        f += (i as i64) * (num as i64);
    }
    let mut max_val = f;
    // Derive each subsequent F(k) in O(1) time
    // We iterate backwards to easily access the element that wraps around
    for i in (1..nums.len()).rev() {
        f = f + sum - n * (nums[i] as i64);
        max_val = max_val.max(f);
    }
    max_val as i32
}

pub fn rotated_digits(n: i32) -> i32 {
    // 788
    fn is_rotated_good(x: &i32) -> bool {
        let mut tmp = *x;
        let mut has_diff = false;
        while tmp > 0 {
            match tmp % 10 {
                3 | 4 | 7 => return false,
                2 | 5 | 6 | 9 => has_diff = true,
                _ => {}
            }
            tmp /= 10
        }
        has_diff
    }
    (1..=n).filter(is_rotated_good).count() as i32
}