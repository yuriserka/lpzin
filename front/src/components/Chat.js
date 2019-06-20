import React from 'react';
import styled from 'styled-components';
import PropTypes from 'prop-types';
import ImagemPerfil from './ImagemPerfil';

const ChatBoxDiv = styled.div`
  display: flex;
  &:hover {
    background: whitesmoke;
  }
`;

const FimChatBox = styled.div`
  border: 0;
  height: 1px;
  background-image: linear-gradient(to right, transparent, silver, transparent);
`;

export class Chat extends React.Component {
  getChatInfoStyles = () => {
    return {
      margin: '20px 0 0 10px',
      maxWidth: '80%',
      minWidth: '50%'
    }
  }

  getPreviewUltimaMensagem = () => {
    const msgs = this.props.chat.Mensagens
    const usrs = this.props.chat.Usuarios
    const ultimaMsg = msgs.length > 0 ? msgs[msgs.length - 1] : undefined
    const ehGrupo = usrs.length > 2 ? true : false

    if (ehGrupo) {
      const autorUltimaMsg = ultimaMsg === undefined ? undefined : usrs.find(usuario => usuario.ID === ultimaMsg.AutorID)
      if (autorUltimaMsg !== undefined) {
        if (autorUltimaMsg.ID === this.props.myID) {
          return ultimaMsg.Conteudo
        }
        return autorUltimaMsg.Nome + ': ' + ultimaMsg.Conteudo
      }
      return ''
    }
    return ultimaMsg === undefined ? '' : ultimaMsg.Conteudo
  }

  getImagem = () => {
    const usrs = this.props.chat.Usuarios
    const ehGrupo = usrs.length > 2 ? true : false
    if (!ehGrupo) {
      const outroUsuario = usrs.find(usuario => usuario.ID !== this.props.myID)
      const ft = outroUsuario === undefined ? '' : outroUsuario.FotoPerfil
      this.props.chat.FotoPerfil = ft
    }
    return <ImagemPerfil obj={this.props.chat} ehGrupo={ehGrupo} />
  }

  getUltimaMensagemStyle = () => {
    return {
      display: 'inline-block',
      flexGrow: '1',
      position: 'relative',
      textOverflow: 'ellipsis',
      whiteSpace: 'nowrap',
      overflow: 'hidden',
      width: '100%',
    }
  }

  render() {
    const { Nome, ID } = this.props.chat
    return (
      <ChatBoxDiv onClick={this.props.getChat.bind(this, ID)}>
        {this.getImagem()}
        <div style={this.getChatInfoStyles()}>
          <span nome={Nome} >
            <strong>{Nome}</strong>
          </span>
          <div>
            <span style={this.getUltimaMensagemStyle()}>
              {this.getPreviewUltimaMensagem()}
            </span>
          </div>
        </div>
        <FimChatBox />
      </ChatBoxDiv>
    )
  }
}

Chat.propTypes = {
  chat: PropTypes.object.isRequired,
  myID: PropTypes.number.isRequired,
}

export default Chat
