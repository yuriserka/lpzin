/* eslint-disable require-jsdoc */
import React from 'react';
import {PainelConversaDiv} from './styles';
import ConversaHeader from './ConversaHeader';
import Mensagens from './Mensagens';
import CaixaEnvioMensagem from './CaixaDeEnvio';

class PainelConversa extends React.Component {
  render() {
    return (
      <span>
        <ConversaHeader/>
        <PainelConversaDiv>
          <div>
            <Mensagens roomID={this.props.roomID} />
          </div>
          <CaixaEnvioMensagem roomID={this.props.roomID} userID={this.props.userID} />
        </PainelConversaDiv>
      </span>
    );
  }
}

export default PainelConversa;
