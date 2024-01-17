use std::iter::FromIterator;

pub struct Node<T> {
    pub data: T,
    pub next: Option<Box<Node<T>>>,
}

impl<T> Node<T> {
    pub fn new(data: T) -> Self {
        Node { data, next: None }
    }
}

pub struct SimpleLinkedList<T> {
    head: Option<Box<Node<T>>>,
    length: usize,
}

impl<T> SimpleLinkedList<T> {
    pub fn new() -> Self {
        SimpleLinkedList {
            head: None,
            length: 0,
        }
    }

    // You may be wondering why it's necessary to have is_empty()
    // when it can easily be determined from len().
    // It's good custom to have both because len() can be expensive for some types,
    // whereas is_empty() is almost always cheap.
    // (Also ask yourself whether len() is expensive for SimpleLinkedList)
    pub fn is_empty(&self) -> bool {
        self.length == 0
    }

    pub fn len(&self) -> usize {
        self.length
    }

    pub fn push(&mut self, _element: T) {
        if self.is_empty() {
            self.head = Some(Box::new(Node::new(_element)));
            self.length += 1;
            return;
        }

        // get last node
        let mut pointer = &mut self.head;
        for _ in 0..self.length - 1 {
            pointer = &mut pointer.as_mut().unwrap().next;
        }

        pointer.as_mut().unwrap().next = Some(Box::new(Node::new(_element)));
    }

    pub fn pop(&mut self) -> Option<T> {
        if self.is_empty() {
            return None;
        }
        if self.length == 1 {
            let head = &self.head;
            self.head = None;
            return Some(head.unwrap().data);
        }

        // get node before last
        let mut pointer = &mut self.head;
        for _ in 0..self.length - 2 {
            pointer = &mut pointer.as_mut().unwrap().next;
        }
        let data = pointer.unwrap().next.unwrap().data;
        pointer.as_mut().unwrap().next = None;
        return Some(data);
    }

    pub fn peek(&self) -> Option<&T> {
        todo!()
    }

    #[must_use]
    pub fn rev(self) -> SimpleLinkedList<T> {
        todo!()
    }
}

impl<T> FromIterator<T> for SimpleLinkedList<T> {
    fn from_iter<I: IntoIterator<Item = T>>(_iter: I) -> Self {
        todo!()
    }
}

// In general, it would be preferable to implement IntoIterator for SimpleLinkedList<T>
// instead of implementing an explicit conversion to a vector. This is because, together,
// FromIterator and IntoIterator enable conversion between arbitrary collections.
// Given that implementation, converting to a vector is trivial:
//
// let vec: Vec<_> = simple_linked_list.into_iter().collect();
//
// The reason this exercise's API includes an explicit conversion to Vec<T> instead
// of IntoIterator is that implementing that interface is fairly complicated, and
// demands more of the student than we expect at this point in the track.

impl<T> From<SimpleLinkedList<T>> for Vec<T> {
    fn from(mut _linked_list: SimpleLinkedList<T>) -> Vec<T> {
        todo!()
    }
}
