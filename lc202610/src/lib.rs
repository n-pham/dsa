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
