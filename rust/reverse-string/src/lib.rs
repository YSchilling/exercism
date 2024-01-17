use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let mut output = String::new();

    for g in input.graphemes(true) {
        output.insert_str(0, g);
    }

    output
}
