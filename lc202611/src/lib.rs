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

pub fn find_the_distance_value(arr1: Vec<i32>, mut arr2: Vec<i32>, d: i32) -> i32 {
    // 1385
    arr2.sort_unstable();
    let mut distance_value = 0;
    for num in arr1 {
        let idx = arr2.binary_search(&num).unwrap_or_else(|idx| idx);
        if idx < arr2.len() && (arr2[idx] - num).abs() <= d {
            continue;
        }
        if idx > 0 && (num - arr2[idx - 1]).abs() <= d {
            continue;
        }
        distance_value += 1;
    }
    distance_value
}
