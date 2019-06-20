import React from 'react';
import PropTypes from 'prop-types';

export class Mensagem extends React.Component {
  getDivStyles = () => {
    return {
      margin: 'auto',
      display: 'flex',
      justifyContent: this.props.mensagem.AutorID === this.props.myID ? 'flex-start' : 'flex-end',
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

  render() {
    return (
      <div style={this.getDivStyles()}>
        <div style={this.getMsgStyles()}>{this.props.mensagem.Conteudo}</div>
      </div>
    )
  }
}

Mensagem.propTypes = {
  mensagem: PropTypes.object.isRequired,
  myID: PropTypes.number.isRequired,
}

export default Mensagem
