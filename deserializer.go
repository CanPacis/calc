package main

import (
	"errors"
)

type Deserialized struct {
	Version      uint32
	ConstantPool ConstantPool
	Instructions []Instruction
}

type Deserializer struct {
	input  []byte
	offset int
}

func (d Deserializer) Deserialize() (*Deserialized, error) {
	deserialized := &Deserialized{}

	if !d.validate_archive() {
		return nil, errors.New("invalid archive")
	}

	deserialized.Version = d.deserialize_version()

	pool, err := d.deserialize_constant_pool()
	if err != nil {
		return nil, err
	}
	deserialized.ConstantPool = pool

	instructions, err := d.deserialize_instructions()
	if err != nil {
		return nil, err
	}
	deserialized.Instructions = instructions

	return deserialized, nil
}

func (d *Deserializer) slice(n int) []byte {
	return d.input[d.offset : d.offset+n]
}

func (d *Deserializer) validate_archive() bool {
	validator_str := "calc.arc"
	archive := d.slice(len(validator_str))

	valid := string(archive) == validator_str
	if valid {
		d.offset += len(validator_str)
	}
	return valid
}

func (d *Deserializer) deserialize_version() uint32 {
	version := bytes_to_uint32(d.slice(4))
	d.offset += 4
	return version
}

func (d *Deserializer) deserialize_constant_pool() (ConstantPool, error) {
	pool := ConstantPool{}
	size := bytes_to_uint32(d.slice(4))
	d.offset += 4

	if size%8 != 0 {
		return pool, errors.New("broken archive")
	}

	for i := 0; i < int(size); i += 8 {
		value := d.slice(8)
		pool.Add(Float64Object{bytes_to_float64(value)})
		d.offset += 8
	}

	return pool, nil
}

func (d *Deserializer) deserialize_instructions() ([]Instruction, error) {
	instructions := []Instruction{}

	for d.offset < len(d.input) {
		instruction, err := d.deserialize_instruction()
		if err != nil {
			return instructions, err
		}
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func (d *Deserializer) deserialize_instruction() (Instruction, error) {
	size_b := d.slice(4)
	d.offset += 4
	size := bytes_to_uint32(size_b)

	if (size-1)%4 != 0 {
		return Instruction{}, errors.New("broken archive")
	}

	op := d.slice(1)
	d.offset += 1

	operands := []int{}

	for i := 0; i < int(size-1); i += 4 {
		value := d.slice(4)
		operands = append(operands, int(bytes_to_uint32(value)))
		d.offset += 4
	}

	return NewInstruction(Op(op[0]), operands...), nil
}

func NewDeserializer(input []byte) *Deserializer {
	return &Deserializer{
		input: input,
	}
}
