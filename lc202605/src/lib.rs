#![allow(dead_code)]

pub fn max_rotate_function(nums: Vec<i32>) -> i32 {
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