/// Determine whether a sentence is a pangram.
pub fn is_pangram(sentence: &str) -> bool {
    let mut letters: Vec<char> = ('a'..='z').collect();

    let sentence = sentence.to_lowercase();
    for character in sentence.chars() {
        let result = letters.binary_search(&character);
        if result.is_ok() {
            letters.remove(result.unwrap());
        }
    }

    letters.len() == 0
}
