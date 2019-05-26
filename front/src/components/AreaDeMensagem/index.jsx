/* eslint-disable require-jsdoc */
import React from 'react';
import Axios from 'axios';
import {PainelEnviarMensagem, InputMensagem} from './styles';

class AreaDeMensagem extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      mensagem: '',
      roomID: props.roomID,
      error: null,
    };

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(event) {
    const data = {
      msgID: 10,
      chatID: parseInt(this.state.roomID),
      autorID: 2,
      conteudo: event.target.value,
    };
    // console.log(data);
    if (event.keyCode === 13) {
      Axios.post('/chat/' + this.state.roomID, data)
          .then(
              (response) => {
                console.log(response.data);
              },
              (error) => {
                this.setState({
                  error,
                });
              }
          );
    }
  }

  render() {
    const {error} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else {
      return (
        <PainelEnviarMensagem onKeyDown={this.handleSubmit}>
          <InputMensagem type="text" name="mensagem" autoComplete="off" />
        </PainelEnviarMensagem>
      );
    }
  }
}

export default AreaDeMensagem;
