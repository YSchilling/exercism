pub fn primes_up_to(upper_bound: u64) -> Vec<u64> {
    // create vec with numbers
    let mut unchecked_values: Vec<Option<u64>> = vec![None, None];
    let mut other_values: Vec<Option<u64>> = (2..=upper_bound).map(|x| Some(x)).collect();
    unchecked_values.append(&mut other_values);
    let mut primes: Vec<u64> = Vec::new();

    loop {
        // search for next unmarked number (next not null), if none finish
        let mut prime = None;
        for n in unchecked_values.as_slice() {
            if n.is_some() {
                prime = *n;
                break;
            }
        }
        if prime.is_none() {
            break;
        }
        // save it in primes
        primes.push(prime.unwrap());
        // go through the vec with unmarked number as step width
        for i in (0..=upper_bound.try_into().unwrap()).step_by(prime.unwrap().try_into().unwrap()) {
            unchecked_values[i] = None;
        }
    }

    primes
}
