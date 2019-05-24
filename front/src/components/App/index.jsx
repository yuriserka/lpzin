/* eslint-disable require-jsdoc */
import React from 'react';
// import Header from '../Header/index';
// import SideBar from '../Sidebar/index';
// // import Mensagem from '../Message/index';
// import PageStyle from './styles.jsx';
// import PainelConversa from '../PainelConversa';
import axios from 'axios';

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      usuarios: [],
    };
  }

  serverRequest() {
    axios.get('/usuarios')
        .then(
            (result) => {
              // console.log(result.data);
              this.setState({
                isLoaded: true,
                usuarios: result.data,
              });
            },
            (error) => {
              this.setState({
                isLoaded: true,
                error,
              });
            }
        );
  }

  componentDidMount() {
    this.serverRequest();
  }

  render() {
    const {error, isLoaded, usuarios} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
      return <div>Loading...</div>;
    } else {
      return (
        <ul>
          {usuarios.map((user) => (
            <li key={user.id}>
              {user.username} enviou {user.message}
            </li>
          ))}
        </ul>
        // <PageStyle>
        //   <Header />
        //   <SideBar />
        //   <PainelConversa />
        // </PageStyle>
      );
    }
  }
}

export default App;
