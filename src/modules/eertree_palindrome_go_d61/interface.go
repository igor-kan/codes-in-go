package mod_eertree_palindrome_go_d61

type PalindromeCounter interface {
	Count(s string) int
}

type TaggedComponent interface {
	ModuleSlug() string
}
