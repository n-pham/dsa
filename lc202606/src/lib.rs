#![allow(dead_code)]

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
