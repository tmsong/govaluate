package common

/*
	Represents a single parsed token.
*/
type ExpressionToken struct {
	Kind  TokenKind
	Value interface{}
}
