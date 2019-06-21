import React from 'react';
import PropTypes from 'prop-types';

export class Mensagem extends React.Component {
  getDivStyles = () => {
    return {
      // margin: 'auto',
      display: 'flex',
      justifyContent: this.props.mensagem.AutorID !== this.props.myID ? 'flex-start' : 'flex-end',
      border: '1px solid rgba(0, 0, 0, 0)',
      margin: '0 25px 0 25px',
    }
  }

  getMsgStyles = () => {
    return {
      width: 'max-content',
      padding: '5px',
      background: this.props.mensagem.AutorID !== this.props.myID ? 'lightgreen' : 'lightblue',
      borderRadius: '10px',
    }
  }

  getAutor = (autorID) => {
    const autor = this.props.chatAtual.Usuarios.find(usuario => usuario.ID === autorID)
    if (autor.ID === this.props.myID) {
      return
    }
    return (
      <span style={{ color: `hsl(${autor.Cor.hue}, ${autor.Cor.sat}%, ${autor.Cor.lum}%)` }}>
        <strong>
          {autor.Nome + ': '}
        </strong>
        <br />
      </span>
    )
  }

  printMensagens = () => {
    if (this.props.ehGrupo) {
      return (
        <div style={this.getMsgStyles()}>
          {this.getAutor(this.props.mensagem.AutorID)}
          <span style={{ margin: '0 10px 0 10px' }}>
            {this.props.mensagem.Conteudo}
          </span>
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
  ehGrupo: PropTypes.bool.isRequired,
}

export default Mensagem
