import React from 'react';
import Header from '../Header/index';
import SideBar from '../Sidebar/index';
import {PageStyle} from './styles.jsx';
import PainelConversa from '../PainelConversa';

class App extends React.Component {
  render() {
    return (
      <PageStyle>
        <Header />
        <SideBar />
        <PainelConversa roomID="0" userID="0"/>
      </PageStyle>
    );
  }
}

export default App;
