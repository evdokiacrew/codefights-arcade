/*
intro 08/60

After they became famous, the CodeBots all decided to move to a new building and live together.
The building is represented by a rectangular matrix of rooms.
Each cell in the matrix contains an integer that represents the price of the room.
Some rooms are free (their cost is 0), but that's probably because they are haunted, so all the bots are afraid of them.
That is why any room that is free or is located anywhere below a free room in the same column is not considered suitable for the bots to live in.
Help the bots calculate the total price of all the rooms that are suitable for them.
*/

func matrixElementsSum(matrix [][]int) int {
    sum := 0
    for c := 0; c < len(matrix[0]); c++ {
        for r := 0; r < len(matrix); r++ {
            if matrix[r][c] != 0 {
                sum += matrix[r][c]
            } else {
                break
            }
        }
    }
    return sum
}
