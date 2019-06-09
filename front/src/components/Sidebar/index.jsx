
import React from 'react';
import Axios from 'axios';
import {
  SidebarDiv,
  ConversaDiv,
  ConversasDivStyle,
  SeparadorConversa,
} from './styles';
import Imagem from '../Imagem';
import BarraPesquisa from '../BarraPesquisa';

class SideBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      usuarios: [],
    };
    this.serverRequest = this.serverRequest.bind(this);
    this.componentDidMount = this.componentDidMount.bind(this);
  }

  serverRequest() {
    Axios.get('/usuarios')
        .then(
            (result) => {
              // console.log(result.data);
              this.setState({
                usuarios: result.data,
              });
            },
            (error) => {
              this.setState({
                error,
              });
            }
        );
  }

  componentDidMount() {
    this.serverRequest();
  }

  render() {
    const {error, usuarios} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else {
      return (
        <SidebarDiv>
          <BarraPesquisa/>
          <div style={{marginTop: '50px'}}>
            <ConversasDivStyle>
              {
                usuarios.map((user) => (
                  <ConversaDiv key={user.id}>
                    <Imagem userID={user.id} />
                    {user.nome} enviou fake_msg_{user.id}
                    <SeparadorConversa />
                  </ConversaDiv>
                ))
              }
            </ConversasDivStyle>
          </div>
        </SidebarDiv>
      );
    }
  }
}

export default SideBar;
