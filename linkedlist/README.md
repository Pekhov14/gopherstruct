# Singly Linked List

Data structure of a Singly Linked List

Node is a struct that holds a value and pointer to the next node, address in memory only for example purposes for illustration current address of Node and next node's address in memory.

Without a `tail` pointer.

We can get the address of a node using `&node`.

```go
type Node struct {
    data int
    next *Node
}
```

The `head` field stores a pointer to the first node in the list.

```go
type LinkedList struct {
    head  *Node
    lenth int
}
```

```mermaid
graph LR
    head["head (pointer variable)
      ┌─────────────┐ 
       value: 0x1000 
      └─────────────┘
    "]
    
    node1["Node by address 0x1000
      ┌─────────────┐
       value: 10    
       next: 0x2000
      └─────────────┘
    "]
    
    node2["Node by address 0x2000
      ┌─────────────┐
       value: 20    
       next: nil
      └─────────────┘
    "]
    
    head -->|"stores address"| node1
    node1 -->|"next points to"| node2
    
    style head fill:#ffd700,stroke:#333,stroke-width:3px
    style node1 fill:#e1f5fe,stroke:#01579b
    style node2 fill:#e1f5fe,stroke:#01579b
```
