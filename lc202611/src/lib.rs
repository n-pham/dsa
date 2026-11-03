pub fn min_moves(nums: Vec<i32>) -> i32 {
    // 453
    let sum: i32 = nums.iter().sum();
    let n = nums.len() as i32;
    sum - n * *nums.iter().min().unwrap_or(&0)
}

pub fn min_moves_3736(nums: Vec<i32>) -> i32 {
    // 3736
    let (sm, mx) = nums.iter().fold((0, 0), |(s, m), &num| (s + num, m.max(num)));
    mx*nums.len() as i32 - sm 
}
