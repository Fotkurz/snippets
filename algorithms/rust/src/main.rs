fn main() {
    println!("Hello, world!");
}

pub fn binary_search(value: i32, sorted: Vec<i32>) -> usize {
    let half = if sorted.len() % 2 == 0 {
        sorted.len() / 2
    } else {
        (sorted.len() - 1) / 2
    };
    let check = sorted[half];
    if check < value {
        binary_search(value, sorted[half:])
    } else if check > value {
        binary_search(value, sorted[:half])
    } else {
        half
    }
}

#[cfg(test)]
mod tests {
    use crate::binary_search;

    #[test]
    fn test_return_the_first_index_of_vec() {
        let got = binary_search(5, vec![5, 6, 7, 8, 9]);
        assert_eq!(0, got);
    }
}
