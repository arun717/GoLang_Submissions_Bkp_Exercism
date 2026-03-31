package chessboard
//import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File([]bool)

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard(map[string]File)

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	for letter, fileSlice := range cb {
        if letter == file {
            count := 0
            for _, boolValue := range fileSlice {
                if boolValue {
                    count++
                }
            }
            return count
        }
    }
    return 0
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
        return 0
    }
    count := 0
	for _, fileSlice := range cb {
        for i, occupied := range fileSlice {
            if i+1 == rank && occupied {
                count++
            }
        }       
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
    for _, fileSlice := range cb {
        for range fileSlice {
            count++
        }
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    count := 0
    letters := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
    for _,l := range letters {
        count += CountInFile(cb, l)
    }
    return count
}
