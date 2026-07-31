func isAnagram(s string, t string) bool {
    var count [26]int

    if len(s) != len(t) {
        return false
    }

    for i := range s {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }

    for i := range count {
        if count[i] != 0 {
            return false
        }
    }

    return true
}
