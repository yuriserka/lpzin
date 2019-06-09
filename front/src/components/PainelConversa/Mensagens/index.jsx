/* eslint-disable require-jsdoc */
import React from 'react';
import Axios from 'axios';
import {MensagemDiv, Sender, MessageContent, GlobalStyles} from './styles';

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
    Axios.get('/chat/' + this.props.roomID)
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
        <div style={{height: '545px', overflow: 'auto'}}>
          <GlobalStyles />
          {
            mensagens.map((msg) => (
              <MensagemDiv key={msg.msgID}>
                <Sender>
                  <span style={{color: 'blue'}}>
                        usr:
                  </span>
                  {msg.autorID}
                </Sender>
                <MessageContent>
                  <span style={{color: 'red', overflowWrap: 'break-word'}}>
                      msg:
                  </span>
                  {msg.conteudo}
                </MessageContent>
              </MensagemDiv>
            ))
          }
        </div>
      );
    }
  }
}

export default Mensagens;
