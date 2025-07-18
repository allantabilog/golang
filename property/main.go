// main.go - Demonstration of the List data structure
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("=== Simple List Data Structure Demo ===")
	fmt.Println()

	// Create a new list
	list := NewList()
	fmt.Printf("Created empty list: %s\n", list)
	fmt.Printf("Size: %d, IsEmpty: %v, Capacity: %d\n\n",
		list.Size(), list.IsEmpty(), list.Capacity())

	// Core Operations - Add
	fmt.Println("📝 Adding elements:")
	list.Add(10)
	fmt.Printf("After Add(10): %s\n", list)
	list.Add(20)
	fmt.Printf("After Add(20): %s\n", list)
	list.Add(30)
	fmt.Printf("After Add(30): %s\n", list)
	fmt.Printf("Size: %d, IsEmpty: %v\n\n", list.Size(), list.IsEmpty())

	// Core Operations - Insert
	fmt.Println("➕ Inserting elements:")
	err := list.Insert(1, 15)
	if err != nil {
		log.Printf("Error inserting: %v", err)
	} else {
		fmt.Printf("After Insert(1, 15): %s\n", list)
	}

	err = list.Insert(0, 5)
	if err != nil {
		log.Printf("Error inserting: %v", err)
	} else {
		fmt.Printf("After Insert(0, 5): %s\n", list)
	}
	fmt.Printf("Size: %d\n\n", list.Size())

	// Core Operations - Get
	fmt.Println("🔍 Getting elements:")
	for i := 0; i < list.Size(); i++ {
		value, err := list.Get(i)
		if err != nil {
			log.Printf("Error getting index %d: %v", i, err)
		} else {
			fmt.Printf("Get(%d) = %d\n", i, value)
		}
	}
	fmt.Println()

	// Core Operations - Set
	fmt.Println("✏️ Setting elements:")
	err = list.Set(2, 25)
	if err != nil {
		log.Printf("Error setting: %v", err)
	} else {
		fmt.Printf("After Set(2, 25): %s\n", list)
	}
	fmt.Println()

	// Query Operations
	fmt.Println("❓ Query operations:")
	fmt.Printf("Contains(15): %v\n", list.Contains(15))
	fmt.Printf("Contains(100): %v\n", list.Contains(100))
	fmt.Printf("IndexOf(25): %d\n", list.IndexOf(25))
	fmt.Printf("IndexOf(999): %d\n", list.IndexOf(999))
	fmt.Printf("LastIndexOf(20): %d\n", list.LastIndexOf(20))
	fmt.Println()

	// Add duplicate values to test LastIndexOf
	fmt.Println("🔁 Testing with duplicate values:")
	list.Add(15) // Add another 15
	list.Add(25) // Add another 25
	fmt.Printf("After adding duplicates: %s\n", list)
	fmt.Printf("IndexOf(15): %d\n", list.IndexOf(15))
	fmt.Printf("LastIndexOf(15): %d\n", list.LastIndexOf(15))
	fmt.Printf("IndexOf(25): %d\n", list.IndexOf(25))
	fmt.Printf("LastIndexOf(25): %d\n", list.LastIndexOf(25))
	fmt.Println()

	// Core Operations - Remove
	fmt.Println("🗑️ Removing elements:")
	fmt.Printf("Before removal: %s\n", list)

	removedValue, err := list.Remove(3)
	if err != nil {
		log.Printf("Error removing: %v", err)
	} else {
		fmt.Printf("Removed value %d at index 3: %s\n", removedValue, list)
	}

	removedValue, err = list.Remove(0)
	if err != nil {
		log.Printf("Error removing: %v", err)
	} else {
		fmt.Printf("Removed value %d at index 0: %s\n", removedValue, list)
	}
	fmt.Printf("Final size: %d, Capacity: %d\n\n", list.Size(), list.Capacity())

	// Error handling demonstration
	fmt.Println("⚠️ Error handling:")

	// Try to access invalid indices
	_, err = list.Get(-1)
	if err != nil {
		fmt.Printf("Expected error for Get(-1): %v\n", err)
	}

	_, err = list.Get(100)
	if err != nil {
		fmt.Printf("Expected error for Get(100): %v\n", err)
	}

	err = list.Insert(-1, 999)
	if err != nil {
		fmt.Printf("Expected error for Insert(-1, 999): %v\n", err)
	}

	err = list.Insert(100, 999)
	if err != nil {
		fmt.Printf("Expected error for Insert(100, 999): %v\n", err)
	}

	// Test capacity management
	fmt.Println("\n🔧 Capacity management test:")
	fmt.Printf("Current: %s (size: %d, capacity: %d)\n",
		list, list.Size(), list.Capacity())

	// Add many elements to trigger resize
	fmt.Println("Adding elements to trigger capacity growth...")
	for i := 100; i < 120; i++ {
		list.Add(i)
	}
	fmt.Printf("After adding many elements: size: %d, capacity: %d\n",
		list.Size(), list.Capacity())

	// Remove many elements to trigger shrinking
	fmt.Println("Removing elements to trigger capacity shrinking...")
	for list.Size() > 2 {
		list.Remove(list.Size() - 1)
	}
	fmt.Printf("After removing many elements: %s (size: %d, capacity: %d)\n",
		list, list.Size(), list.Capacity())

	fmt.Println("\n✅ Demo completed successfully!")
}
