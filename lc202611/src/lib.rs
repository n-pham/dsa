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

pub fn count_servers(grid: Vec<Vec<i32>>) -> i32 {
    // 1267
    let m = grid.len();
    let n = grid[0].len();
    let mut total_servers = 0;

    let mut row_counts = vec![0; m];
    let mut col_counts = vec![0; n];
    
    for i in 0..m {
        for j in 0..n {
            row_counts[i] += grid[i][j];
            col_counts[j] += grid[i][j];
        }
    }
            
    for i in 0..m {
        for j in 0..n {
            if grid[i][j] == 1 && (row_counts[i] > 1 || col_counts[j] > 1) {
                total_servers += 1;
            }
        }
    }
    
    total_servers
}

pub fn longest_monotonic_subarray(nums: Vec<i32>) -> i32 {
    // 3105
    let n = nums.len();
    let (mut i, mut j) = (0, 1);
    let mut longest = 1i32;
    while j < n {
        if nums[j] > nums[j - 1] {
            if i + 1 < j && nums[j - 1] <= nums[j - 2] {
                i = j - 1;
            }
            longest = longest.max((j - i + 1) as i32);
        } else if nums[j] < nums[j - 1] {
            if i + 1 < j && nums[j - 1] >= nums[j - 2] {
                i = j - 1;
            }
            longest = longest.max((j - i + 1) as i32);
        } else {
            i = j;
        }
        j += 1;
    }

    longest
}
