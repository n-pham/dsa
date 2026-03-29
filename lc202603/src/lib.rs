use std::collections::HashSet;

pub fn min_partitions(n: String) -> i32 {
    // 1689
    // The answer is the maximum digit in the string.
    n.chars()
        .map(|c| (c as u8 - b'0') as i32)
        .max()
        .unwrap_or(0)
}

pub fn find_kth_bit(n: i32, k: i32) -> char {
    // 1545
    let mut chars: Vec<char> = vec!['0'];
    for _ in 1..n {
        let mut part2: Vec<char> = chars
            .iter()
            .rev()
            .map(|&c| if c == '0' { '1' } else { '0' })
            .collect();
        chars.push('1');
        chars.append(&mut part2);
    }
    chars[k as usize - 1]
}

pub fn num_special(mat: Vec<Vec<i32>>) -> i32 {
    // 1582
    let m = mat.len();
    let n = mat[0].len();

    // Count ones in each row and column
    let mut row_count = vec![0; m];
    let mut col_count = vec![0; n];

    for i in 0..m {
        for j in 0..n {
            if mat[i][j] == 1 {
                row_count[i] += 1;
                col_count[j] += 1;
            }
        }
    }

    // Check for special positions
    let mut result = 0;
    for i in 0..m {
        for j in 0..n {
            if mat[i][j] == 1 && row_count[i] == 1 && col_count[j] == 1 {
                result += 1;
            }
        }
    }

    result
}

pub fn min_operations(s: String) -> i32 {
    // 1758
    let mut count_if_0_first = 0;
    for (i, c) in s.chars().enumerate() {
        if i % 2 == 0 {
            if c != '0' {
                count_if_0_first += 1;
            }
        } else {
            if c != '1' {
                count_if_0_first += 1;
            }
        }
    }
    let count_if_1_first = s.len() as i32 - count_if_0_first;
    count_if_0_first.min(count_if_1_first)
}

pub fn check_ones_segment_fail(s: String) -> bool {
    // 1784
    let mut prev_is_1 = false;
    for char in s.chars() {
        if char == '1' && prev_is_1 {
            return true;
        }
        prev_is_1 = char == '1'
    }
    false
}

pub fn check_ones_segment(s: String) -> bool {
    // 1784
    !s.contains("01")
}

pub fn min_flips(s: String) -> i32 {
    // 1888
    let n = s.len();
    if n == 0 {
        return 0;
    }
    let t = s.repeat(2);
    let bytes = t.as_bytes();
    let mut even0 = 0;
    let mut even1 = 0;
    let mut odd0 = 0;
    let mut odd1 = 0;

    // initial window: indices 0..n
    for idx in 0..n {
        let b = bytes[idx]; // FIX: use square brackets
        if idx % 2 == 0 {
            if b == b'0' {
                even0 += 1;
            } else {
                even1 += 1;
            }
        } else {
            if b == b'0' {
                odd0 += 1;
            } else {
                odd1 += 1;
            }
        }
    }

    let p0_initial = 0 % 2; // always 0
    let flips0_initial = if p0_initial == 0 {
        even1 + odd0
    } else {
        odd1 + even0
    };
    let flips1_initial = if p0_initial == 0 {
        even0 + odd1
    } else {
        odd0 + even1
    };
    let mut ans = flips0_initial.min(flips1_initial);

    // iterate i from 1 to n-1
    for i in 1..n {
        // remove old index i-1
        let old = bytes[i - 1]; // FIX
        let old_parity = (i - 1) % 2;
        if old == b'0' {
            if old_parity == 0 {
                even0 -= 1;
            } else {
                odd0 -= 1;
            }
        } else {
            if old_parity == 0 {
                even1 -= 1;
            } else {
                odd1 -= 1;
            }
        }

        // add new index i+n-1
        let new_idx = i + n - 1;
        let newb = bytes[new_idx]; // FIX
        let new_parity = new_idx % 2;
        if newb == b'0' {
            if new_parity == 0 {
                even0 += 1;
            } else {
                odd0 += 1;
            }
        } else {
            if new_parity == 0 {
                even1 += 1;
            } else {
                odd1 += 1;
            }
        }

        let p0 = i % 2;
        let flips0 = if p0 == 0 { even1 + odd0 } else { odd1 + even0 };
        let flips1 = if p0 == 0 { even0 + odd1 } else { odd0 + even1 };
        ans = ans.min(flips0.min(flips1));
    }

    ans
}

pub fn find_different_binary_string(nums: Vec<String>) -> String {
    // 1980
    let n = nums.len();
    let int_set: HashSet<i32> = nums
        .iter()
        .map(|s| i32::from_str_radix(s, 2).expect("Not a valid binary number!"))
        .collect();
    // Search all n-bit numbers
    for num in 0..(1 << n) {
        if !int_set.contains(&num) {
            return format!("{:0width$b}", num, width = n);
        }
    }
    unreachable!("There must always be a solution");
}

pub fn number_of_stable_arrays(zero: i32, one: i32, limit: i32) -> i32 {
    // 3130
    let mod_val = 1_000_000_007;
    let z = zero as usize;
    let o = one as usize;
    let lim = limit as usize;

    let mut dp = vec![vec![[0i32; 2]; o + 1]; z + 1];

    for i in 1..=z.min(lim) {
        dp[i][0][0] = 1;
    }
    for j in 1..=o.min(lim) {
        dp[0][j][1] = 1;
    }

    for i in 1..=z {
        for j in 1..=o {
            dp[i][j][0] = (dp[i - 1][j][0] + dp[i - 1][j][1]) % mod_val;
            if i > lim {
                let sub = dp[i - lim - 1][j][1];
                dp[i][j][0] = (dp[i][j][0] - sub + mod_val) % mod_val;
            }

            dp[i][j][1] = (dp[i][j - 1][0] + dp[i][j - 1][1]) % mod_val;
            if j > lim {
                let sub = dp[i][j - lim - 1][0];
                dp[i][j][1] = (dp[i][j][1] - sub + mod_val) % mod_val;
            }
        }
    }

    (dp[z][o][0] + dp[z][o][1]) % mod_val
}

pub fn bitwise_complement(n: i32) -> i32 {
    // 1009
    if n == 0 {
        return 1;
    }
    let bit_length = 32 - n.leading_zeros();
    let mask = (1 << bit_length) - 1;
    n ^ mask
}

const MOD: i64 = 1_000_000_007;

fn power(mut x: i64, mut y: i64) -> i64 {
    let mut res = 1;
    x %= MOD;
    while y > 0 {
        if y % 2 == 1 {
            res = (res * x) % MOD;
        }
        y /= 2;
        x = (x * x) % MOD;
    }
    res
}

fn mod_inverse(n: i64) -> i64 {
    power(n, MOD - 2)
}

#[derive(Default)]
pub struct Fancy {
    nums: Vec<i64>,
    a: i64,
    b: i64,
}

impl Fancy {
    // 1622
    pub fn new() -> Self {
        Self {
            nums: Vec::new(),
            a: 1,
            b: 0,
        }
    }

    pub fn append(&mut self, val: i32) {
        let inv_a = mod_inverse(self.a);
        let val = val as i64;
        let v = ((val - self.b + MOD) % MOD * inv_a) % MOD;
        self.nums.push(v);
    }

    pub fn add_all(&mut self, inc: i32) {
        self.b = (self.b + inc as i64) % MOD;
    }

    pub fn mult_all(&mut self, m: i32) {
        self.a = (self.a * m as i64) % MOD;
        self.b = (self.b * m as i64) % MOD;
    }

    pub fn get_index(&self, idx: i32) -> i32 {
        let i = idx as usize;
        if i >= self.nums.len() {
            return -1;
        }
        let res = (self.nums[i] * self.a + self.b) % MOD;
        res as i32
    }
}

pub fn largest_submatrix(matrix: Vec<Vec<i32>>) -> i32 {
    // 1727
    let num_cols = matrix[0].len();
    let mut max_area = 0;

    let mut col_running_sums = vec![0; num_cols];
    for row in &matrix {
        for c in 0..num_cols {
            if row[c] == 1 {
                col_running_sums[c] += 1;
            } else {
                col_running_sums[c] = 0;
            }
        }

        let mut sorted_heights = col_running_sums.clone();
        sorted_heights.sort_by(|a, b| b.cmp(a));

        for (i, &h) in sorted_heights.iter().enumerate() {
            let width = (i + 1) as i32;
            max_area = max_area.max(h * width);
        }
    }
    max_area
}

pub fn recite(start_bottles: u32, take_down: u32) -> String {
    let lower_words = [
        "no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten",
    ];
    let upper_words = [
        "No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
    ];
    let lines: Vec<String> = (0..take_down)
        .map(|i| {
            let current = (start_bottles - i) as usize;
            let remaining = (start_bottles - i - 1) as usize;
            let current_suffix = if current == 1 { "bottle" } else { "bottles" };
            let remaining_suffix = if remaining == 1 { "bottle" } else { "bottles" };
            format!(
                "{} green {} hanging on the wall,\n\
            {} green {} hanging on the wall,\n\
            And if one green bottle should accidentally fall,\n\
            There'll be {} green {} hanging on the wall.",
                upper_words[current],
                current_suffix,
                upper_words[current],
                current_suffix,
                lower_words[remaining],
                remaining_suffix
            )
        })
        .collect();
    lines.join("\n\n")
}

pub fn square_of_sum(n: u32) -> u32 {
    n * (n + 1) * n * (n + 1) / 4
}

pub fn sum_of_squares(n: u32) -> u32 {
    (1..=n).map(|x| x * x).sum()
}

pub fn difference(n: u32) -> u32 {
    square_of_sum(n) - sum_of_squares(n)
}

pub fn square(s: u32) -> u64 {
    1 << (s - 1)
}

pub fn total() -> u64 {
    (1..=64).map(square).sum()
}

pub fn are_similar(mat: Vec<Vec<i32>>, k: i32) -> bool {
    // 2946
    let mut shifted = mat.clone();
    let k = k as usize;
    for (i, row) in shifted.iter_mut().enumerate() {
        let shift_amount = k % row.len();
        if i % 2 == 0 {
            row.rotate_left(shift_amount);
        } else {
            row.rotate_right(shift_amount);
        }
    }
    mat == shifted
}

pub fn can_be_equal(s1: String, s2: String) -> bool {
    // 2839
    let mut c1 = s1.chars();
    let (c10, c11, c12, c13) = (c1.next(), c1.next(), c1.next(), c1.next());

    let mut c2 = s2.chars();
    let (c20, c21, c22, c23) = (c2.next(), c2.next(), c2.next(), c2.next());

    ((c10 == c20 && c12 == c22) || (c10 == c22 && c12 == c20))
        && ((c11 == c21 && c13 == c23) || (c11 == c23 && c13 == c21))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_fancy() {
        let mut fancy = Fancy::new();
        fancy.append(2);
        fancy.add_all(3);
        fancy.append(7);
        fancy.mult_all(2);
        assert_eq!(fancy.get_index(0), 10);
        fancy.add_all(3);
        fancy.append(10);
        fancy.mult_all(2);
        assert_eq!(fancy.get_index(0), 26);
        assert_eq!(fancy.get_index(1), 34);
        assert_eq!(fancy.get_index(2), 20);
    }

    #[test]
    fn test_bitwise_complement() {
        assert_eq!(bitwise_complement(5), 2);
        assert_eq!(bitwise_complement(7), 0);
        assert_eq!(bitwise_complement(10), 5);
        assert_eq!(bitwise_complement(0), 1);
    }

    #[test]
    fn test_num_special() {
        assert_eq!(
            num_special(vec![vec![1, 0, 0], vec![0, 0, 1], vec![1, 0, 0]]),
            1
        );
        assert_eq!(
            num_special(vec![vec![1, 0, 0], vec![0, 1, 0], vec![0, 0, 1]]),
            3
        );
    }

    #[test]
    fn test_check_ones_segment() {
        assert_eq!(check_ones_segment("1001".to_string()), false);
        assert_eq!(check_ones_segment("110".to_string()), true);
        assert_eq!(check_ones_segment("1".to_string()), true);
    }

    #[test]
    fn test_check_ones_segment_fail() {
        // This function seems to check for presence of "11"
        assert_eq!(check_ones_segment_fail("1001".to_string()), false);
        assert_eq!(check_ones_segment_fail("110".to_string()), true);
    }

    #[test]
    fn test_min_flips() {
        assert_eq!(min_flips("111000".to_string()), 2);
        assert_eq!(min_flips("010".to_string()), 0);
        assert_eq!(min_flips("111".to_string()), 1);
        assert_eq!(min_flips("".to_string()), 0);
    }

    #[test]
    fn test_find_different_binary_string() {
        let nums = vec!["01".to_string(), "10".to_string()];
        let res = find_different_binary_string(nums);
        assert!(res == "00" || res == "11");

        let nums2 = vec!["00".to_string(), "01".to_string()];
        let res2 = find_different_binary_string(nums2);
        assert!(res2 == "10" || res2 == "11");
    }

    #[test]
    fn test_number_of_stable_arrays() {
        assert_eq!(number_of_stable_arrays(1, 1, 2), 2);
        assert_eq!(number_of_stable_arrays(1, 2, 1), 1);
        assert_eq!(number_of_stable_arrays(3, 3, 2), 14);
    }

    #[test]
    fn test_min_operations() {
        assert_eq!(min_operations("0100".to_string()), 1);
        assert_eq!(min_operations("10".to_string()), 0);
        assert_eq!(min_operations("1111".to_string()), 2);
    }

    #[test]
    fn test_find_kth_bit() {
        assert_eq!(find_kth_bit(3, 1), '0');
        assert_eq!(find_kth_bit(4, 11), '1');
    }

    #[test]
    fn test_min_partitions() {
        assert_eq!(min_partitions("32".to_string()), 3);
        assert_eq!(min_partitions("82734".to_string()), 8);
        assert_eq!(min_partitions("27346209830709182346".to_string()), 9);
    }
}
