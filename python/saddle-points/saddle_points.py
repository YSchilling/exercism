def saddle_points(matrix):
    if len(matrix) == 0: return []
    row_length = len(matrix[0])
    for row in matrix:
        if len(row) != row_length: raise ValueError("irregular matrix")

    result_points = []

    for i, row in enumerate(matrix):
        for j in range(len(row)):
            current_element = matrix[i][j]
            col = [matrix[k][j] for k in range(len(matrix))]
            if current_element == max(row) and current_element == min(col):
                result_points.append({"row": i+1, "column": j+1})
    return result_points
