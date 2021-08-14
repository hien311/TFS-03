function uniqueUnion(...theArgs) {
    a = []
    while (theArgs.length > 0) {
      a = [...a,...theArgs.shift()]
    }
    return  [... new Set(a)]
  }
  