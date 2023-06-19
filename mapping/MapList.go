package mapping

func MapList() (map[string][]string) {
    // Create map of string slices.
    m := map[string][]string {
        "cat": {"orange", "grey"},
        "dog": {"black"},
    }

    // Add a string at the dog key.
    // ... Append returns the new string slice.
    m["dog"] = append(m["dog"], "brown")

    // Add a key for fish.
    m["fish"] = []string{"orange", "red"}

	return m
}

func MapString() (map[string]string) {
    m := map[string]string{
        "A": "Dog",
        "B": "Cat",
        "C": "Cow",
    }

    return m
}

func MapPointer(pointer *map[string]string) (map[string]string) {

    copyPointer := *pointer

    copyPointer["D"] = "Goat"

    return copyPointer
}