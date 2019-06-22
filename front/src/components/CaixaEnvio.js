import React from 'react'
import PropTypes from 'prop-types'
import ImagemPerfil from './ImagemPerfil'
import styled from 'styled-components'

const Input = styled.input`
  height: 100%;
  resize: none;
  width: 100%;
  border-radius: 100px;
  outline: none;
  border: none !important;
  padding-left: 10px; 
  &:focus {
    border-bottom: 20px solid blue;
  }
`

export class CaixaEnvio extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      inputMsg: ''
    }
  }

  getContainerStyle = () => {
    return {
      display: 'flex',
      float: 'left',
      width: '70%',
      height: '100%',
      flexWrap: 'wrap',
      justifyContent: 'center',
      paddingTop: '20px',
    }
  }

  postHandler = (e) => {
    if (e.keyCode === 13 && e.target.value.length > 0) {
      this.props.addMensagem({
        ID: this.props.chatAtual.Mensagens.length + 1,
        ChatID: this.props.chatAtual.ID,
        AutorID: this.props.usuarioAtual.ID,
        Conteudo: e.target.value,
      })
      this.setState({ inputMsg: '' })
    }
  }

  changeHandler = (e) => this.setState({ inputMsg: e.target.value });

  render() {
    if (this.props.chatAtual === null) {
      return <div></div>
    }
    const chat = this.props.chatAtual
    return (
      <div style={this.getContainerStyle()}>
        <ImagemPerfil obj={this.props.usuarioAtual} />
        <div style={{ padding: '0px 20px 0 20px', height: '100%', width: '50%' }}>
          <div style={{ height: '60px' }}>
            <Input placeholder="Escreva uma mensagem..." autoFocus="autofocus"
              onChange={this.changeHandler} value={this.state.inputMsg}
              onKeyDown={this.postHandler} />
          </div>
        </div>
        <ImagemPerfil obj={chat} ehGrupo={this.props.ehGrupo} />
      </div>
    )
  }
}

CaixaEnvio.propTypes = {
  usuarioAtual: PropTypes.object,
  chatAtual: PropTypes.object,
  addMensagem: PropTypes.func.isRequired,
  ehGrupo: PropTypes.bool.isRequired,
}

export default CaixaEnvio
