key = (value) => {
  let isOk = true
  if (value.length < 5 || value.length > 7) {
     console.log('Tidak')
     return
  }

  for (let i = 0; i < value.length; i++) {
    if (value[i] === '5') {
      isOk = true
      continue
    }
    
    if (value[i] > value[i+1]) {
      if ((value[i] - value[i+1]) >= 4) {
        isOk = false
      }
    }
    
    if ((value[i+1] - value[i]) >= 4) {
      isOk = false
    }
  }
  
  if (!isOk) {
    console.log('TIDAK')
    return
  }

  console.log('YA')
}


key('514369')

//!!!wrong
