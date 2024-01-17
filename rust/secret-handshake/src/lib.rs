pub fn actions(n: u8) -> Vec<&'static str> {
    //create a iterable representations of the binary of n
    let bits: Vec<u8> = (0..5).map(|x| (n >> x) & 1).collect();
    let mut actions: Vec<&'static str> = Vec::new();

    //iterate over it and create the actions vec
    for (i, val) in bits.iter().enumerate() {
        match (i, val) {
            (0, 1) => actions.push("wink"),
            (1, 1) => actions.push("double blink"),
            (2, 1) => actions.push("close your eyes"),
            (3, 1) => actions.push("jump"),
            (4, 1) => actions.reverse(),
            (_, _) => (),
        }
    }

    //return it
    actions
}
