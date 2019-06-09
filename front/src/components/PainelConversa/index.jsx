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
      contadorMsgs: 0,
    };
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
        (result) => {
          console.log(result.data);
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
      error: null,
      mensagens: [],
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
      return <div>Error: {error.message}</div>;
    } else {
      return (
        <footer>
          <div >
          contador de mensagens: {mensagens.length}
            {
              mensagens.map((msg) => (
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
              ))
            }
          </div>
        </footer>
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
