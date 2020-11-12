package main

import "fmt"

var (
	registers [16]int
	I         int
	Mem       = map[int]int{}
)

func init() {
	Mem[0x2ea] = 0x80
	Mem[0x2eb] = 0x80
	Mem[0x2ec] = 0x80
	Mem[0x2ed] = 0x80
	Mem[0x2ee] = 0x80
	Mem[0x2ef] = 0x80
}

func main() {
	setVxNN(0xa, 02)   // 0x00
	setVxNN(0xb, 0x0c) // 0x02
	setVxNN(0xc, 0x3f) // 0x04
	setVxNN(0xd, 0x0c) // 0x06

	setItoAddr(0x2ea) // 0x08
	draw(0xa, 0xb, 6) // 0x0a draw the left paddle
	draw(0xc, 0xd, 6) // 0x0c draw the right paddle

	setVxNN(0xe, 0) // 0x0e

	subroutine2d4() // 0x10 Call subroutine at 0x2d4

	setVxNN(6, 3) // 0x12
	setVxNN(8, 2)
	setVxNN(6, 60)

	setDelayTimer

	fmt.Printf("regs: %v, I: %04x, M:%x", registers, I, Mem)
}

func subroutine2d4() {
	setItoAddr(0x2f2) // 0xd4
	storeXBCD(0xe)    //0xd6
	loadVxFromMem(2)  // 0xd8 loads the BCD numbers (just the 100 and 10) into reg 0 and 1
	setItoChar(1)     // 0xda set I to memory location of sprite for the character stored in reg 1
	setVxNN(6, 14)    // 0xdc
	setVxNN(5, 00)    // 0xde

	draw(4, 5, 5)  // 0xe0
	addToVx(4, 15) // 0xe2

	setItoChar(2) // 0xe4

	draw(4, 5, 5)

	// 0xe8 return from sub routine
}

func addToVx(reg, i int) {
	registers[reg] += i
}

func setItoChar(reg int) {
	I = getMemLocationForChar(registers[reg])
}

// This will return the actual mem location 0 = 0, 1 = 5, 2 = 10 etc
func getMemLocationForChar(character int) int {
	return character * 5
}

func loadVxFromMem(reg int) {
	for i := 0; i < reg; i++ {
		registers[i] = Mem[I+i]
	}
}

func storeXBCD(reg int) {
	Mem[I] = registers[reg] / 100
	Mem[I+1] = (registers[reg] / 10) % 10
	Mem[I+2] = (registers[reg] % 100) % 10
}

func setVxNN(register int, value int) {
	registers[register] = int(value)
}

func setItoAddr(memoryAddress int) {
	I = memoryAddress
}

func draw(xReg, yReg, height int) {
	x := registers[xReg]
	y := registers[yReg]

	fmt.Println(x, y)

	for i := 0; i < height; i++ {
		fmt.Printf("%x\n", Mem[I+i])
	}

}
