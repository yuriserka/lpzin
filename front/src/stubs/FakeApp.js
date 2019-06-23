import React from 'react';
import ListaMensagens from '../components/ListaMensagens';
import Sidebar from '../components/Sidebar';
import CaixaEnvio from '../components/CaixaEnvio';
import Header from '../components/Header';
import { StubChats, StubUsuarios } from './index'

class FakeApp extends React.Component {
  constructor() {
    super()
    this.state = {
      chats: null,
      chatAtual: null,
      error: null,
      usuarioAtual: null,
      loaded: false,
      usuariosAtivos: null,
      stubchats: StubChats,
      stubusuarios: StubUsuarios,
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

  StubGetChat = (id) => {
    this.setState({
      chatAtual: this.state.chats[id]
    }, () => console.log('chatAtual', this.state.chatAtual))
  }

  StubGetChats = () => {
    this.setState({
      chats: this.state.stubchats
    }, () => console.log('chats', this.state.chats))
  }

  StubGetUsuarioAtual = () => {
    this.setState({
      usuarioAtual: this.state.stubusuarios[0]
    }, this.StubGetChats)
  }

  StubGetUsuariosAtivos = () => {
    this.setState({
      usuariosAtivos: this.state.stubusuarios
    }, () => console.log('usuariosAtivos', this.state.usuariosAtivos))
  }

  StubAddMensagem = (msg) => {
    this.setState({
      chats: this.state.chats.map(c => {
        if (c.ID === msg.ChatID) {
          c.Mensagens.push(msg)
        }
        return c
      })
    }, () => console.log('chatAtivo.Mensagens', this.state.chatAtual.Mensagens))
  }

  StubAddChat = (chat, usuarios) => {
    if (usuarios.length < 1) {
      alert('É necessário selecionar pelo menos um usuário selecionado')
      return
    }

    chat.ID = this.state.chats.length
    chat.Mensagens = []
    chat.Usuarios = usuarios
    chat.Usuarios.push(this.state.usuarioAtual)
    const noSelecionado = chat.Usuarios.map(u => { delete u.Selecionado; return u })
    chat.Usuarios = noSelecionado

    delete chat.CriadorID

    this.state.chats.push(chat)
    this.setState({
      chats: this.state.chats,
    }, () => console.log('NovoChat.Usuarios', chat.Usuarios))
  }

  async componentDidMount() {
    await this.StubGetUsuarioAtual()
    await this.StubGetUsuariosAtivos()

    this.setState({ loaded: true }, () => console.log('terminou de carregar'))
  }

  render() {
    const {chatAtual, usuarioAtual} = this.state
    const ehGrupo = chatAtual === null ? false : (chatAtual.Usuarios.length > 2 ? true : false)
    if (!this.state.loaded) {
      return <div>Carregando esta poha</div>
    } else {
      return (
        <div >
          {
            usuarioAtual && this.state.chats && this.state.usuariosAtivos &&
            <div className="App" style={this.getAppStyle()}>
              <Header chatAtual={chatAtual} usuarioAtual={usuarioAtual} ehGrupo={ehGrupo} />
              <Sidebar chats={this.state.chats} myID={usuarioAtual.ID} getChat={this.StubGetChat}
                addChat={this.StubAddChat} usuariosAtivos={this.state.usuariosAtivos} />
              <ListaMensagens chatAtual={chatAtual} myID={usuarioAtual.ID} ehGrupo={ehGrupo} />
              <CaixaEnvio chatAtual={chatAtual} usuarioAtual={usuarioAtual} addMensagem={this.StubAddMensagem}
                ehGrupo={ehGrupo} />
            </div>
          }
        </div>
      );
    }
  }
}

export default FakeApp;
