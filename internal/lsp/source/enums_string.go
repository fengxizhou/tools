// Code generated by "stringer -output enums_string.go -type ImportShortcut,HoverKind,Matcher,SymbolMatcher,SymbolStyle"; DO NOT EDIT.

package source

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Both-0]
	_ = x[Link-1]
	_ = x[Definition-2]
}

const _ImportShortcut_name = "BothLinkDefinition"

var _ImportShortcut_index = [...]uint8{0, 4, 8, 18}

func (i ImportShortcut) String() string {
	if i < 0 || i >= ImportShortcut(len(_ImportShortcut_index)-1) {
		return "ImportShortcut(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ImportShortcut_name[_ImportShortcut_index[i]:_ImportShortcut_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SingleLine-0]
	_ = x[NoDocumentation-1]
	_ = x[SynopsisDocumentation-2]
	_ = x[FullDocumentation-3]
	_ = x[Structured-4]
}

const _HoverKind_name = "SingleLineNoDocumentationSynopsisDocumentationFullDocumentationStructured"

var _HoverKind_index = [...]uint8{0, 10, 25, 46, 63, 73}

func (i HoverKind) String() string {
	if i < 0 || i >= HoverKind(len(_HoverKind_index)-1) {
		return "HoverKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _HoverKind_name[_HoverKind_index[i]:_HoverKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Fuzzy-0]
	_ = x[CaseInsensitive-1]
	_ = x[CaseSensitive-2]
}

const _Matcher_name = "FuzzyCaseInsensitiveCaseSensitive"

var _Matcher_index = [...]uint8{0, 5, 20, 33}

func (i Matcher) String() string {
	if i < 0 || i >= Matcher(len(_Matcher_index)-1) {
		return "Matcher(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Matcher_name[_Matcher_index[i]:_Matcher_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SymbolFuzzy-0]
	_ = x[SymbolCaseInsensitive-1]
	_ = x[SymbolCaseSensitive-2]
}

const _SymbolMatcher_name = "SymbolFuzzySymbolCaseInsensitiveSymbolCaseSensitive"

var _SymbolMatcher_index = [...]uint8{0, 11, 32, 51}

func (i SymbolMatcher) String() string {
	if i < 0 || i >= SymbolMatcher(len(_SymbolMatcher_index)-1) {
		return "SymbolMatcher(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SymbolMatcher_name[_SymbolMatcher_index[i]:_SymbolMatcher_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PackageQualifiedSymbols-0]
	_ = x[FullyQualifiedSymbols-1]
	_ = x[DynamicSymbols-2]
}

const _SymbolStyle_name = "PackageQualifiedSymbolsFullyQualifiedSymbolsDynamicSymbols"

var _SymbolStyle_index = [...]uint8{0, 23, 44, 58}

func (i SymbolStyle) String() string {
	if i < 0 || i >= SymbolStyle(len(_SymbolStyle_index)-1) {
		return "SymbolStyle(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SymbolStyle_name[_SymbolStyle_index[i]:_SymbolStyle_index[i+1]]
}
