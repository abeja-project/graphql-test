package abeja

// create this schema from a schema.graphql-file with gogenerate or go 1.16 file
// embedding.
const graphQLSchema = `
type Query {
	todos: [Todo!]!
}

type Mutation {
	createTodo(input: NewTodo!): Todo!
}

type Todo {
	id: ID!
	text: String!
	done: Boolean!
	user: User!
}

type User {
	id: ID!
	name: String!
}

input NewTodo {
	text: String!
	userId: ID!
}
`
