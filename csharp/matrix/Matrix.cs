using System;

public class Matrix
{
    public int[][] matrix;
    public Matrix(string input)
    {
        string[] rows = input.Split('\n');
        string[][] matrix = [];

        for (int i = 0; i < rows.Length; i++)
        {
            matrix[i] = rows[i].Split(' ');
        }


        for (int i = 0; i < matrix.Length; i++)
        {
            for (int j = 0; j < matrix[0].Length; j++)
            {
                this.matrix[i][j] = Int32.Parse(matrix[i][j]);
            }
        }
    }

    public int[] Row(int row)
    {
        return matrix[row];
    }

    public int[] Column(int col)
    {
        throw new NotImplementedException("You need to implement this function.");
    }
}