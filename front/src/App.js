import React from 'react';
import ListaMensagens from './components/ListaMensagens';
import { StubMensagens, StubChats } from './stubs';
import Sidebar from './components/Sidebar';
import CaixaEnvio from './components/CaixaEnvio';

class App extends React.Component {
  constructor() {
    super()
    this.state = {
      mensagens: StubMensagens,
      chats: StubChats,
      chatAtual: null,
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

  render() {
    return (
      <div className="App" style={this.getAppStyle()}>
        <Sidebar chats={this.state.chats} myID={0} getChat={this.getChat} />
        <ListaMensagens chatAtual={this.state.chatAtual} mensagens={this.state.mensagens} myID={0} />
        <CaixaEnvio chatAtual={this.state.chatAtual} myID={0} />
      </div>
    );
  }
}

export default App;
