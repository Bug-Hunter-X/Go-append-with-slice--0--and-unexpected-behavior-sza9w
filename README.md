# Go Append with Slice[:0] Unexpected Behavior

This example demonstrates an unexpected behavior when using `append` with a slice literal `x[:0]`.  It showcases a scenario where the original slice remains unchanged after the append operation, leading to potentially incorrect results.

The issue stems from the use of `x[:0]`, which creates a *new* empty slice instead of using the existing capacity of the original slice `x`.  The append operation then allocates a new slice to accommodate the appended element, leaving the original slice `x` unmodified.