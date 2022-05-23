package main

type Root struct {
	Debug       bool      `yaml:"debug"`
	FirstBlock  Block     `yaml:"first_block,omitempty"`
	SecondBlock *Block    `yaml:"second_block,omitempty"`
	ThirdBlock  *AltBlock `yaml:"block_three,omitempty"`
}

type Block struct {
	Name  string `yaml:"name"`
	Count int    `yaml:"count"`
	Valid bool   `yaml:"valid"`
}

type AltBlock struct {
	Description string `yaml:"description"`
	Reference   string `yaml:"reference"`
}
