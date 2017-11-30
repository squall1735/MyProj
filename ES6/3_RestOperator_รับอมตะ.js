const addPrefix = (prefix, firstname, lastname) => {
  return prefix + " " + firstname + " " + lastname
}
console.log('\n// รับแบบปกติ :', addPrefix('Mr.', 'Thanakorn', 'ChokChai4'));
// -------------------------------------------------------------
const addPrefix2 = (prefix, ...msg) => {
  return prefix + " " + msg
}
console.log('// รับแบบรัวๆ :', addPrefix2('Mr.', 'Thanakorn', 'ChokChai4', 'Vue', 'VueJS'));
// -------------------------------------------------------------
const addPrefix3 = (prefix, ...msg) => {
  return msg.map(m => prefix + m)
}
console.log('// วน Map เติม prefix(ตัวแรก) :', addPrefix3('Mr.', 'Thanakorn', 'ChokChai4', 'Vue', 'VueJS'));
// -------------------------------------------------------------
const addPrefix4 = (prefix, ...msg) => {
  return msg.map(m => 'Hello, ' + m)
}
console.log('// วน Map เติม Hello แทน prefix :\n', addPrefix4('Mr.', 'Thanakorn', 'ChokChai4', 'Vue', 'VueJS'));
// -------------------------------------------------------------
const addPrefix5 = (config, ...msg) => {
  return msg.map(m => config.prefix + m + config.postfix)
}
console.log('\n// ตั้ง config ตัวแรก=Hello ตัวหลัง=! :\n', addPrefix5({prefix: 'Hello, ', postfix: '!'}, 'Thanakorn', 'ChokChai4', 'Vue', 'VueJS'));
