/* eslint-disable require-jsdoc */
import React from 'react';
import {ConversaHeaderDiv, ConversaInfoDiv} from './styles';

class ConversaHeader extends React.Component {
  render() {
    return (
      <ConversaHeaderDiv>
        <ConversaInfoDiv>
          Eric visto por ultimo hoje
        </ConversaInfoDiv>
      </ConversaHeaderDiv>
    );
  }
}

export default ConversaHeader;
