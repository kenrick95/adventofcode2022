package main

import (
	"sort"
	"strconv"
	"strings"
	"utils"
)

type File struct {
	name         string
	absolutePath string
	size         int
	parentDir    *Directory
}

type Directory struct {
	name         string
	absolutePath string
	files        []*File
	directories  []*Directory
	parentDir    *Directory
	totalSize    int
}

func main() {
	pwd := []string{}
	currentDir := &Directory{
		parentDir:    nil,
		absolutePath: "/",
		name:         "/",
		files:        []*File{},
		directories:  []*Directory{},
		totalSize:    0,
	}
	rootPath := currentDir
	dirMap := map[string]*Directory{}
	dirMap["/"] = rootPath

	lines := utils.ReadFileToLines("day07.in")

	lastCommand := ""
	for _, line := range lines {
		// fmt.Println("> ", line)
		if line == "" {
			continue
		} else if line[0] == '$' {
			// command
			cmd := line[2:]
			if cmd[0:2] == "ls" {
				lastCommand = "ls"
				continue
			} else if cmd[0:2] == "cd" {
				cdTarget := cmd[3:]
				lastCommand = "cd"
				if cdTarget == "/" {
					pwd = []string{}
					currentDir = rootPath

				} else if cdTarget == ".." {
					// pop
					pwd = pwd[0 : len(pwd)-1]

					if currentDir.parentDir != nil {
						currentDir = currentDir.parentDir
					} else {
						currentDir = rootPath
					}

				} else {
					// push
					pwd = append(pwd, cdTarget)
					for _, dir := range currentDir.directories {
						if dir.name == cdTarget {
							currentDir = dir
							break
						}
					}

				}
			}
		} else if lastCommand == "ls" {
			// parse command output
			parts := strings.Split(line, " ")
			name := parts[1]

			absolutePathBase := ""
			if currentDir.absolutePath == "/" {
				absolutePathBase = "/"
			} else {
				absolutePathBase = currentDir.absolutePath + "/"
			}

			if parts[0] == "dir" {
				dir := Directory{
					name:         name,
					files:        []*File{},
					directories:  []*Directory{},
					parentDir:    currentDir,
					absolutePath: absolutePathBase + name,
					totalSize:    0,
				}
				dirMap[dir.absolutePath] = &dir
				currentDir.directories = append(currentDir.directories, &dir)
			} else {
				fileSize, _ := strconv.Atoi(parts[0])
				file := File{
					name:         name,
					absolutePath: absolutePathBase + name,
					size:         fileSize,
					parentDir:    currentDir,
				}
				currentDir.files = append(currentDir.files, &file)

				tmpDir := currentDir
				for tmpDir != nil {
					tmpDir.totalSize += fileSize

					tmpDir = tmpDir.parentDir
				}
			}

		}
	}

	// Part 1 (sample) totalSize: 48_381_165
	// Part 1 (real) totalSize: 43_562_874
	println("rootPath", rootPath.totalSize)

	// Part 1: Find dirs with totalSize <= 100_000
	{
		ansPart1 := visitPart1(rootPath)
		println("Part 1:", ansPart1)
	}

	// Part 2: Find smallest dir that's at >= freeSpaceDeficit

	{
		totalDiskSpace := 70_000_000
		currentFreeSpace := totalDiskSpace - rootPath.totalSize
		freeSpaceRequired := 30_000_000
		freeSpaceDeficit := freeSpaceRequired - currentFreeSpace
		allCandidates := []int{}
		for _, dir := range dirMap {
			if dir.totalSize >= freeSpaceDeficit {
				allCandidates = append(allCandidates, dir.totalSize)
			}
		}
		sort.Ints(allCandidates)

		println("Part 2:", allCandidates[0])
	}

}

func visitPart1(dir *Directory) int {
	sum := 0
	if dir.totalSize <= 100000 {
		sum += dir.totalSize
	}
	for _, nextDir := range dir.directories {
		sum += visitPart1(nextDir)
	}
	return sum
}
