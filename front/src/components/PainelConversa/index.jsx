/* eslint-disable require-jsdoc */
import React from 'react';
import {PainelConversaDiv, InputMensagem, PainelEnviarMensagem} from './styles';
import Axios from 'axios';

class MensagemBox extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      mensagem: '',
      error: null,
    };
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  serverRequest(room, data) {
    Axios.post('/chat/' + room, data)
        .then(
            (response) => {
              console.log('enviado...', response.data);
            },
            (error) => {
              this.setState({
                error,
              });
            }
        );
  }

  handleSubmit(event) {
    if (event.keyCode === 13) {
      const data = {
        msgID: 2,
        chatID: parseInt(this.props.roomID),
        autorID: parseInt(this.props.userID),
        conteudo: event.target.value,
      };
      this.serverRequest(this.props.roomID, data);
    }
  }

  render() {
    const {error} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else {
      return (
        <PainelEnviarMensagem onKeyDown={this.handleSubmit}>
          <InputMensagem type="text" name="mensagem" autoComplete="off"
            placeholder="mande uma mensagem"/>
        </PainelEnviarMensagem>
      );
    }
  }
}

class Mensagens extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      mensagens: [],
      error: null,
    };
    this.serverRequest = this.serverRequest.bind(this);
  }

  serverRequest() {
    Axios.get('/chat/' + this.props.room)
        .then(
            (result) => {
              console.log('MENSAGENS:', result.data);
              this.setState({
                mensagens: result.data,
              });
            },
            (error) => {
              this.setState({
                error,
              });
            }
        );
  }

  componentDidMount() {
    this.serverRequest();
  }

  render() {
    const {error, mensagens} = this.state;
    if (error) {
      alert(error);
      return <div>Error: {error.message}</div>;
    } else {
      return (
        <div>
          <div >
            {mensagens.map((msg) => (
              <div key={msg.msgID} style={{
                background: 'whitesmoke',
                width: '50%', borderRadius: '10px',
                margin: ' auto', marginBottom: '3px',
              }}>
                <div>
                  <span style={{color: 'blue'}}>
                    usr:
                  </span>
                  {msg.autorID}
                </div>
                <div>
                  <span style={{color: 'red'}}>
                    msg:
                  </span>
                  {msg.conteudo}
                </div>
              </div>
            ))}
          </div>
        </div>
      );
    }
  }
}

class PainelConversa extends React.Component {
  render() {
    return (
      <PainelConversaDiv>
        <div>
          <Mensagens room={this.props.roomID} />
        </div>
        <MensagemBox roomID={this.props.roomID} userID={this.props.userID} />
      </PainelConversaDiv>
    );
  }
}

export default PainelConversa;
