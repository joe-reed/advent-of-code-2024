package main

import (
	"github.com/samber/lo"
	"slices"
	"sort"
	"strings"
	"utils"
)

func puzzle1(input string) (result int) {
	blocks := parseBlocks(input, true)

	lastFilled := 0

	for i := len(blocks) - 1; i >= lastFilled; i-- {
		if !blocks[i].isOccupied {
			continue
		}

		dest := attemptToMoveBlock(blocks, i)

		if dest != -1 {
			lastFilled = dest
		}
	}

	return blocks.checksum()
}

func puzzle2(input string) (result int) {
	blocks := parseBlocks(input, false)

	ids := lo.Filter(lo.Map(blocks, func(b Block, _ int) int {
		return b.id
	}), func(id int, _ int) bool {
		return id != -1
	})

	sort.Slice(ids, func(i, j int) bool {
		return ids[i] > ids[j]
	})

	for _, id := range ids {
		for i := 0; i < len(blocks); i++ {
			if blocks[i].id != id {
				continue
			}

			attemptToMoveBlock(blocks, i)
		}
	}

	return blocks.checksum()
}

func attemptToMoveBlock(blocks Blocks, i int) int {
	for j := 0; j < len(blocks); j++ {
		if i <= j {
			return -1
		}
		if !blocks[j].isOccupied && blocks[j].length >= blocks[i].length {
			blocks.moveBlock(i, j)
			return j
		}
	}

	return -1
}

type Block struct {
	isOccupied bool
	id         int
	length     int
}

type Blocks []Block

func (blocks Blocks) checksum() (result int) {
	idx := 0
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < blocks[i].length; j++ {
			if blocks[i].isOccupied {
				result += blocks[i].id * idx
			}
			idx++
		}
	}

	return
}

func (blocks Blocks) moveBlock(src int, dest int) {
	srcBlock := &blocks[src]
	destBlock := &blocks[dest]

	srcLength := srcBlock.length
	destLength := destBlock.length

	destBlock.copy(srcBlock)
	srcBlock.empty()

	if destLength > srcLength {
		slices.Insert(blocks, dest+1, Block{isOccupied: false, id: -1, length: destLength - srcLength})
	}
}

func (b *Block) copy(c *Block) {
	b.id = c.id
	b.isOccupied = c.isOccupied
	b.length = c.length
}

func (b *Block) empty() {
	b.id = -1
	b.isOccupied = false
}

func parseBlocks(input string, isSingleBlock bool) (blocks Blocks) {
	diskMap := utils.MapToInts(strings.Split(input, ""))
	currentId := 0

	for i, digit := range diskMap {
		id := -1
		isOccupied := false
		if i%2 == 0 {
			id = currentId
			isOccupied = true

			currentId++
		}

		if isSingleBlock {
			for j := 0; j < digit; j++ {
				blocks = append(blocks, Block{isOccupied: isOccupied, id: id, length: 1})
			}
		} else {
			blocks = append(blocks, Block{isOccupied: isOccupied, id: id, length: digit})
		}
	}

	return
}
