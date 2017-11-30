// OLD
function add1(a, b) {
  return a + b
}
let total1 = add1(1, 4)
console.log('\nแบบเก่า :', total1)
// ------------------------------------------------
// ไม่มีปีก
const add2 = (a, b) => a + b
let total2 = add2(2, 5)
console.log('ไม่มีปีก :', total2);
// ------------------------------------------------
// มีปีก: ทำงาน หลายๆบรรทัก
const add3 = (a, b) => {
  let total3 = a + b
  return total3
}
let total3 = add3(3, 5)
console.log('มีปีก: ทำงาน หลายๆบรรทัก :', total3);
// ------------------------------------------------
// ไม่มีปีก รับ 1 ตัว (สายย่อ)
const add4 = a => a + 5
let total4 = add4(1)
console.log('ไม่มีปีก รับ 1 ตัว (สายย่อ) :', total4);
// ------------------------------------------------
// ทำ FUNC ให้เป็น OBJ
// const newUser = name => { username: name, age: 5 } // ปีกกาของ Array ไปซ้อนกับ {} ของ สายย่อ
const newUser = name => ({
  username: name,
  age: 5
})
let total5 = newUser("Thanakorn")
console.log('ทำ FUNC ให้เป็น OBJ :', total5);
// ------------------------------------------------
// ใช้กับ Map สายวน แล้วย่อ
const nameList = ['Man', 'John', 'Mark']
// แบบที่ 1
// let total6 = nameList.map(function(name) { // แบบยาว
// let total6 = nameList.map(name => {
//   return 'My name is ' + name
// })
// แบบที่ 2
let total6 = nameList.map(name => 'My name is ' + name)
console.log('ใช้กับ Map สายวน แล้วย่อ :', total6);
// ------------------------------------------------
// สาย พลาด
const add7 = (a, b) => {
  a + b
}
// แก้แบบนี้
const add8 = (a, b) => {
  return a + b
}
const add9 = (a, b) => a + b

let total7 = add7(4, 3)
let total8 = add8(5, 3)
let total9 = add9(9, 3)
console.log('สายพลาด :', total7);
console.log('แก้แบบนี้ 1 :', total8);
console.log('แก้แบบนี้ 2 :', total9);
