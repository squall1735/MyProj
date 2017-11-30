var a = 5
// var b = 10
if (a === 5) {
  var b = 1
  console.log('\nvar B in if :', b);
}
var b = 10
console.log('var B out if :', b, '(เสือกใช้ได้ นอก IF, แถมเปลี่ยนค่า ได้ด้วย)\n');
// ------------------------------------------------
let c = 1
let d = 2
if (c === 1) {
  let d = 20
  console.log('let D in :', d);
}
console.log('let D out :', d, 'ดีมาก อยู่ใน IF แหล่ะ ไม่ต้องออกมาข้างนอก\n');
// ------------------------------------------------
const e = 1
// e = 4 // เปลี่ยน ค่า const ไม่ได้นะ
console.log('const E :', e);
let y = 5
if (y === 5) {
  const z = 1
}
// console.log('const Z :', z) // ใครให้ประกาศ const ใน IF มันใช้ไม่ได้
console.log('const Y :', y);
