#![allow(dead_code)]

pub fn max_rotate_function(nums: Vec<i32>) -> i32 {
    // 396
    let n = nums.len() as i64;
    let mut sum = 0i64;
    let mut f = 0i64;
    // F(0)
    for (i, &num) in nums.iter().enumerate() {
        sum += num as i64;
        f += (i as i64) * (num as i64);
    }
    let mut max_val = f;
    // Derive each subsequent F(k) in O(1) time
    // We iterate backwards to easily access the element that wraps around
    for i in (1..nums.len()).rev() {
        f = f + sum - n * (nums[i] as i64);
        max_val = max_val.max(f);
    }
    max_val as i32
}

pub fn rotated_digits(n: i32) -> i32 {
    // 788
    fn is_rotated_good(x: &i32) -> bool {
        let mut tmp = *x;
        let mut has_diff = false;
        while tmp > 0 {
            match tmp % 10 {
                3 | 4 | 7 => return false,
                2 | 5 | 6 | 9 => has_diff = true,
                _ => {}
            }
            tmp /= 10
        }
        has_diff
    }
    (1..=n).filter(is_rotated_good).count() as i32
}

pub fn rotate_string(s: String, goal: String) -> bool {
    // 796
    let ss = format!("{}{}", s, s);
    ss.contains(&goal) && s.len() == goal.len()
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

pub fn rotate_right(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
    // 61
    if head.is_none() || k == 0 {
        return head;
    }

    let mut head = head;
    let mut len = 0;
    let mut curr = &head;
    while let Some(node) = curr {
        curr = &node.next;
        len += 1;
    }
    let k = k % len;
    if k == 0 {
        return head;
    }
    let mut curr = &mut head;
    for _ in 0..(len - k - 1) {
        curr = &mut curr.as_mut().unwrap().next;
    }
    // Detach the new head and the new tail
    // Use std::mem::take to move the value out of the Option
    let mut new_head = std::mem::take(&mut curr.as_mut().unwrap().next);
    // Connect the old tail to the old head
    let mut tail = &mut new_head;
    while tail.as_mut().unwrap().next.is_some() {
        tail = &mut tail.as_mut().unwrap().next;
    }
    tail.as_mut().unwrap().next = head;
    new_head
}