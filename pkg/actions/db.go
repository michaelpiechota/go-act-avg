package actions

// Temporary data storage needed for project scope.
// This should be a k-v store such as DynamoDB, Redis,
// or Cassandra.
// See READMEfor implementation justification.
// Could use sync.Map?
// TODO: Implementation documentation for choosing map.
var TempData = map[string]ActionData{}
