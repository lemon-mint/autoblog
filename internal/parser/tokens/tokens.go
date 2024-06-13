package tokens

import "fmt"

type TokenType uint16

//go:generate go run golang.org/x/tools/cmd/stringer -type=TokenType
const (
	EOF TokenType = iota
	NEWLINE

	// Block-level tokens
	HEADING1
	HEADING2
	HEADING3
	HEADING4
	HEADING5
	HEADING6
	PARAGRAPH
	BLOCKQUOTE
	CODEBLOCK
	LIST_ITEM // Can be ordered or unordered
	HORIZONTAL_RULE

	// Inline-level tokens
	TEXT
	BOLD
	ITALIC
	CODE
	LINK
	IMAGE

	// Special tokens for parsing
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACKET
	RIGHT_BRACKET
	ASTERISK
	UNDERSCORE
	BACKTICK
	HASH
	GREATER_THAN
	HYPHEN
	PLUS
	DOT
	COLON
	SPACE

	// For ordered lists
	NUMERIC

	// For links
	LINK_TEXT
	LINK_DESTINATION
	LINK_TITLE
)

// Token represents a single token in the input stream.
type Token struct {
	Type    TokenType
	Literal string
	Line    int // Line number where the token starts
	Column  int // Column number where the token starts
}

// String provides a string representation of a Token for debugging.
func (t Token) String() string {
	return fmt.Sprintf("%s(%q)", t.Type, t.Literal)
}
