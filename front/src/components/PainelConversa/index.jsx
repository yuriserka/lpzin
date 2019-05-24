/* eslint-disable require-jsdoc */
import React from 'react';
import PainelConversaDiv from './styles';
import Mensagem from '../Message';

class PainelConversa extends React.Component {
  render() {
    return (
      <PainelConversaDiv>
        <Mensagem autor="teste_usr" conteudo="teste_msg" />
      </PainelConversaDiv>
    );
  }
}

export default PainelConversa;
