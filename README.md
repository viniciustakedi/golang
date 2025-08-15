# 🚀 Golang Studies

A comprehensive repository dedicated to learning and experimenting with Go programming language concepts, patterns, and best practices.

## 📚 Overview

This repository serves as my personal learning journey through Go's core concepts and advanced features. Each directory contains focused examples and experiments that demonstrate different aspects of Go programming.

## 🎯 Learning Objectives

- Master Go's concurrency primitives (goroutines, channels, select)
- Understand interface design and implementation
- Practice with synchronization mechanisms (WaitGroups, Mutexes)
- Explore JSON handling and data manipulation
- Learn about race conditions and how to prevent them
- Experiment with Go's standard library features

## 📁 Study Topics

### 🔄 **Concurrency & Channels**
- **`channels/`** - Basic channel operations and communication patterns
- **`channels/select/`** - Using select statements for non-blocking operations
- **`channels/another_example/`** - Additional channel examples and patterns

### ⚡ **Goroutines & Synchronization**
- **`goroutines-waitgroups/`** - Managing concurrent operations with WaitGroups
- **`race-condition/`** - Demonstrating and preventing race conditions with mutexes

### 🏗️ **Interfaces & Types**
- **`interfaces/`** - Interface implementation and method receivers

### 📊 **Data Handling**
- **`json/`** - JSON marshaling, unmarshaling, and manipulation
- **`sort-slices/`** - Sorting algorithms and slice operations

### 🛠️ **System Programming**
- **`my-os/`** - Operating system interactions and system calls

## 🚀 Getting Started

### Prerequisites
- Go 1.24.5 or higher
- Basic understanding of programming concepts

### Running Examples

Each directory contains a `main.go` file that can be executed independently:

```bash
# Navigate to any study topic
cd channels
go run main.go

# Or run with race detection
cd race-condition
go run -race main.go
```

## 🔍 Key Concepts Explored

### Concurrency Patterns
- **Channels**: Bidirectional and directional channel usage
- **Select Statements**: Non-blocking channel operations
- **Goroutines**: Lightweight thread management
- **WaitGroups**: Synchronizing multiple goroutines

### Memory Safety
- **Race Conditions**: Identifying and preventing data races
- **Mutexes**: Mutual exclusion for shared resources
- **Atomic Operations**: Thread-safe operations

### Go Idioms
- **Interface Design**: Implementing and using interfaces
- **Method Receivers**: Pointer vs value receivers
- **Error Handling**: Go's error handling patterns

## 🛠️ Development Environment

- **Go Version**: 1.24.5
- **Module**: `studies`
- **OS**: Cross-platform (tested on macOS)

## 📝 Notes

- Each example is self-contained and can be run independently
- Comments in the code explain the concepts being demonstrated
- Some examples include commented-out code showing alternative approaches
- Race condition examples include instructions for using Go's race detector

## 🎓 Learning Resources

This repository complements various Go learning resources including:
- Official Go documentation
- Go by Example
- Effective Go
- Go Concurrency Patterns

## 🤝 Contributing

This is a personal learning repository, but suggestions and improvements are welcome! Feel free to:
- Report issues or bugs
- Suggest additional topics to explore
- Share better examples or explanations

---

**Happy coding and learning! 🎉**
