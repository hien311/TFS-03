function isPalindrome(s) {
    t = ""
    s = s.toLowerCase()
    for (v of s) {
      if (v.codePointAt(0) >= 97 && v.codePointAt(0) <= 122) {
          t += v
      } 
    }
    i = 0
    while (i <= Math.ceil(t.length / 2)) {
      if (t[i] != t[t.length - 1 - i]) {
        return false
      }
      i++
    }
    return true
}

function uniqueUnion(...theArgs) {
  a = []
  while (theArgs.length > 0) {
    a = [...a,...theArgs.shift()]
  }
  return  [... new Set(a)]
}




