import React from 'react'
import PropTypes from 'prop-types'

export class CaixaEnvio extends React.Component {

  getContainerStyle = () => {
    return {
      background: 'red',
      height: '100%',
    }
  }

  render() {
    return (
      <div style={this.getContainerStyle()}>
        <input placeholder="Digite uma mensagem" />
      </div>
    )
  }
}

CaixaEnvio.propTypes = {
  myID: PropTypes.number.isRequired,
  chatAtual: PropTypes.object.isRequired,
}

export default CaixaEnvio
