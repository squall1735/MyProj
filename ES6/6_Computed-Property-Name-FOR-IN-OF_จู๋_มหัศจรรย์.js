let obj = {
  ['first'+Math.random()]: "john"
}
console.log('\nobj:', obj, '\n')


let getObj = (prefix) => ({
  [prefix + "_name"]: "john"
})
console.log(getObj('first'))
console.log(getObj('second'))
console.log(getObj('third'))

// obj
console.log('\nOBJ : { a1:.. , a2:.. , a3:.. , a4:.. }');
let obj2 = {
  option1: "1",
  option2: "2",
  option3: "3",
  option4: "4",
  option5: "5"
}
console.log(obj2['option'+1])
console.log(obj2['option'+2])
console.log(obj2['option'+3])
console.log(obj2['option'+4])
console.log(obj2['option'+5])

console.log('\nFOR IN : OBJ');
for (var prop in obj2) {  // FOR IN...
  console.log(`obj.${prop} = ${obj2[prop]}`);
}

// Array
console.log('\nArray : [ {..}, {..}, {..} ]');
let array = [
  {option: "1"},
  {option: "2"},
  {option: "3"},
  {option: "4"},
  {option: "5"}
]
for (var i = 0; i < array.length; i++) {
  console.log('OPTION :', array[i].option)
}

// Array
console.log('\nArray : [ .., .., .. ]');
let iterable = ["a:10", "b:20", "c:30"];
for (let value of iterable) {
  // value += 1;
  console.log(value);
}
// let obj = { a:1 , a2:2 } //in - Obj
// let arr = [ a:1 , a2:2 ] //of - Array
