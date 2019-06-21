import React from 'react';
import ListaMensagens from './components/ListaMensagens';
import { StubMensagens, StubChats, StubUsuarios } from './stubs';
import Sidebar from './components/Sidebar';
import CaixaEnvio from './components/CaixaEnvio';
import Header from './components/Header';

class App extends React.Component {
  constructor() {
    super()
    this.state = {
      mensagens: StubMensagens,
      chats: StubChats,
      chatAtual: null,
      usuarioAtual: StubUsuarios[0],
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
    this.setState({
      chatAtual: StubChats[IDchatClicado],
    })
  }

  addMensagem = (msg) => {
    console.log(msg)
    this.setState({
      chats: this.state.chats.map(c => {
        if (c.ID === msg.ChatID) {
          c.Mensagens.push(msg)
        }
        return c
      })
    })
  }

  render() {
    return (
      <div className="App" style={this.getAppStyle()}>
        <Header chatAtual={this.state.chatAtual} usuarioAtual={this.state.usuarioAtual} />
        <Sidebar chats={this.state.chats} myID={this.state.usuarioAtual.ID} getChat={this.getChat} />
        <ListaMensagens chatAtual={this.state.chatAtual} mensagens={this.state.mensagens} myID={this.state.usuarioAtual.ID} />
        <CaixaEnvio chatAtual={this.state.chatAtual} usuarioAtual={this.state.usuarioAtual} addMensagem={this.addMensagem}/>
      </div>
    );
  }
}

export default App;
