/* eslint-disable require-jsdoc */
import React from 'react';
import {PainelConversaDiv} from './styles';
import Mensagem from '../Message';
import AreaDeMensagem from '../AreaDeMensagem';

// const DUMMY_DATA = [
//   {
//     senderId: "perborgen",
//     text: "who'll win?"
//   },
//   {
//     senderId: "janedoe",
//     text: "who'll win?"
//   }
// ];

class PainelConversa extends React.Component {
  render() {
    return (
      // <ChatTitulo />
      <PainelConversaDiv>
        <Mensagem autor="teste_usr" conteudo="teste_msg" />
        <AreaDeMensagem roomID="0"/>
      </PainelConversaDiv>
    );
  }
}

export default PainelConversa;
