import React from 'react';
import Mensagem from './Mensagem';
import PropTypes from 'prop-types'
import DefaultPage from './DefaultPage';

class ListaMensagens extends React.Component {
  getMensagens = () => {
    if (this.props.chatAtual === null) {
      return <DefaultPage />
    }
    return this.props.chatAtual.Mensagens.map(
      (msg) => (
        <Mensagem key={msg.ID} chatAtual={this.props.chatAtual} mensagem={msg}
          myID={this.props.myID} />
      )
    )
  }

  render() {
    return (
      <div style={{ height: '500px', overflow: 'auto' }}>
        {this.getMensagens()}
      </div>
    )
  }
}

ListaMensagens.propTypes = {
  mensagens: PropTypes.array.isRequired,
  myID: PropTypes.number.isRequired,
  chatAtual: PropTypes.object,
}

export default ListaMensagens;