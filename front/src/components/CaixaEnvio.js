import React from 'react'
import PropTypes from 'prop-types'
import ImagemPerfil from './ImagemPerfil';

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
      width: '69%',
      height: '100%',
      flexWrap: 'wrap',
      justifyContent: 'center',
    }
  }

  getInputStyle = () => {
    return {
      height: '100%',
      resize: 'none',
      width: '100%',
      borderRadius: '10px',
      outline: 'none',
      WebkitUserSelect: 'none',
      WebkitTouchCallout: 'none',
      MozUserSelect: 'none',
      msUserSelect: 'none',
      KhtmlUserSelect: 'none',
    }
  }

  render() {
    if (this.props.chatAtual === null) {
      return <div></div>
    }
    const chat = this.props.chatAtual
    return (
      <div style={this.getContainerStyle()}>
        <ImagemPerfil obj={this.props.usuarioAtual} />
        <div style={{ padding: '20px 20px 0 20px', height: '100%', width: '50%', }}>
          <div>
            <input placeholder="Escreva uma mensagem..." autoFocus="autofocus" style={this.getInputStyle()}
              onChange={this.changeHandler} onDoubleClick={this.PostHandler(chat)} />
          </div>
        </div>
        <div>
          <ImagemPerfil obj={chat} />
        </div>
      </div>
    )
  }

  PostHandler = (chat) => {
    if (this.state.inputMsg.length > 0) {
      const msg = this.state.inputMsg
      return (
        this.props.addMensagem.bind(this,
          {
            ID: chat.Mensagens.length + 1,
            ChatID: chat.ID,
            AutorID: this.props.usuarioAtual.ID,
            Conteudo: msg,
          }
        )
      )
    }
  }

  changeHandler = (e) => {
    e.preventDefault()
    this.setState({
      inputMsg: e.target.value,
    });
  }
}

CaixaEnvio.propTypes = {
  usuarioAtual: PropTypes.object,
  chatAtual: PropTypes.object,
  addMensagem: PropTypes.func.isRequired,
}

export default CaixaEnvio
