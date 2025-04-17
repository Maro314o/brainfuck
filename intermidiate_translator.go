package main

import "log"

var (
	VALID_SYMBOLS = []Symbol{'+', '-', '>', '<', ',', '.', '[', ']'}
)

const (
	READ          = true
	WRITE         = false
	CELL_MODIFIER = iota
	NULL_TOKEN
	POINTER_MODIFIER
	IO_TOKEN
	CLOSING_BRACKET
	OPENING_BRACKET

	NONE
	BASIC
	COLLAPSED_VALUES
	FINISHED
)

type Token struct {
	enum_value int
	Cell_modifier_value
	Pointer_modifier_value
	IO_token
	Bracket
}
type Cell_modifier_value int
type Pointer_modifier_value int
type IO_token bool
type Bracket struct {
	Opening_bracket
	Closing_bracket
}
type Opening_bracket struct {
	companion_bracket_position int
}
type Closing_bracket struct {
	companion_bracket_position int
}

type Bf_source_code []Symbol
type Intermidiate_code struct {
	state       int
	code_string []Token
	code_lenght int
}
type Opening_bracket_stack struct {
	stack  []int
	lenght int
}

func (stack *Opening_bracket_stack) pop() int {
	stack.lenght--
	return stack.stack[stack.lenght]
}
func (stack *Opening_bracket_stack) push(position int) {
	if stack.lenght >= len(stack.stack) {
		stack.stack = append(stack.stack, position)
		stack.lenght++
		return
	}
	stack.stack[stack.lenght] = position
	stack.lenght++

}

func (code *Bf_source_code) remove_invalid_symbos(valid_symbol_set []Symbol) {
	var temp_source_code Bf_source_code
	for _, value := range *code {
		//magari implementare un set per ?velocizzare? questo passaggio
		for _, valid_symbol := range valid_symbol_set {
			if value == valid_symbol {
				temp_source_code = append(temp_source_code, value)
			}

		}
	}
	*code = temp_source_code

}

func (intermidiate_code *Intermidiate_code) collapse_value() {
	var code_string []Token

	offset := 0
	prev_token := NULL_TOKEN

	for index, token := range intermidiate_code.code_string {
		if index != 0 {
			switch token.enum_value {
			case CELL_MODIFIER:
      case POINTER_MODIFIER:
			case IO_TOKEN:
      case OPENING_BRACKET:
			case CLOSING_BRACKET:
				panic("")

			}
		} else {
			code_string = append(code_string,token)
		}

	}

	//offset the value of brackets
}
func (code *Bf_source_code) into_intermidiate_code() Intermidiate_code {
	var intermidiate_string []Token
	var value int
	var new_token Token
	var open_bracket_stack Opening_bracket_stack

	for position, symbol := range *code {
		new_token = Token{}
		switch symbol {
		case '+':
		case '-':
			new_token.enum_value = CELL_MODIFIER
			new_token.Cell_modifier_value = Cell_modifier_value(-(symbol - 44)) // ascii magic
			intermidiate_string = append(intermidiate_string, new_token)

		case '<':
		case '>':
			new_token.enum_value = POINTER_MODIFIER
			new_token.Pointer_modifier_value = Pointer_modifier_value((symbol - 60)) // ascii magic
			intermidiate_string = append(intermidiate_string, new_token)

		case ',':
		case '.':
			new_token.enum_value = IO_TOKEN
			new_token.IO_token = IO_token(symbol == ',')
			intermidiate_string = append(intermidiate_string, new_token)

		case '[':
			open_bracket_stack.push(position)
			new_token.enum_value = OPENING_BRACKET
			new_token.Opening_bracket.companion_bracket_position = 0
			intermidiate_string = append(intermidiate_string, new_token)

		case ']':

			value = open_bracket_stack.pop()
			new_token.enum_value = CLOSING_BRACKET
			new_token.Closing_bracket.companion_bracket_position = value
			intermidiate_string[value].Opening_bracket.companion_bracket_position = position
			intermidiate_string = append(intermidiate_string, new_token)
		}

	}
	if open_bracket_stack.lenght > 0 {
		log.Panic("invalid parentesis")
	}

	return Intermidiate_code{
		BASIC,
		intermidiate_string,
		len(intermidiate_string),
	}
}
func translate(original_source_code Bf_source_code) Intermidiate_code {

	var intermidiate_rappresentation Intermidiate_code
	intermidiate_rappresentation.state = NONE
	original_source_code.remove_invalid_symbos(VALID_SYMBOLS)

	return intermidiate_rappresentation
}
