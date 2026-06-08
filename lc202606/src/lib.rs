#![allow(dead_code)]

use std::cmp::{max, min};

pub fn minimum_cost(mut cost: Vec<i32>) -> i32 {
    // 2144
    cost.sort_unstable_by(|a, b| b.cmp(a));
    let mut sum = 0;
    for (index, &num) in cost.iter().enumerate() {
        if index % 3 != 2 {
            sum += num;
        }
    }
    sum
}

pub fn earliest_finish_time(land_start_time: Vec<i32>, land_duration: Vec<i32>, water_start_time: Vec<i32>, water_duration: Vec<i32>) -> i32 {
    // 3633
    let calculate_sequence = |s1_start: &Vec<i32>, s1_dur: &Vec<i32>, s2_start: &Vec<i32>, s2_dur: &Vec<i32>| -> i32 {
        // Step 1: Find the earliest time stage 1 can finish
        let mut min_stage1_end = i32::MAX;
        for i in 0..s1_start.len() {
            min_stage1_end = min(min_stage1_end, s1_start[i] + s1_dur[i]);
        }
        
        // Step 2: Use that end time to find the earliest stage 2 can finish
        let mut min_total_end = i32::MAX;
        for i in 0..s2_start.len() {
            let current_finish = max(min_stage1_end, s2_start[i]) + s2_dur[i];
            min_total_end = min(min_total_end, current_finish);
        }
        
        min_total_end
    };

    // Try Land -> Water
    let land_then_water = calculate_sequence(
        &land_start_time, 
        &land_duration, 
        &water_start_time, 
        &water_duration
    );

    // Try Water -> Land
    let water_then_land = calculate_sequence(
        &water_start_time, 
        &water_duration, 
        &land_start_time, 
        &land_duration
    );

    // Return the best path overall
    min(land_then_water, water_then_land)
}

pub fn earliest_finish_time_3635(land_start_time: Vec<i32>, land_duration: Vec<i32>, water_start_time: Vec<i32>, water_duration: Vec<i32>) -> i32 {
    // 3635
    struct Ride {
        start: i32,
        duration: i32,
    }
    fn solve_sequence(
            s1_start: &[i32],
            s1_dur: &[i32],
            s2_start: &[i32],
            s2_dur: &[i32]
        ) -> i32 {
            let n = s2_start.len();
            let mut rides: Vec<Ride> = (0..n)
                .map(|i| Ride { start: s2_start[i], duration: s2_dur[i] })
                .collect();
            
            // Sort Stage 2 rides by start time to enable binary search
            rides.sort_unstable_by_key(|r| r.start);

            // Precompute prefix minimum of durations for rides starting <= X
            let mut pref_min_dur = vec![i32::MAX; n];
            let mut current_min_dur = i32::MAX;
            for i in 0..n {
                current_min_dur = min(current_min_dur, rides[i].duration);
                pref_min_dur[i] = current_min_dur;
            }

            // Precompute suffix minimum of (start + duration) for rides starting > X
            let mut suff_min_end = vec![i32::MAX; n];
            let mut current_min_end = i32::MAX;
            for i in (0..n).rev() {
                current_min_end = min(current_min_end, rides[i].start + rides[i].duration);
                suff_min_end[i] = current_min_end;
            }

            let mut min_total_finish = i32::MAX;

            // Iterate through all possible first stage choices
            for i in 0..s1_start.len() {
                let x = s1_start[i] + s1_dur[i]; // Finish time of Stage 1 ride

                // Use binary search to find the partition point where ride.start > x
                let idx = rides.partition_point(|r| r.start <= x);

                // Case 1: Choose a Stage 2 ride that opens BEFORE or AT time X
                if idx > 0 {
                    let best_prefix_dur = pref_min_dur[idx - 1];
                    min_total_finish = min(min_total_finish, x + best_prefix_dur);
                }

                // Case 2: Choose a Stage 2 ride that opens AFTER time X
                if idx < n {
                    let best_suffix_end = suff_min_end[idx];
                    min_total_finish = min(min_total_finish, best_suffix_end);
                }
            }

            min_total_finish
        }
    // Evaluate both sequence combinations and take the global minimum
    let land_then_water = solve_sequence(&land_start_time, &land_duration, &water_start_time, &water_duration);
    let water_then_land = solve_sequence(&water_start_time, &water_duration, &land_start_time, &land_duration);
    
    min(land_then_water, water_then_land)
}

pub fn total_waviness(num1: i32, num2: i32) -> i32 {
    // 3751 
    fn calculate_waviness(mut n: i32) -> i32 {
        if n < 100 {
            return 0;
        }
        let mut digits = [0u8; 6];
        let mut len = 0;
        // digits in reverse order, which doesn't alter peak/valley logic
        while n > 0 {
            digits[len] = (n % 10) as u8;
            n /= 10;
            len += 1;
        }
        let mut waviness = 0;
        // Check middle elements (exclude first and last digits)
        for i in 1..(len - 1) {
            let curr = digits[i];
            let prev = digits[i - 1];
            let next = digits[i + 1];
            if (curr > prev && curr > next) || (curr < prev && curr < next) {
                waviness += 1;
            }
        }
        waviness
    }
    
    let mut waviness_sum = 0;
    for num in num1..=num2 {
        waviness_sum += calculate_waviness(num);
    }
    waviness_sum
}

pub fn left_right_difference(nums: Vec<i32>) -> Vec<i32> {
    // 2574
    let mut diff_sums = vec![0; nums.len()];
    let mut left_sum = 0;
    let mut right_sum: i32 = nums.iter().sum(); 
    for i in 0..nums.len() {
        right_sum -= nums[i];
        diff_sums[i] = (left_sum - right_sum).abs();
        left_sum += nums[i];
    }
    diff_sums
}

pub fn pivot_array_(nums: Vec<i32>, pivot: i32) -> Vec<i32> {
    // 2161
    let (mut left_nums, mut mid_nums, mut right_nums) = (Vec::new(), Vec::new(), Vec::new());
    for num in nums {
        if num < pivot {
            left_nums.push(num);
        } else if num > pivot {
            right_nums.push(num);
        } else {
            mid_nums.push(pivot);
        }
    }
    left_nums.append(&mut mid_nums);
    left_nums.append(&mut right_nums);
    left_nums
}

pub fn pivot_array_mem(nums: Vec<i32>, pivot: i32) -> Vec<i32> {
    // 2161
    let n = nums.len();
    let mut ans = vec![pivot; n]; // Pre-fill with pivot
    
    let mut left = 0;
    let mut right = n - 1;

    for i in 0..n {
        // Place elements smaller than pivot from the left
        if nums[i] < pivot {
            ans[left] = nums[i];
            left += 1;
        }
        
        // Place elements larger than pivot from the right
        let j = n - 1 - i;
        if nums[j] > pivot {
            ans[right] = nums[j];
            right -= 1;
        }
    }
    
    ans
}
