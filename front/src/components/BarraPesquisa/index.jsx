import React from 'react';
import {SearchWrap, SearchBox, SearchTerm} from './styles';
import Axios from 'axios';

class BarraPesquisa extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      matches: [],
      error: null,
      value: '',
    };
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
    Axios.get('/usuarios').then(
        (result) => {
          const matches = result.data.filter(
              (obj) => Object.keys(obj).some(() =>
                obj.nome.startsWith(this.state.value)));
          this.setState({
            matches: matches,
          });
          console.log(this.state.matches);
        },
        (error) => {
          this.setState({
            error,
          });
        }
    );
  }

  render() {
    const {pesquisados, error} = this.state;
    if (error) {
      return <div>Error: {error.message}</div>;
    } else {
      return (
        <SearchWrap>
          <SearchBox method="get">
            <SearchTerm type="text" placeholder="Pesquisar"
              value={this.state.value} onChange={this.handleChange}
              result={pesquisados} />
          </SearchBox>
        </SearchWrap>
      );
    }
  }
}

export default BarraPesquisa;
