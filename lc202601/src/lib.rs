use std::collections::HashMap;
use std::collections::HashSet;

pub fn can_construct(ransom_note: &str, magazine: &str) -> bool {
    let mut counts = [0; 26];
    for b in magazine.bytes() {
        let idx = (b - b'a') as usize;
        counts[idx] += 1;
    }
    for b in ransom_note.bytes() {
        let idx = (b - b'a') as usize;
        if counts[idx] <= 0 {
            return false;
        }
        counts[idx] -= 1;
    }
    true
}

pub fn can_construct_failed(ransom_note: &str, magazine: &str) -> bool {
    let unique_chars: HashSet<char> = magazine.chars().collect();
    ransom_note.chars().all(|c| unique_chars.contains(&c))
}

pub fn can_construct_map(ransom_note: &str, magazine: &str) -> bool {
    let mut counts = HashMap::new();
    for char_val in magazine.chars() {
        *counts.entry(char_val).or_insert(0) += 1;
    }
    for char_val in ransom_note.chars() {
        match counts.entry(char_val) {
            std::collections::hash_map::Entry::Occupied(mut entry) => {
                let val = entry.get_mut();
                if *val <= 0 {
                    return false;
                }
                *val -= 1;
            }
            std::collections::hash_map::Entry::Vacant(_) => {
                return false;
            }
        }
    }
    true
}

pub fn repeated_n_times(nums: Vec<i32>) -> i32 {
    let mut seen = HashSet::new();
    for num in nums {
        if seen.contains(&num) {
            return num;
        }
        seen.insert(num);
    }
    // According to the problem description, there will always be an N-repeated element,
    // so this line should theoretically not be reached.
    // However, Rust requires a return value for all paths.
    0
}

pub fn sum_four_divisors(nums: Vec<i32>) -> i32 {
    let mut total_sum = 0;
    for num in nums {
        if num < 6 {
            continue;
        }
        let mut count = 2; // 1 and num are always divisors for num > 1
        let mut sum = 1 + num;
        let root = (num as f64).sqrt() as i32;
        for i in 2..=root {
            if num % i == 0 {
                if i * i == num {
                    count += 1;
                    sum += i;
                } else {
                    count += 2;
                    sum += i + num / i;
                }
            }
            if count > 4 {
                break;
            }
        }
        if count == 4 {
            total_sum += sum;
        }
    }
    total_sum
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_sum_four_divisors() {
        assert_eq!(sum_four_divisors(vec![21, 4, 7]), 32);
        assert_eq!(sum_four_divisors(vec![21, 21]), 64);
        assert_eq!(sum_four_divisors(vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10]), 12 + 15 + 18); // 6 (1,2,3,6 sum 12), 8 (1,2,4,8 sum 15), 10 (1,2,5,10 sum 18)
    }

    #[test]
    fn test_can_construct_failed() {
        assert!(!can_construct_failed("a", "b"));
        assert!(can_construct_failed("aa", "ab"));
        assert!(can_construct_failed("aa", "aab")); // Should be false because `unique_chars` only checks presence, not count.
        assert!(can_construct_failed("a", "a"));
        assert!(can_construct_failed("", "abc")); // Empty ransom note should always be true
    }

    #[test]
    fn test_can_construct_map() {
        // False cases
        assert!(!can_construct_map("a", "b"));
        assert!(!can_construct_map("aa", "ab"));
        assert!(!can_construct_map("flee", "foobar"));
        assert!(!can_construct_map("apple", "aple"));
        assert!(!can_construct_map("test", ""));

        // True cases
        assert!(can_construct_map("a", "a"));
        assert!(can_construct_map("aa", "aba"));
        assert!(can_construct_map("foobar", "foobarfoobar"));
        assert!(can_construct_map("abc", "abcdef"));
        assert!(can_construct_map("", "abc"));
        assert!(can_construct_map("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"));
    }

    #[test]
    fn test_can_construct() {
        // False cases
        assert!(!can_construct("a", "b"));
        assert!(!can_construct("aa", "ab"));
        assert!(!can_construct("flee", "foobar"));
        assert!(!can_construct("apple", "aple"));
        assert!(!can_construct("test", ""));

        // True cases
        assert!(can_construct("a", "a"));
        assert!(can_construct("aa", "aba"));
        assert!(can_construct("foobar", "foobarfoobar"));
        assert!(can_construct("abc", "abcdef"));
        assert!(can_construct("", "abc"));
        assert!(can_construct("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"));
    }

    #[test]
    fn test_repeated_n_times() {
        assert_eq!(repeated_n_times(vec![1, 2, 3, 3]), 3);
        assert_eq!(repeated_n_times(vec![2, 1, 2, 5, 3, 2]), 2);
        assert_eq!(repeated_n_times(vec![5, 1, 5, 2, 5, 3, 5, 4]), 5);
        assert_eq!(repeated_n_times(vec![1, 1]), 1);
    }
}