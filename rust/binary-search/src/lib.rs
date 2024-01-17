pub fn find<T: Ord>(array: &[T], key: T) -> Option<usize> {
    if array.len() == 0 {
        return None;
    }

    let lower = 0;
    let upper = array.len() - 1;
    search(array, lower, upper, key)
}

fn search<T: Ord>(array: &[T], mut lower: usize, mut upper: usize, key: T) -> Option<usize> {
    let middle = (upper - lower) / 2 + lower;
    // break conditions
    if array[middle] == key {
        return Some(middle);
    }
    if lower == upper {
        return None;
    }

    // adjust bounds
    if array[middle] < key {
        lower = middle + 1;
    } else {
        upper = middle;
    }

    // recursive call
    search(array, lower, upper, key)
}
