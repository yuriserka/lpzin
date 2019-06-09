/* eslint-disable require-jsdoc */
import React from 'react';
import {PainelConversaDiv} from './styles';
import ConversaHeader from './ConversaHeader/index';
import Mensagens from './Mensagens/index';
import CaixaEnvioMensagem from './CaixaDeEnvio/index';

class PainelConversa extends React.Component {
  render() {
    return (
      <div>
        <ConversaHeader/>
        <PainelConversaDiv>
          <Mensagens roomID={this.props.roomID} />
          <CaixaEnvioMensagem roomID={this.props.roomID} userID={this.props.userID} />
        </PainelConversaDiv>
      </div>
    );
  }
}

export default PainelConversa;
