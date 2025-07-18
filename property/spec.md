### Simple List Data Structure Specification
### Overview

A Simple List is a linear data structure that maintains an ordered collection of elements. It provides dynamic resizing capabilities and supports insertion, deletion, and access operations at any position.

### Core Properties

### Basic Characteristics

- Ordered: Elements maintain their insertion order and position
- Dynamic: The list can grow and shrink during runtime
- Zero-indexed: First element is at index 0, second at index 1, etc.
- Homogeneous: All elements should be of the same type (integers for this implementation)
- Mutable: Elements can be modified after insertion

### Capacity Management
- Initial capacity: Starts with a reasonable default size (e.g., 10 elements)
- Auto-resize: Automatically grows when more space is needed
- Shrink capability: Can reduce capacity when significantly under-utilized
- Memory efficient: Doesn't waste excessive memory

### Required Operations

### Core Operations

- Add(value) - Appends an element to the end of the list
- Insert(index, value) - Inserts an element at a specific position
- Remove(index) - Removes an element at a specific position
- Get(index) - Retrieves the element at a specific position
- Set(index, value) - Updates the element at a specific position

### Query Operations
- Size() - Returns the number of elements currently in the list
- IsEmpty() - Returns true if the list contains no elements
- Contains(value) - Returns true if the value exists in the list
- IndexOf(value) - Returns the first index where the value is found
- LastIndexOf(value) - Returns the last index where the value is found

### Bulk Operations
- Clear() - Removes all elements from the list
- AddAll(values) - Adds multiple elements to the end of the list
- RemoveAll(value) - Removes all occurrences of a specific value
- ToArray() - Returns a copy of the list as an array/slice

### Utility Operations
- Clone() - Creates a deep copy of the list
- Reverse() - Reverses the order of elements in place
- Sort() - Sorts the elements in ascending order
- Capacity() - Returns the current maximum capacity before resize

### Behavioral Specifications

### Index Validation
- All index-based operations must validate that the index is within bounds
- Valid range: 0 ≤ index < size for access operations
- Valid range: 0 ≤ index ≤ size for insertion operations
- Invalid indices should trigger appropriate error handling

### Memory Management
- When adding elements exceeds capacity, double the current capacity
- When size drops below 25% of capacity and capacity > initial size, halve the capacity
- Never shrink below the initial capacity
- Ensure no memory leaks when removing elements

### Edge Cases
- Operations on empty lists should handle gracefully
- Adding to a full list should trigger automatic resize
- Removing from single-element list should result in empty list
- Negative indices should be rejected
- Null/nil value handling should be clearly defined

### Performance Expectations
- Access (Get/Set): O(1) - Constant time
- Append (Add): O(1) amortized - Constant time on average
- Insert/Remove: O(n) - Linear time due to shifting elements
- Search (Contains/IndexOf): O(n) - Linear time
- Sort: O(n log n) - Efficient sorting algorithm

### Error Handling
### Invalid Operations
- Index out of bounds: Clearly indicate the invalid index and valid range
- Operations on empty list: Handle gracefully with appropriate messages
- Capacity overflow: Handle maximum size limitations
- Invalid input: Validate parameters and provide meaningful errors

### Recovery Strategies
- Invalid operations should not corrupt the list state
- Partial failures in bulk operations should be handled consistently
- Error messages should be descriptive and actionable
