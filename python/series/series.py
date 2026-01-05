def slices(series, length):
    if length == 0:
        raise ValueError("slice length cannot be zero")
    
    if length < 0:
        raise ValueError("slice length cannot be negative")
    
    if len(series) == 0:
        raise ValueError("series cannot be empty")
    
    if length > len(series):
        raise ValueError("slice length cannot be greater than series length")
    
    results = []
    for i in range(0, len(series) - length + 1):
        results.append(series[i:i+length])
    
    return results