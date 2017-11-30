let state = {
  showing: true,
  title: "My Custom Page",
  choice: {
    first: 1,
    second: 2,
    third: 3
  },
  options: [
    { value: 1, text: 'one' },
    { value: 2, text: 'two' },
    { value: 3, text: 'third' }
  ],
  subTitle: "Helooooooooooo"
}
// แบบโบราณ
// let showing = state.showing
// let title = state.title
// console.log(showing);
// console.log(title);

//                                                แบบ 1 ดึงค่า เอา Field ไม่ลึก obj
// let { showing, title, choice, options } = state
// console.log(showing);
// console.log(title);

//                                                แบบ 2 ดึงค่า เอา Field ลึก ลึก obj
// let val =state.choice.third // เก่า
// console.log('ดึงแบบเก่า', val);
// let { choice: {third} } = state // ใหม่
// console.log('ดึงแบบใหม่', third);

// ใช้รวมกับด้านบนได้ (เปิดพิเศษ)
// let { showing, title, choice: {third}, options } = state
// console.log(third);
// console.log(title);

//                                                          แบบ 3 Array Matching
// let { showing, title, options } = state  //c
// // let [ ant, bird ] = options //c1
// let [ ant, , bird ] = options //c2
// console.log(ant);
// console.log(bird);

// let { showing, title, options: [ optOne, optTwo ]}= state
// console.log(title);
// console.log(optOne);
// console.log(optTwo);

// Default-Destructuring
// let { showing, title, subTitle = 'Default Sub Title' } = state
// console.log(subTitle);

                                          // แบบ 4 : ไม่ต้องจำ ตัวแปร ตอน ส่ง
let state2 = {
  showing: true,
  title: 'My Custom Page'
}
let user = {
  title: 'John'
}
// แบบเก่า
// const printTitle = (title) => {
//   console.log(title);
// }
// printTitle(state2.title)
// printTitle(user.title)

                                // แบบไม่ต้องจำ ตัวแปร 1
// const printTitle = ({title}) => {
//   console.log(title);
// }
// printTitle(state2)
// printTitle(user)

                              // แบบไม่ต้องจำ ตัวแปร 2
let userPageState = {
  username: 'john123',
  password: '12345',
  userType: 'NORMAL',
  remember: true
}
let userPageNewState = {
  username: 'xxxxxxxx',
  password: '11111111',
  userType: 'SPECIAL',
  option: [ 1, 2, 3],
  choices: [ 3, 4, 5]
}
// แบบเก่า
// const doLogin = (username, password, userType) => {
//   console.log(username, password, userType);
// }
// doLogin(userPageState.username, userPageState.password, userPageState.remember)
// วิธีแก้
const doLogin = ({username, password, userType}) => {
  console.log(username, password, userType);
}
doLogin(userPageState)
doLogin(userPageNewState)
doLogin({password: '123123', userType: 'NEW', username: 'John'})
doLogin({password: '123123', userType: 'NEW', username: 'John', ABCABC: 'TESTGEE'})
