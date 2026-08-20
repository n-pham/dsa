use std::collections::HashSet;

pub fn predict_the_winner_greedy(nums: Vec<i32>) -> bool {
    let (mut left, mut right) = (0, nums.len()-1);
    let (mut score_1, mut score_2) = (0, 0);
    while left <= right {
        if nums[left] < nums[right] {
            score_1 += nums[right];
            right -= 1;
        } else {
            score_1 += nums[left];
            left += 1;
        }
        if left > right { break; }
        if nums[left] < nums[right] {
            score_2 += nums[right];
            right -= 1;
        } else {
            score_2 += nums[left];
            left += 1;
        }
    }
    println!("1: {score_1} 2: {score_2}");
    score_1 >= score_2
}
pub fn predict_the_winner(nums: Vec<i32>) -> bool {
    // 486
    let n = nums.len();
    // dp[i][j] stores the maximum net score the current player can gain from subarray nums[i..=j]
    let mut dp = vec![vec![0; n]; n];

    // Base case: when there is only one number, the player must take it
    for i in 0..n {
        dp[i][i] = nums[i];
    }

    // Build the DP table for subarrays of length 2 to n
    for len in 2..=n {
        for i in 0..=n - len {
            let j = i + len - 1;
            // The current player chooses either nums[i] or nums[j], 
            // and the opponent will optimally choose from the remaining subarray.
            dp[i][j] = (nums[i] - dp[i + 1][j]).max(nums[j] - dp[i][j - 1]);
        }
    }

    // If Player 1's net score for the entire array is >= 0, they can win
    dp[0][n - 1] >= 0
}

pub fn stone_game(nums: Vec<i32>) -> bool {
    // 877
    let n = nums.len();
    // dp[i][j] stores the maximum net score the current player can gain from subarray nums[i..=j]
    let mut dp = vec![vec![0; n]; n];

    // Base case: when there is only one number, the player must take it
    for i in 0..n {
        dp[i][i] = nums[i];
    }

    // Build the DP table for subarrays of length 2 to n
    for len in 2..=n {
        for i in 0..=n - len {
            let j = i + len - 1;
            // The current player chooses either nums[i] or nums[j], 
            // and the opponent will optimally choose from the remaining subarray.
            dp[i][j] = (nums[i] - dp[i + 1][j]).max(nums[j] - dp[i][j - 1]);
        }
    }

    // If Player 1's net score for the entire array is >= 0, they can win
    dp[0][n - 1] >= 0
}

pub fn find_missing_elements(nums: Vec<i32>) -> Vec<i32> {
    // 3731
    let (mut mn, mut mx) = (i32::MAX, i32::MIN);
    let mut set = HashSet::with_capacity(nums.len());
    for num in nums {
        if num > mx {
            mx = num
        }
        if num < mn {
            mn = num
        }
        set.insert(num);
    }
    let mut res = Vec::new();
    for num in mn..=mx {
        if !set.contains(&num) {
            res.push(num);
        }
    }
    res
}

pub fn smallest_number(n: i32, t: i32) -> i32 {
    // 3345
    for num in n..=(n+9) {
        let mut product = 1;
        let mut tmp = num;
        while tmp > 0 {
            product = product * (tmp%10);
            tmp = tmp / 10;
        }
        if product % t == 0 {
            return num
        }
    }
    0
}

pub fn missing_integer(nums: Vec<i32>) -> i32 {
    // 2996
    let mut sm = nums[0];
    let mut prev = nums[0];
    let mut is_sequential = true;
    let mut num_set = HashSet::with_capacity(nums.len());
    num_set.insert(nums[0]);
    for &num in nums.iter().skip(1) {
        num_set.insert(num);
        if is_sequential {
            if num == prev + 1 {
                sm += num;
                prev = num;
            } else {
                is_sequential = false; // Sequential prefix broke
            }
        }
    }
    while num_set.contains(&sm) {
        sm += 1;
    }
    sm
}

pub fn result_array(nums: Vec<i32>) -> Vec<i32> {
    // 3069
    let mut arr1 = vec![nums[0]];
    let mut arr2 = vec![nums[1]];
    for &num in nums.iter().skip(2) {
        if arr1.last() > arr2.last() {
            arr1.push(num);
        } else {
            arr2.push(num);
        }
    }
    arr1.extend(arr2);
    arr1
}
