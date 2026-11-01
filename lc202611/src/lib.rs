pub fn min_moves(nums: Vec<i32>) -> i32 {
    // 453
    let sum: i32 = nums.iter().sum();
    let n = nums.len() as i32;
    sum - n * *nums.iter().min().unwrap_or(&0)
}