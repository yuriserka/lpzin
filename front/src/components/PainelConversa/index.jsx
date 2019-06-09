
import React from 'react';
import {PainelConversaDiv} from './styles';
import ConversaHeader from './ConversaHeader/index';
import Mensagens from './Mensagens/index';
import CaixaEnvioMensagem from './CaixaDeEnvio/index';
import PropTypes from 'prop-types';

class PainelConversa extends React.Component {
  render() {
    return (
      <div>
        <ConversaHeader/>
        <PainelConversaDiv>
          <Mensagens roomID={this.props.roomID} />
          <CaixaEnvioMensagem roomID={this.props.roomID}
            userID={this.props.userID} />
        </PainelConversaDiv>
      </div>
    );
  }
}

PainelConversa.propTypes = {
  roomID: PropTypes.string,
  userID: PropTypes.string,
};

export default PainelConversa;
