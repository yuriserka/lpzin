
import React from 'react';
import {ImagemWrap} from './styles';
import Axios from 'axios';
import PropTypes from 'prop-types';

class Imagem extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      user: {},
      userID: props.userID,
    };

    this.serverRequest.bind(this);
    this.componentDidMount = this.componentDidMount.bind(this);
  }

  serverRequest() {
    Axios.get('/usuarios/' + this.state.userID)
        .then(
            (result) => {
              this.setState({
                user: result.data,
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
    const {error, user} = this.state;
    if (error) {
      return <div>Error: {this.state.error.message}</div>;
    } else {
      return (
        <ImagemWrap>
          <img src={user.fotoPerfil} alt=""
            style={{width: 'auto', height: '100%'}}>
          </img>
        </ImagemWrap>
      );
    }
  }
}

Imagem.propTypes = {
  userID: PropTypes.string,
};

export default Imagem;
