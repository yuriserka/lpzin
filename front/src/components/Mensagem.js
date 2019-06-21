import React from 'react';
import PropTypes from 'prop-types';

export class Mensagem extends React.Component {
  getDivStyles = () => {
    return {
      margin: 'auto',
      display: 'flex',
      justifyContent: this.props.mensagem.AutorID !== this.props.myID ? 'flex-start' : 'flex-end',
      border: '1px solid rgba(0, 0, 0, 0)',
    }
  }

  getMsgStyles = () => {
    return {
      width: 'max-content',
      padding: '5px',
      background: this.props.mensagem.AutorID === this.props.myID ? 'lightgreen' : 'lightblue',
      borderRadius: '10px',
    }
  }

  getSender = (autorID) => {
    const autor = this.props.chatAtual.Usuarios.find(usuario => usuario.ID === autorID)
    if (autor.ID === this.props.myID) {
      return ''
    }
    return (
      <span style={{ color: autor.Cor }}>
        {autor.Nome + ': '}
        <br />
      </span>
    )
  }

  printMensagens = () => {
    const ehGrupo = this.props.chatAtual.Usuarios.length > 2 ? true : false
    if (ehGrupo) {
      return (
        <div>
          <div style={this.getMsgStyles()}>
            {this.getSender(this.props.mensagem.AutorID)}
            <span style={{ margin: '0 10px 0 10px' }}>
              {this.props.mensagem.Conteudo}
            </span>
          </div>
        </div>
      )
    }
    return (
      <div style={this.getMsgStyles()}>
        <span style={{ margin: '0 10px 0 10px' }}>
          {this.props.mensagem.Conteudo}
        </span>
      </div>
    )
  }

  render() {
    return (
      <div style={this.getDivStyles()}>
        {this.printMensagens()}
      </div>
    )
  }
}

Mensagem.propTypes = {
  mensagem: PropTypes.object.isRequired,
  myID: PropTypes.number.isRequired,
  chatAtual: PropTypes.object.isRequired,
}

export default Mensagem
