def transform(legacy_data: dict[int, list[str]]):
    result: dict[str, int] = {}
    for k, v in legacy_data.items():
        for n in v:
            result[n.lower()] = k
    return result