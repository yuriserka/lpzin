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
  constructor(props) {
    super(props)
    this.state = {
      open: false
    }
  }

  openModal = () => {
    this.setState({ open: true })
  }
  closeModal = () => {
    this.setState({ open: false })
  }

  render() {
    return (
      <div>
        <Popup trigger={
          <NovoChatDiv onClick={this.openModal}>
            <ImagemPerfil obj={{
              FotoPerfil: process.env.PUBLIC_URL + '/images/addBtn.png'
            }} />
            <div style={{ margin: '20px 0 0 10px', maxWidth: '75%', minWidth: '50%' }}>
              <span>
                <strong>Adicionar Novo Chat</strong>
              </span>
            </div>
          </NovoChatDiv>
        } modal closeOnDocumentClick>
          <InnerAddChatPopUp usuariosAtivos={this.props.usuariosAtivos} chats={this.props.chats}
            myID={this.props.myID} addChat={this.props.addChat} />
        </Popup>
      </div>
    )
  }
}

NovoChat.propTypes = {
  chats: PropTypes.array.isRequired,
  usuariosAtivos: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
  addChat: PropTypes.func.isRequired,
}

export default NovoChat
