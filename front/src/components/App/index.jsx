/* eslint-disable require-jsdoc */
import React from 'react';
import Header from '../Header/index';
import SideBar from '../Sidebar/index';
// import Mensagem from '../Message/index';
import {PageStyle} from './styles.jsx';
import PainelConversa from '../PainelConversa';

class App extends React.Component {
  render() {
    return (
      <PageStyle>
        <Header />
        <SideBar />
        <PainelConversa />
      </PageStyle>
    );
  }
}

export default App;
