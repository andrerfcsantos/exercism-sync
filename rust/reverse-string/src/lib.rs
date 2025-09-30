use unicode_reverse::reverse_grapheme_clusters_in_place;

pub fn reverse(input: &str) -> String {
    let mut string_copy: String = input.clone().to_string();
    reverse_grapheme_clusters_in_place(&mut string_copy);
    return string_copy;
}
