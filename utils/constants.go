package utils

const MantleKeyTag = "mantle"
const MantleIndexTag = "index"
const MantleQueryTag = "query"
const GraphQLAllowedCharactersRegex = "^[_A-Za-z][_0-9A-Za-z]*$"

const DepsResolverKey = "graph::depsresolver"
const QuerierKey = "graph::querier"
const DependenciesKey = "graph::dependencies"

type DependenciesKeyType map[string]bool
