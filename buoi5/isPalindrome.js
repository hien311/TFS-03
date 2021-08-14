function isPalindrome(s) {
    let t = ""
    s = s.toLowerCase()
    for (let v of s) {
      if (v.codePointAt(0) >= 97 && v.codePointAt(0) <= 122) {
          t += v
      } 
    }
    let i = 0
    while (i <= Math.ceil(t.length / 2)) {
      if (t[i] != t[t.length - 1 - i]) {
        return false
      }
      i++
    }
    return true
}

export  {isPalindrome}