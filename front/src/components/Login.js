// import React, { Component } from 'react'

// export class login extends Component {
//   constructor(props) {
//     super(props)
//     this.state = {
//       username: '',
//       senha: ''
//     }
//   }

//   handleChange = (e) => {
//     this.setState({
//       [e.target.name]: e.target.value
//     })
//   }

//   render() {
//     return (
//       <div>
//         <form action="/login" method="post">
//           <label>
//             Username:
//             <input placeholder="username" name="username" value={this.state.username}
//               onChange={this.handleChange} autoComplete="off" />
//           </label>
//           <label>
//             Senha:
//             <input type="password" name="senha" placeholder="senha" value={this.state.senha}
//               onChange={this.handleChange} />
//           </label>
//           <input type="submit" value="Submit" />
//         </form>
//       </div>
//     )
//   }
// }

// export default login
