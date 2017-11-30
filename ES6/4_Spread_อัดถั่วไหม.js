let user = {name: 'john', age: '50'}
let userWithTel = {...user, tel: '0863880040'}
let userOverride = {...userWithTel,tel: '2222222222'}
let userOverride2 = {tel: '2222222222',...userWithTel}
console.log('\n// ดึง obj :', userWithTel);
console.log('// ทับตัวแปร :', userOverride);
console.log('// กลับด้าน :', userOverride2, '// ไม่เกิดผล\n');

let user1 = {userId: 1,name: 'John',age: '50'}
let address = {userId: 1,provice: 'Bangkok'}
let userWithAddress = {...user1,...address}
console.log('// ตัวแปรเหมือนกัน คัดให้ :', userWithAddress, '\n');

let studentMale = ['John', 'Man', 'Mike']
let studentFemale = ['Natalee', 'Boa']
let allStudent = studentMale.concat(studentFemale)
let allStudent1 = ['Micheal', ...studentMale,...studentFemale
]
console.log('// concat แบบเก่า :', allStudent);
console.log('// concat แบบ spread :', allStudent1, '\n');

let allStudent2 = ['HEADER', ...studentMale,...studentFemale,'LOAD']
console.log('// เขียน FUNC :', allStudent2);
// จอ HEAD วาด View สรุปข้อมูล-จำนวนรวม
// if (item === 'HEADER') {
//   วาด View HEADER
// } else if (item === 'LOAD') {
//   วาดปุ่ม LOAD MORE
// } else {
//   วาด รายชื่อ นักเรียน
// }
