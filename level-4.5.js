getArray = (value, angka) =>{
  let arrData = []
  let hasil = ''
  for (let j = 0; j < value.length; j++) {
    for (let i = 0; i < angka; i++) {
      if (value[j+i] !== undefined) {
        hasil += value[j+i]
      }
    }
    arrData.push(hasil)
    hasil = ''
    
    if (angka === 2) {
      j = j + 1
    } else if (angka === 3) {
      j = j + 2
    }
  }
  return arrData
}

myFunction = (value) =>{
  let isOk = false
  let arr = []
  let hasil = []
  for (let i = 1; i <= 3; i++) {
   arr.push(getArray(value, i)) 
  }

  for (let j = 0; j < arr.length; j++) {
    for (let k = 0; k < arr[j].length; k++) {
      if (arr[j][k+1] !== undefined) {
        if (arr[j][k+1] - arr[j][k] === 1) {
          isOk = true
          hasil.push(arr[j][k])
        }
        continue
      }
    }
    if (hasil.length < (arr.length / 2)) {
      hasil = []
    }
  }

  return hasil
}

console.log(myFunction('232425272829'))