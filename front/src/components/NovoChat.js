import React, { Component } from 'react'
import styled from 'styled-components'
import ImagemPerfil from './ImagemPerfil';
import PropTypes from 'prop-types'
import Popup from 'reactjs-popup'
import InnerAddChatPopUp from './InnerAddChatPopUp';

const NovoChatDiv = styled.div`
  display: flex;
  &:hover {
    background: lightgrey;
  }
  cursor: pointer;
  height: 72px;
`;

export class NovoChat extends Component {
  PopUpStyle = () => {
    return {
      margin: 'auto',
      position: 'relative',
      width: '50%',
      background: 'lightgreen',
      overflowY: 'scroll',
      float: 'left',
      zIndex: '9999'
    }
  }

  PopUpTrigger = () => {
    return (
      <NovoChatDiv>
        <ImagemPerfil obj={{
          FotoPerfil: process.env.PUBLIC_URL + '/images/addBtn.png'
        }} encoded={false} />
        <div style={{ margin: '20px 0 0 10px', maxWidth: '75%', minWidth: '50%' }}>
          <span>
            <strong>Adicionar Novo Chat</strong>
          </span>
        </div>
      </NovoChatDiv>
    )
  }

  render() {
    return (
      <Popup trigger={this.PopUpTrigger()} modal contentStyle={this.PopUpStyle()}
        closeOnDocumentClick closeOnEscape>
        {
          close => (
            <>
              <InnerAddChatPopUp usuariosAtivos={this.props.usuariosAtivos} chats={this.props.chats}
                myID={this.props.myID} addChat={this.props.addChat} onEnter={close} />
            </>
          )
        }

      </Popup>
    )
  }
}

NovoChat.propTypes = {
  chats: PropTypes.array,
  usuariosAtivos: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
  addChat: PropTypes.func.isRequired,
}

export default NovoChat
