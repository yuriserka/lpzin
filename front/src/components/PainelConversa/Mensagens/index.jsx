
import React from 'react';
import Axios from 'axios';
import {MensagemDiv, Sender, MessageContent} from './styles';
import PropTypes from 'prop-types';

class Mensagens extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      mensagens: [],
    };

    this.serverRequest = this.serverRequest.bind(this);
    this.componentDidMount = this.componentDidMount.bind(this);
  }

  serverRequest() {
    Axios.get('/chat/' + this.props.roomID)
        .then(
            (result) => {
              this.setState({
                mensagens: result.data,
              });
              console.log('MENSAGENS:', result.data);
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

Mensagens.propTypes = {
  roomID: PropTypes.string,
};

export default Mensagens;
