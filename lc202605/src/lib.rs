#![allow(dead_code)]
use std::collections::HashSet;

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

pub fn separate_digits(nums: Vec<i32>) -> Vec<i32> {
    // 2553
    let mut result_digits = Vec::new();
    for mut num in nums {
        let mut digits = Vec::new();
        while num > 0 {
            digits.push((num % 10) as i32);
            num /= 10;
        }
        result_digits.extend(digits.into_iter().rev());
    }
    result_digits
}

pub fn is_good(nums: Vec<i32>) -> bool {
    // 2784
    let mut max_num = nums[0];
    let mut already_has_dup = false;
    let mut current_sum: i64 = nums[0] as i64;
    for &num in &nums[1..] {
        current_sum += num as i64;
        if num == max_num {
                if already_has_dup {
                return false;
            }
            already_has_dup = true;
        } else if num > max_num {
            if already_has_dup {
                return false; 
            }
            max_num = num;
        }
    }
    if !already_has_dup {
        return false;
    }
    // The expected sum of [1, 2, ..., n, n] is (n * (n + 1) / 2) + n
    let n = max_num as i64;
    let expected_sum = (n * (n + 1)) / 2 + n;
    // The length must also match n + 1
    current_sum == expected_sum && nums.len() == (max_num as usize + 1)
}

pub fn get_common(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
    // 2540
    let (len1, len2) = (nums1.len(), nums2.len());
    let (mut i1, mut i2) = (0, 0);
    while i1 < len1 && i2 < len2 {
        if nums1[i1] == nums2[i2] {
            return nums1[i1];
        } else if nums1[i1] < nums2[i2] {
            i1 += 1;
        } else {
            i2 += 1;
        }
    }
    -1
}

pub fn find_the_prefix_common_array(a: Vec<i32>, b: Vec<i32>) -> Vec<i32> {
    // 2657
    let n = a.len();
    let mut seen = vec![false; n + 1];
    let mut rs = vec![0; n];
    let mut common_count = 0;
    for i in 0..n {
        let val_a = a[i] as usize;
        if seen[val_a] {
            common_count += 1;
        } else {
            seen[val_a] = true;
        }

        let val_b = b[i] as usize;
        if seen[val_b] {
            common_count += 1;
        } else {
            seen[val_b] = true;
        }

        rs[i] = common_count;
    }

    rs
}

pub fn number_of_special_chars(word: String) -> i32 {
    // 3120
    let (mut upper_set, mut lower_set) = (HashSet::new(), HashSet::new());
    for ch in word.chars() {
        if ch.is_ascii_uppercase() {
            upper_set.insert(ch);
        } else if ch.is_ascii_lowercase() {
            lower_set.insert(ch.to_ascii_uppercase());
        }
    }
    upper_set.intersection(&lower_set).count() as i32
}

pub fn number_of_special_chars_3121(word: String) -> i32 {
    // 3121
    let mut last_lower = [-1; 26];
    let mut first_upper = [-1; 26];
    
    for (i, ch) in word.chars().enumerate() {
        let idx = i as i32;
        if ch.is_ascii_lowercase() {
            let letter_idx = (ch as u8 - b'a') as usize;
            last_lower[letter_idx] = idx; // Always update to find the LAST occurrence
        } else if ch.is_ascii_uppercase() {
            let letter_idx = (ch as u8 - b'A') as usize;
            if first_upper[letter_idx] == -1 {
                first_upper[letter_idx] = idx; // Only set once to find the FIRST occurrence
            }
        }
    }
    
    let mut special_count = 0;
    for i in 0..26 {
        // Condition 1: Must see both versions
        // Condition 2: Last lowercase index must be strictly smaller than the first uppercase index
        if last_lower[i] != -1 && first_upper[i] != -1 && last_lower[i] < first_upper[i] {
            special_count += 1;
        }
    }
    
    special_count
}

pub fn min_element(nums: Vec<i32>) -> i32 {
    // 3300
    let mut min_digit_sum = 36;  // 9999 for 1 <= nums[i] <= 10^4
    for num in nums {
        let mut tmp_num = num;
        let mut digit_sum = 0;
        while tmp_num > 0 {
            digit_sum += tmp_num % 10;
            tmp_num /= 10;
        }
        min_digit_sum = min_digit_sum.min(digit_sum)
    }
    min_digit_sum
}

pub fn asteroids_destroyed_long(mass: i32, mut asteroids: Vec<i32>) -> bool {
    // 2126
    asteroids.sort();
    let mut new_mass = mass as i64;
    for asteroid_mass in asteroids {
        if asteroid_mass as i64 > new_mass { return false }
        new_mass += asteroid_mass as i64;
    }
    true
}

pub fn asteroids_destroyed(mass: i32, mut asteroids: Vec<i32>) -> bool {
    // 2126
    asteroids.sort();
    asteroids.iter().try_fold(mass as i64, |acc, &ast| {
        (acc >= ast as i64).then(|| acc + ast as i64)
    }).is_some()
}