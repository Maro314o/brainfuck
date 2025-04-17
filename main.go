//TODO
/*
implementare "cache" per velocizzare i salti tra le parentesi
implementare un sistema di errori
implementare la visualizzazione della memoria (altro file)
implementare supporto per i commenti

implementare una versione alternativa di brainfuck con import ad altri file e "variabili" in modo da fare riferimento ad una lunga espressione con un solo nome


*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	NEGATIVE = true
	POSITIVE = false
)

type Symbol byte
type Instructions []Symbol

type Memory_vec struct {
	memory         []byte
	memory_pointer int
}
type Interpreter struct {
	instructions        Instructions
	instruction_pointer int
	memory              Memory_vec
}

func (interpreter *Interpreter) load_instructions(instructions Instructions) {
	interpreter.instructions = instructions
}

func (memory *Memory_vec) get_address_of_pointed_location(location int) *byte{
	memory_len := len(memory.memory)
	memory_pointer := location
	if memory_pointer >= memory_len {
		memory.memory = append(memory.memory, make([]byte, memory_pointer-memory_len+1)...)
	}
	return &memory.memory[location]
}
func (memory *Memory_vec) get_address_of_current_pointed_location() *byte {
	return memory.get_address_of_pointed_location(memory.memory_pointer)
}
/*
func (memory *Memory_vec) get_value_of_pointed_location(pointed_location int) byte{
	return *memory.get_address_of_pointed_location(pointed_location)
}
*/
func (memory *Memory_vec) get_value_of_current_pointed_location() byte {
	return *memory.get_address_of_current_pointed_location()
}
func (memory *Memory_vec) output_pointed_location() {
	value := memory.get_value_of_current_pointed_location()
	fmt.Print(value)
}
func (memory *Memory_vec) input_pointed_location() {
	memory_addr := memory.get_address_of_current_pointed_location()
  _, err := fmt.Scanf("%d", memory_addr)
  if err != nil {
    log.Panic("Please input only a valid number in this range [0 - 255]")
  }
  
}

func (memory *Memory_vec) add_signed_value_to_position(value byte, sign bool) {
	memory_position := memory.get_address_of_current_pointed_location()
	if sign == NEGATIVE {
		*memory_position -= value
	} else {
		*memory_position += value
	}
}
func (memory *Memory_vec) add_signed_value_to_memory_pointer(value int) {
	memory.memory_pointer += value
}
func (memory *Memory_vec) add_value_to_position(value byte) {
	memory.add_signed_value_to_position(value, POSITIVE)
}
func (memory *Memory_vec) subtract_value_to_position(value byte) {
	memory.add_signed_value_to_position(value, NEGATIVE)
}

func (interpreter *Interpreter) get_current_corrisponding_bracket_position(direction int) int {
	traversal_instruction_pointer :=interpreter.instruction_pointer+ direction
	opening_bracket_counter, closing_bracket_counter := 0, 0
	if direction == -1 {
		closing_bracket_counter = 1
	} else {
		opening_bracket_counter = 1
	}
	var current_token Symbol
	for opening_bracket_counter != closing_bracket_counter {
		current_token = interpreter.instructions[traversal_instruction_pointer]
		if current_token == '[' {
			opening_bracket_counter++
		} else if current_token == ']' {
			closing_bracket_counter++
		}
		traversal_instruction_pointer += direction
	}
	return traversal_instruction_pointer - direction
}
func (interpreter *Interpreter) get_current_corrisponding_open_bracket_position() int {
	return interpreter.get_current_corrisponding_bracket_position(-1)
}

func (interpreter *Interpreter) get_current_corrisponding_close_bracket_position() int {
	return interpreter.get_current_corrisponding_bracket_position(1)
}
func (interpreter *Interpreter) exec_symbol(symbol Symbol) {
	switch symbol {
	case '+':
		interpreter.memory.add_value_to_position(1)
	case '-':
		interpreter.memory.subtract_value_to_position(1)
	case ',':
		interpreter.memory.input_pointed_location()
	case '.':
		interpreter.memory.output_pointed_location()
	case '<':
		interpreter.memory.add_signed_value_to_memory_pointer(-1)
	case '>':
		interpreter.memory.add_signed_value_to_memory_pointer(1)
	case '[':
    if interpreter.memory.get_value_of_current_pointed_location() != 0{
      return 
    }
    interpreter.instruction_pointer = interpreter.get_current_corrisponding_close_bracket_position()

	case ']':
    if interpreter.memory.get_value_of_current_pointed_location() == 0{
      return 
    }
    interpreter.instruction_pointer = interpreter.get_current_corrisponding_open_bracket_position()

	default:
		log.Panic("invalid symbol : ", symbol)
	}

}
func (interpreter *Interpreter) excel_current_symbol(){
  interpreter.exec_symbol(interpreter.instructions[interpreter.instruction_pointer])
}
func main() {

	args := os.Args
	if len(args) > 2 {
		log.Panic("too many parameters")
	}
	file_bytes, err := os.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
	var bf_interprer Interpreter
	file := strings.ReplaceAll(string(file_bytes), "\n", "")
	file = strings.ReplaceAll(file, " ", "")
	bf_interprer.load_instructions(Instructions(file))
	for bf_interprer.instruction_pointer < len(bf_interprer.instructions) {
    bf_interprer.excel_current_symbol()
    bf_interprer.instruction_pointer++
	}

}
