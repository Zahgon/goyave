package httputil

type byPriority []HeaderValue

func (s byPriority) Len() int { _ = "STUB: not implemented"; return 0 }

func (s byPriority) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s byPriority) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func specificity(value HeaderValue) int { _ = "STUB: not implemented"; return 0 }
