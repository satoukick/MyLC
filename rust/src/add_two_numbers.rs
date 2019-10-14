// Definition for singly-linked list.
#![allow(dead_code)]

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

struct Solution;

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}
impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        Self::add(l1, l2, 0)
    }

    fn add(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
        carry: i32,
    ) -> Option<Box<ListNode>> {
        if l1.is_none() && l2.is_none() && carry == 0 {
            return None;
        }

        let (a, x1) = Self::extract(l1);
        let (b, n2) = Self::extract(l2);

        let sum = a + b + carry;
        let new_carry = sum / 10;

        let mut node = ListNode::new(sum % 10);
        node.next = Self::add(x1, n2, new_carry);
        Some(Box::new(node))
    }

    #[inline]
    fn extract(l: Option<Box<ListNode>>) -> (i32, Option<Box<ListNode>>) {
        if let Some(n) = l {
            (n.val, n.next)
        } else {
            (0, None)
        }
    }
}

// fn main() {
//     println!("Hello");
//     let l1 = ListNode::new(1);
// }
