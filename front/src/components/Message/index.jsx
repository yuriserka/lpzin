/* eslint-disable require-jsdoc */
import React from 'react';
import {MessageDiv, SenderNameDiv, MsgContentDiv} from './styles';

const Mensagem = (props, {conteudo, autor}) => (
  <MessageDiv>
    <SenderNameDiv>
      {autor}
    </SenderNameDiv>
    <MsgContentDiv>
      {conteudo}
    </MsgContentDiv>
  </MessageDiv>
);

export default Mensagem;
