# Clear Duplicate

## For Tags

Pull all tags from DB, using the unique rule. For each tag `x`, pull all `similar tags`. Replaces all references to the `similar tags` with reference to `x`, that means updating the linking table. Then delete all similar tags.