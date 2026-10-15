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
