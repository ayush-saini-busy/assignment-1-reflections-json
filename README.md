Problem Overview

Given a JSON input with arbitrary nesting, the goal is to iterate through every element in the JSON structure and print the type and value of each entity.
If an entity is a data structure (such as a map or slice), the traversal should continue recursively until a primitive data type is reached.

This solution uses Go’s reflection (reflect) package to dynamically inspect and process values at runtime.

Approach

The core idea behind solving this problem is to recursively identify the kind of the underlying value in the JSON using switch statements.

Since JSON unmarshalling in Go results in values of type interface{}, we must first determine the actual underlying type before processing.
