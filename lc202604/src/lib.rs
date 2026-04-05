pub fn is_leap_year(year: u64) -> bool {
    year.is_multiple_of(400) || (year.is_multiple_of(4) && !year.is_multiple_of(100))
}

pub fn decode_ciphertext_mem(encoded_text: String, rows: i32) -> String {
    let n = encoded_text.len();
    let rows = rows as usize;
    let cols = n / rows;
    let chars: Vec<char> = encoded_text.chars().collect();
    let mut result = String::new();
    for c in 0..cols {
        let mut row = 0;
        let mut col = c;
        // (0, c) -> (1, c+1) -> (2, c+2) ...
        while row < rows && col < cols {
            let index = row * cols + col;
            result.push(chars[index]);
            row += 1;
            col += 1;
        }
    }
    result.trim_end().to_string()
}

pub fn decode_ciphertext(encoded_text: String, rows: i32) -> String {
    // 2075
    let bytes = encoded_text.as_bytes(); // No collection/allocation!
    let rows = rows as usize;
    let cols = bytes.len() / rows;
    let mut result = Vec::new();
    for c in 0..cols {
        let mut r = 0;
        let mut current_col = c;
        // (0, c) -> (1, c+1) -> (2, c+2) ...
        while r < rows && current_col < cols {
            let index = r * cols + current_col;
            if index < bytes.len() {
                result.push(bytes[index]);
            }
            r += 1;
            current_col += 1;
        }
    }
    String::from_utf8_lossy(&result).trim_end().to_string()
}

pub fn judge_circle(moves: String) -> bool {
    // 657
    let (mut x, mut y) = (0, 0);
    for c in moves.chars() {
        match c {
            'R' => x += 1,
            'L' => x -= 1,
            'U' => y -= 1,
            'D' => y += 1,
            _ => return false,
        }
    }
    x == 0 && y == 0
}