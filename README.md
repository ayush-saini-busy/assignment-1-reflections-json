## Problem Overview

Given a JSON input with arbitrary nesting, the goal is to iterate through every element in the JSON structure and print the type and value of each entity.
If an entity is a data structure (such as a map or slice), the traversal should continue recursively until a primitive data type is reached.

This solution uses Go’s reflection (reflect) package to dynamically inspect and process values at runtime.

## Approach

The core idea behind solving this problem is to recursively identify the kind of the underlying value in the JSON using switch statements.

Since JSON unmarshalling in Go results in values of type interface{}, we must first determine the actual underlying type before processing.

## Solution
```
*** Inspecting the JSON ***
Map detected
Key (string): age_in_years
Type: float64 | Value: 8.5
Key (string): origin
Type: string | Value: Noida
Key (string): address
Slice detected
Map detected
Key (string): street
Type: string | Value: 91 Springboard
Key (string): landmark
Type: string | Value: Axis Bank
Key (string): city
Type: string | Value: Noida
Key (string): pincode
Type: float64 | Value: 201301
Key (string): state
Type: string | Value: Uttar Pradesh
Map detected
Key (string): landmark
Type: string | Value: Axis Bank
Key (string): city
Type: string | Value: Noida
Key (string): pincode
Type: float64 | Value: 201301
Key (string): state
Type: string | Value: Uttar Pradesh
Key (string): street
Type: string | Value: 91 Springboard
Key (string): str_text
Slice detected
Type: string | Value: one
Type: string | Value: two
Key (string): int_text
Slice detected
Type: float64 | Value: 1
Type: float64 | Value: 3
Type: float64 | Value: 4
Key (string): name
Type: string | Value: Tolexo Online Pvt. Ltd
Key (string): head_office
Type: string | Value: Noida, Uttar Pradesh
Key (string): sponsers
Map detected
Key (string): name
Type: string | Value: One
Key (string): revenue
Type: string | Value: 19.8 million$
Key (string): no_of_employee
Type: float64 | Value: 630
```
