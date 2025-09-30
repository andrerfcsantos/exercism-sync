use std::collections::HashSet;
use std::collections::HashMap;
use unicode_segmentation::UnicodeSegmentation;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let mut anagrams: HashSet<&'a str> = HashSet::new();
    let word = word.to_lowercase();
    let word_anagram = make_anagram_map(&word);

    for &possible_anagram in possible_anagrams {
        let candidate_lowercase = possible_anagram.to_lowercase();
        let candidate_anagram = make_anagram_map(&candidate_lowercase);
        if word_anagram == candidate_anagram && candidate_lowercase != word {
            anagrams.insert(possible_anagram);
        }
    }

    return anagrams;
}

pub fn make_anagram_map(word: &str) -> HashMap<&str, i32> {
    let mut res = HashMap::new();

    for c in word.graphemes(true) {
        match res.get_mut(&c) {
            None => { res.insert(c, 1); }
            Some(v) => { *v += 1 }
        }
    }

    return res;
}
