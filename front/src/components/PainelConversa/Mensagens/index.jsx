/* eslint-disable require-jsdoc */
import React from 'react';
import Axios from 'axios';
import {MensagemDiv} from './styles';

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
        <footer>
          <div >
            contador de mensagens: {mensagens.length}
            {
              mensagens.map((msg) => (
                <MensagemDiv key={msg.msgID}>
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
                </MensagemDiv>
              ))
            }
          </div>
        </footer>
      );
    }
  }
}

export default Mensagens;
