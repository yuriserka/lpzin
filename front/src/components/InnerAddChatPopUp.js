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
    this.usuariosSelecionados = []
  }

  changeHandler = (e) => {
    this.setState({ nome: e.target.value })
  }

  addChatHandler = (e) => {
    if (e.keyCode === 13 && e.target.value.length > 0) {
      this.props.addChat({
        ID: this.props.chats + 1,
        Nome: e.target.value,
        CriadorID: this.props.myID,
        FotoPerfil: '',
      }, this.usuariosSelecionados)
      this.setState({ nome: '' })
      this.usuariosSelecionados = []
      this.props.onEnter()
    }
  }

  getInnerStyle = () => {
    return {
      height: '300px',
      margin: '15px 10px',
      // background: 'red',
      padding: '15px',
      justifyContent: 'center',
    }
  }

  selecionar = (usuario) => {
    const uss = this.usuariosSelecionados
    uss.push(usuario)
    const realmenteSelecionados = uss.filter(u => u.Selecionado === true)
    this.usuariosSelecionados = realmenteSelecionados
  }


  render() {
    return (
      <div style={this.getInnerStyle()}>
        <div style={{ width: '100%' }}>
          <div style={{ marginBottom: '10px', height: '64px' }}>
            <Input placeholder="Digite o nome do Chat" onChange={this.changeHandler} value={this.state.nome}
              onKeyDown={this.addChatHandler} />
          </div>
          <ListaUsuarios usuariosAtivos={this.props.usuariosAtivos} myID={this.props.myID}
            selecionar={this.selecionar} />
        </div>
      </div>
    )
  }
}

InnerAddChatPopUp.propTypes = {
  chats: PropTypes.array,
  usuariosAtivos: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
  addChat: PropTypes.func.isRequired,
  onEnter: PropTypes.func.isRequired,
}

export default InnerAddChatPopUp
