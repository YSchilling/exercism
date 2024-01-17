pub fn brackets_are_balanced(string: &str) -> bool {
    let mut stack = Vec::new();
    for ch in string.chars() {
        match ch {
            '(' | '{' | '[' => stack.push(ch),
            ')' | '}' | ']' => {
                if check_bracket(&mut stack, ch) == false {
                    return false;
                }
            }
            _ => (),
        }
    }

    if stack.len() != 0 {
        return false;
    }

    true
}

fn check_bracket(stack: &mut Vec<char>, ch: char) -> bool {
    let bracket = match stack.pop() {
        None => return false,
        Some(br) => match br {
            '(' => ')',
            '{' => '}',
            '[' => ']',
            _ => return false,
        },
    };

    bracket == ch
}
