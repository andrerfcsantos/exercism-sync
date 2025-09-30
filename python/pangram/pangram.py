alphabet = set(['a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'])

def is_pangram(sentence):
    lowered = sentence.lower()
    only_alphabet = filter(lambda x: x in alphabet, lowered)
    chars_used = set("".join(only_alphabet))
    return len(chars_used) == len(alphabet)
