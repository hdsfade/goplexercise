//@author: hdsfade
//@date: 2020-11-01-19:54
package main

func nonempty (strings []string) []string {
	i := 0
	for _, s := strings{
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}
