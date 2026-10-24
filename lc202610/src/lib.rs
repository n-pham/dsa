use std::collections::HashSet;

pub fn is_valid(s: String) -> bool {
    // 20
    let mut stack = Vec::new();
    for ch in s.chars() {
        match ch {
            '(' | '[' | '{' => stack.push(ch),
            ')' => if stack.pop() != Some('(') { return false; },
            ']' => if stack.pop() != Some('[') { return false; },
            '}' => if stack.pop() != Some('{') { return false; },
            _ => return false,
        }
    }
    stack.is_empty()
}

pub fn check_valid_string(s: String) -> bool {
    // 678
    let mut min_open = 0i32;
    let mut max_open = 0i32;
    for ch in s.chars() {
        match ch {
            '(' => {
                min_open += 1;
                max_open += 1;
            }
            ')' => {
                min_open = (min_open - 1).max(0);
                max_open -= 1;
            }
            '*' => {
                // '*' can be ')': min_open decreases
                // '*' can be '(': max_open increases
                // '*' can be "": min_open stays the same
                min_open = (min_open - 1).max(0);
                max_open += 1;
            }
            _ => return false,
        }
        if max_open < 0 {
            return false;
        }
    }
    min_open == 0
}

pub fn remove_outer_parentheses(s: String) -> String {
    // 1021
    let mut opened = 0;
    s.chars()
        .filter(|&c| {
            if c == '(' {
                opened += 1;
                opened > 1 // Keep if it's not the outermost opening brace
            } else {
                opened -= 1;
                opened > 0 // Keep if it's not the outermost closing brace
            }
        })
        .collect()
}

pub fn sum_of_squares(nums: Vec<i32>) -> i32 {
    // 2778
    let n = nums.len();
    nums.iter()
        .enumerate()
        .filter(|&(index, _)| n % (index + 1) == 0)
        .map(|(_, &number)| number * number)
        .sum()
}

pub fn sum_counts(nums: Vec<i32>) -> i32 {
    // 2913
    let n = nums.len();
    let mut sm = 0;
    for i in 0..n {
        let mut distinct_set: HashSet<i32> = HashSet::new();
        for j in i..n {
            distinct_set.insert(nums[j]);
            let distinct_count = distinct_set.len() as i32;
            sm += distinct_count * distinct_count;
        }
    }
    sm
}

pub fn target_indices(mut nums: Vec<i32>, target: i32) -> Vec<i32> {
    // 2089
    nums.sort_unstable();
    nums.iter()
        .enumerate()
        .skip_while(|&(_, &num)| num < target)
        .take_while(|&(_, &num)| num == target)
        .map(|(i, _)| i as i32)
        .collect()
}

pub fn count_days_together(arrive_alice: String, leave_alice: String, arrive_bob: String, leave_bob: String) -> i32 {
    // 2409
    let alice_start = date_to_day(&arrive_alice);
    let alice_end = date_to_day(&leave_alice);
    let bob_start = date_to_day(&arrive_bob);
    let bob_end = date_to_day(&leave_bob);
    let start_overlap = alice_start.max(bob_start);
    let end_overlap = alice_end.min(bob_end);
    if start_overlap <= end_overlap {
        end_overlap - start_overlap + 1
    } else {
        0
    }
}
fn date_to_day(date: &str) -> i32 {
    let days_in_months = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
    let month: usize = date[0..2].parse().unwrap();
    let day: i32 = date[3..5].parse().unwrap();
    let mut total_days = day;
    for i in 0..(month - 1) {
        total_days += days_in_months[i];
    }
    total_days
}
