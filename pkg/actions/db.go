package actions

// Temporary data storage needed for project scope.
// This could use a k-v data storage such as DynamoDB, Redis,
// or Cassandra.
// See README for implementation justification.
// Could use sync.Map?
// TODO: Implementation documentation for choosing map.
var TempData = map[string]ActionData{}
