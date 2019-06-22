import React, { Component } from 'react'
import styled from 'styled-components'
import PropTypes from 'prop-types'
import ListaUsuarios from './ListaUsuarios';

const Input = styled.input`
  height: 48px;
  resize: none;
  width: 100%;
  border-radius: 100px;
  outline: none;
  border: none !important;
  padding-left: 10px;
`;

export class InnerAddChatPopUp extends Component {
  constructor(props) {
    super(props)
    this.state = {
      nome: '',
    }
  }

  changeHandler = (e) => {
    this.setState({ nome: e.target.value })
  }

  addChatHandler = (e) => {
    if (e.keyCode === 13 && e.target.value.length > 0) {
      this.props.addChat({
        ID: this.props.chats.length,
        Nome: e.target.value,
        CriadorID: this.props.myID,
        FotoPerfil: '',
      })
      this.setState({ nome: '' })
    }
  }

  getInnerStyle = () => {
    return {
      height: '250px',
      margin: '15px 10px',
      background: 'red',
      padding: '15px',
      display: 'flex',
      justifyContent: 'center',
    }
  }

  render() {
    return (
      <div style={this.getInnerStyle()}>
        <div style={{ width: '100%' }}>
          <div style={{ marginBottom: '10px', height: '64px' }}>
            <Input placeholder="Digite o nome do Chat" onChange={this.changeHandler} value={this.state.nome}
              onKeyDown={this.addChatHandler} />
          </div>
          <ListaUsuarios usuariosAtivos={this.props.usuariosAtivos} myID={this.props.myID} />
        </div>
      </div>
    )
  }
}

InnerAddChatPopUp.propTypes = {
  chats: PropTypes.array.isRequired,
  usuariosAtivos: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
  addChat: PropTypes.func.isRequired,
}

export default InnerAddChatPopUp
