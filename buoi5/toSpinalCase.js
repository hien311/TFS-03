function toSpinalCase(s) {
    i = 1
    s = s.replace(s[0],s[0].toLowerCase())
    while (i < s.length) {
      if (s[i].codePointAt() > 59 && s[i].codePointAt() < 91) {
        if (s[i-1] != "-") {
          s = s.replace(s[i],"-" + s[i].toLowerCase())
          i +=2
        } else {
          s = s.replace(s[i],s[i].toLowerCase())
          i +=2
        }
      }
      if (s[i].codePointAt() < 97 || s[i].codePointAt() > 122 && s[i] != "-") {
        s = s.replace(s[i], "-")
      }
      i++
    }
    s = s.split("-").reduce((a,b) => {
      if (b != "") {
        return a + "-" + b
      } else return a
    })
    return s 
  }
  
  toSpinalCase("My---Name---is--Hien")