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