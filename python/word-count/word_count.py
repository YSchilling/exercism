def count_words(sentence):
    sentence = sentence.lower()
    chars = []
    for char in sentence:
        if char.isalnum() or char == "'" or char == " ":
            chars.append(char)
        else:
            chars.append(" ")
    
    words = ("".join(chars)).split()
    words = map(lambda x: x.strip("'"), words)

    result = {}
    for word in words:
        if word == "":
            continue
        if word not in result.keys():
            result[word] = 1
        else:
            result[word] += 1
        
    return result