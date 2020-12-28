package main

import (
	"fmt"
	"os"
)

//点
type point struct {
	i, j int
}

//相对于当前点的四周
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

//add 移动点
func (p point) add(r point) point {
	return point{
		p.i + r.i,
		p.j + r.j,
	}
}

//at 返回下一个点，判断点是否合法
func (p point) at(grid [][]int) (int, bool) {
	//行跑到墙外面去了
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	//列跑到墙外面去了
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	//没问题返回
	return grid[p.i][p.j], true
}

//读取迷宫并保存到slice中
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Println(row, col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

//walk 算法代码 maze迷宫数据  start起始点  end终点 返回结果：
func walk(maze [][]int, start, end point) [][]int {
	//steps 用于保存搜索数据
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
	}

	//待搜索队列 从起始开始
	queue := []point{start}

	//队列不空搜索
	for len(queue) > 0 {
		//取掉队首内容
		cur := queue[0]
		queue = queue[1:]

		//走到出口 退出
		if cur == end {
			break
		}

		//搜索四周
		for _, dir := range dirs {
			//找到下一个需要搜索的点，将下标add得到
			next := cur.add(dir)

			//判断点是否能走
			val, ok := next.at(maze)
			//撞墙
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			//走过了
			if !ok || val != 0 {
				continue
			}

			//跑到起点了
			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			//把探索到的点加1
			steps[next.i][next.j] = curSteps + 1
			//把这个点加到队列
			queue = append(queue, next)

		}

	}

	return steps
}

func main() {
	maze := readMaze("maze.in")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
