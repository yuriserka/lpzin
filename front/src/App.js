import React from 'react';
import ListaMensagens from './components/ListaMensagens';
import Sidebar from './components/Sidebar';
import CaixaEnvio from './components/CaixaEnvio';
import Header from './components/Header';
import axios from 'axios';

class App extends React.Component {
  constructor() {
    super()
    this.state = {
      chats: null,
      chatAtual: null,
      error: null,
      usuarioAtual: null,
      loaded: false,
    };
  }

  getAppStyle = () => {
    return {
      backgroundImage: `url(${process.env.PUBLIC_URL + '/images/defaultBG.jpg'})`,
      backgroundRepeat: 'no-repeat',
      backgroundSize: 'cover',
      height: '650px',
      maxWidth: '1010px',
      minWidth: '300px',
      display: 'block',
      borderBottom: '0',
      borderRadius: '0',
      boxShadow: 'none',
      margin: '0 auto',
      borderLeft: '1px solid #dfe5ec',
      borderRight: '1px solid #dfe5ec',
      overflow: 'hidden',
    }
  }

  getChat = (IDchatClicado) => {
    const url = `/chats/${IDchatClicado}`
    axios.get(url).then(
      (res) => {
        this.setState({
          chatAtual: res.data
        }, () => { console.log('GetChat: this.chatAtual terminou', this.chatAtual) })
      },
      (error) => {
        this.setState({
          error: error
        })
      }
    )
  }

  addMensagem = (msg) => {
    const url = `/chats/${String(msg.ChatID)}/mensagens`
    axios.post(url, {
      Conteudo: msg.Conteudo,
      AutorID: msg.AutorID,
    })
      .then(
        (res) => { },
        (error) => {
          this.setState({ error: error })
        }
      )
  }

  getChats = () => {
    const url = `/usuarios/${this.state.usuarioAtual.ID}/chats`
    axios.get(url).then(
      (res) => {
        this.setState({
          chats: res.data,
        },
          () => { console.log('this.state.chats terminou', this.state.chats) })
      }, (error) => {
        this.setState({
          error: error,
        })
      }
    )
  }

  getUsuarioAtual = () => {
    axios.get('/usuarios/1')
      .then(
        (res) => {
          this.setState({
            usuarioAtual: res.data
          }, this.getChats)
        },
        (error) => {
          this.setState({
            error: error,
          })
        }
      )
  }

  async componentDidMount() {
    this.getUsuarioAtual()

    this.setState({ loaded: true })
  }

  render() {
    const chatAtual = this.state.chatAtual
    const usuarioAtual = this.state.usuarioAtual
    const ehGrupo = chatAtual === null ? false : (chatAtual.Usuarios.length > 2 ? true : false)
    if (!this.state.loaded) {
      return <div>Carregando esta poha</div>
    } else {
      return (
        <div >
          {
            usuarioAtual && this.state.chats && <div className="App" style={this.getAppStyle()}>
              <Header chatAtual={chatAtual} usuarioAtual={usuarioAtual} ehGrupo={ehGrupo} />
              <Sidebar chats={this.state.chats} myID={usuarioAtual.ID} getChat={this.getChat} />
              <ListaMensagens chatAtual={chatAtual} myID={usuarioAtual.ID} ehGrupo={ehGrupo} />
              <CaixaEnvio chatAtual={chatAtual} usuarioAtual={usuarioAtual} addMensagem={this.addMensagem}
                ehGrupo={ehGrupo} />
            </div>
          }
        </div>
      );
    }
  }
}

export default App;
