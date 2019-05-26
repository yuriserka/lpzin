/* eslint-disable require-jsdoc */
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
      pesquisando: '',
    };
    this.digitando.bind(this);
  }

  serverRequest() {
    Axios.get('/usuarios')
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

  digitando(texto) {
    console.log(texto);
  }

  render() {
    const {error, usuarios} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else {
      return (
        <SidebarDiv>
          <BarraPesquisa value={this.digitando}/>
          <div style={{marginTop: '50px'}}>
            <ConversasDivStyle> {/* copiar pra msg*/}
              {usuarios.map((user) => (
                <ConversaDiv key={user.id}>
                  <Imagem userID={user.id} />
                  {user.nome} enviou {user.mensagem.conteudo}
                  <SeparadorConversa />
                </ConversaDiv>
              ))}
            </ConversasDivStyle>
          </div>
        </SidebarDiv>
      );
    }
  }
}

export default SideBar;
