fn main() {
    println!("Hello");
}

#[warn(dead_code)]
/// recursive_sum
fn recursive_sum(v: Vec<i32>) -> i32 {
    let length = v.len();
    if length == 0 {
        return 0;
    }
    if length == 1 {
        return v.get(0).unwrap().clone();
    }
    v[0] + recursive_sum(v[1..].to_vec())
}

#[warn(dead_code)]
/// recursive_count count all elements of vec recursively
fn recursive_count(v: Vec<i32>) -> i32 {
    let length = v.len();
    if length == 0 {
        return 0;
    }
    if length == 1 {
        return 1;
    }

    1 + recursive_count(v[1..].to_vec())
}

#[warn(dead_code)]
fn recursive_max(v: Vec<i32>) -> i32 {
    let length = v.len();
    if length == 0 {
        return 0;
    }
    if length == 1 {
        return v[0];
    }

    let max = recursive_max(v[1..].to_vec());
    if v[0] > max {
        v[0]
    } else {
        max
    }
}

#[cfg(test)]
mod test {

    #[test]
    fn recursive_sum_of_len_0_vec_is_0() {
        let want: i32 = 0;
        let got = crate::recursive_sum(vec![]);

        assert_eq!(want, got);
    }

    #[test]
    fn recursive_sum_of_len_1_vec_is_first_ele() {
        let want: i32 = 10;
        let got = crate::recursive_sum(vec![10]);

        assert_eq!(want, got);
    }

    #[test]
    fn recursive_sum_of_vec() {
        let want: i32 = 12;
        let got = crate::recursive_sum(vec![2, 4, 6]);

        assert_eq!(want, got);
    }

    #[test]
    fn recursive_count_of_len_0_vec_is_0() {
        let want: i32 = 0;
        let got = crate::recursive_count(vec![]);

        assert_eq!(want, got);
    }

    #[test]
    fn recursive_count_of_len_1_vec_is_1() {
        let want: i32 = 1;
        let got = crate::recursive_count(vec![100]);

        assert_eq!(want, got);
    }

    #[test]
    fn recursive_count_of_vec() {
        let want: i32 = 5;
        let got = crate::recursive_count(vec![1, 2, 3, 4, 5]);

        assert_eq!(want, got);
    }

    #[test]
    fn recursive_max_of_len_0_is_0() {
        let want: i32 = 0;
        let got = crate::recursive_max(vec![]);

        assert_eq!(want, got);
    }

    #[test]
    fn recursive_max_of_len_1_is_ele() {
        let want: i32 = 5;
        let got = crate::recursive_max(vec![5]);

        assert_eq!(want, got);
    }

    #[test]
    fn recursive_max_of_vec() {
        let want: i32 = 7;
        let got = crate::recursive_max(vec![3, 5, 7, 2, 4]);

        assert_eq!(want, got);
    }
}