struct Solution {}

impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        nums.retain(|x| x != &val);
        nums.len() as i32
    }
}

fn main() {
    let mut n = vec![0, 1, 2, 2, 3, 0, 4, 2];
    let a = Solution::remove_element(&mut n, 2);
    println!("{}", a);
}
