package tinycc

type OutputMode int

const (
	OutputToMemory OutputMode = iota
	OutputToBinary
	OutputToDynamicLibrary
	OutputToObjectFile
	PreprocessOnly
)


