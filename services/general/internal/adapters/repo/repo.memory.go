package repo

type Memory struct {
	Course   *MemoryCourse
	Material *MemoryMaterial
}

func NewMemory() Memory {
	return Memory{
		Course:   NewMemoryCourse(),
		Material: NewMemoryMaterial(),
	}
}
