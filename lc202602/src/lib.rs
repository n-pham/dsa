pub fn sort_by_bits(mut arr: Vec<i32>) -> Vec<i32> {
    arr.sort_by_cached_key(|&num| (num.count_ones(), num));
    arr
}

pub fn num_steps(s: String) -> i32 {
    // 1404
    // s cannot be converted to very big int
    let mut steps = 0;
    let mut carry = 0;
    let bytes = s.as_bytes();
    for i in (1..bytes.len()).rev() {
        let current_bit = (bytes[i] - b'0') as i32;
        if current_bit + carry == 1 {
            steps += 2;
            carry = 1;
        } else {
            steps += 1;
        }
    }
    steps + carry
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_sort_by_bits() {
        assert_eq!(
            sort_by_bits(vec![0, 1, 2, 3, 4, 5, 6, 7, 8]),
            vec![0, 1, 2, 4, 8, 3, 5, 6, 7]
        );
        assert_eq!(
            sort_by_bits(vec![1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1]),
            vec![1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024]
        );
    }
}
