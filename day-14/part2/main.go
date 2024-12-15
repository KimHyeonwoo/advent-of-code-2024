package main

import (
	"fmt"
	"os"
	"strings"

	common "github.com/KimHyeonwoo/advent-of-code-2024/day-14"
)

func main() {
	robots, err := common.ParseInput("input")
	if err != nil {
		panic(err)
	}

	width := 101
	height := 103

	file, err := os.Create("tree.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for range 39 {
		for r := range robots {
			robots[r].Move(width, height)
		}
	}

	fmt.Fprint(file, "Loop Count:", 39, "\n")
	robotCount := make([][]int, height)
	for i := range robotCount {
		robotCount[i] = make([]int, width)
	}

	for r := range robots {
		robotCount[robots[r].PositionRow][robots[r].PositionCol]++
	}

	for row := range robotCount {
		for col := range robotCount[row] {
			if robotCount[row][col] > 0 {
				fmt.Fprint(file, "*")
			} else {
				fmt.Fprint(file, " ")
			}
		}
		fmt.Fprintln(file)
	}

	for loopCount := range height {
		for r := range robots {
			robots[r].MoveWidthTimes(width, height)
		}

		robotCount = make([][]int, height)
		for i := range robotCount {
			robotCount[i] = make([]int, width)
		}

		for r := range robots {
			robotCount[robots[r].PositionRow][robots[r].PositionCol]++
		}

		fmt.Fprint(file, "Loop Count:", 39+(loopCount+1)*width, "\n")
		for row := range robotCount {
			for col := range robotCount[row] {
				if robotCount[row][col] > 0 {
					fmt.Fprint(file, "*")
				} else {
					fmt.Fprint(file, " ")
				}
			}
			fmt.Fprintln(file)
		}
		fmt.Fprint(file, strings.Repeat("-", width), "\n")
	}

	// Loop Count:7412
	//                                            *   *
	//           *       *
	//     *
	//                                           *
	//                          *              *
	//            *               *
	//                                                                              *                  *
	//                  *
	//                                         *
	//
	//                               *                     *
	//*                           *                *
	//           *                                                                *           *
	//         *            *                               *
	//
	//   *        *    *                                          *    *
	//             *                                                 *
	//                                         *
	//                  *
	//
	//                                    *
	//
	//                *******************************                                       *            *
	//                *                             *      *   *
	//                *                             *
	//                *                             *                             *
	//*               *                             *
	//  *             *              *              *
	//       *        *             ***             *                            *
	//                *            *****            *                     *
	//                *           *******           *                  *                          *     *
	//         *      *          *********          *
	//                *            *****            *                 *
	//                *           *******           *
	//                *          *********          *    *
	//                *         ***********         *                    *
	//                *        *************        *                            *
	//     *          *          *********          *                                          *
	//                *         ***********         *
	//         *      *        *************        *
	//                *       ***************       *             *
	//                *      *****************      *
	//                *        *************        *                                                     *
	//                *       ***************       *
	//                *      *****************      *
	//                *     *******************     *       *                *
	//                *    *********************    *
	//                *             ***             *                             **
	//                *             ***             *
	//                *             ***             *                 *
	//*       *       *                             *                             *
	//                *                             *                                          *
	//                *                             *
	//                *                             *
	//                *******************************                                     *
	//
	//                    *
	//                                                                                 *       *
	//
	//                                                                      *
	//                                             *                                                  *
	//                                            *              *           *       *
	//                                                                                    *
	//                                    *                                                            *
	//                                                                                          *
	//                    *
	//                             *                                                             *   *
	//                                  *                           *
	//                                *     *              *
	//                                              *          *  *
	//
	//            *                         *                        *                         *
	//                                                                 *
	//
	//                                                                  *
	//                       *                                              *
	//                                              *   *                                  *
	//    *               *
	//                                                                                                *
	//
	//                                        *                                               *
	//       *               *
	//
	//
	//  *
	//                                                *                           *
	//                                     *
	// *                                             *                            *
	//
	//         *                 *                                             *               *
	//                                                                         *                     *
	//
	//                                        *                   *                             *
	//                                   *                          *
	//                                                                          **
	//                       *
	//                                                                                     *
	//            *              *
	//  *             *         *
	//                                 *                                              *
	//                                                                                                    *
	//           *                                * *
	//             *    *
}
