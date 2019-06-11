
import React from 'react';
import Axios from 'axios';
import {PainelEnviarMensagem, InputMensagem,
  PainelDigitarMensagem} from './styles';
import PropTypes from 'prop-types';

class CaixaEnvioMensagem extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      mensagem: '',
      error: null,
      contadorMsgs: 0,
    };

    this.handleSubmit = this.handleSubmit.bind(this);
    this.getCurrentTime = this.getCurrentTime.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  getCurrentTime() {
    const today = new Date();
    let hh = today.getHours();
    let mm = today.getMinutes();

    if (hh < 10) {
      hh = '0' + hh;
    }
    if (mm < 10) {
      mm = '0' + mm;
    }
    return hh + ':' + mm;
  }

  postMensagem(conteudo) {
    Axios.post('/chat/' + this.props.roomID, {
      msgID: this.state.contadorMsgs,
      chatID: parseInt(this.props.roomID),
      autorID: parseInt(this.props.userID),
      conteudo: conteudo,
      hora_enviada: this.getCurrentTime(),
    }).then(
        () => {},
        (error) => {
          this.setState({
            error,
          });
        }
    );
  }

  handleSubmit(event) {
    if (event.keyCode === 13) {
      this.setState(function(state) {
        return {
          counter: state.counter + 1,
        };
      });
      this.postMensagem(event.target.value);
    }
  }


  render() {
    const {error} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else {
      return (
        <PainelEnviarMensagem>
          <PainelDigitarMensagem>
            <form onKeyDown={this.handleSubmit}>
              <InputMensagem type="text" name="mensagem" autoComplete="off"
                placeholder="mande uma mensagem"/>
            </form>
          </PainelDigitarMensagem>
        </PainelEnviarMensagem>
      );
    }
  }
}

CaixaEnvioMensagem.propTypes = {
  roomID: PropTypes.string,
  userID: PropTypes.string,
};

export default CaixaEnvioMensagem;
